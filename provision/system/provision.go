package system

import (
	"context"
	"github.com/cortezaproject/corteza-server/pkg/id"
	impAux "github.com/cortezaproject/corteza-server/pkg/importer"
	"github.com/cortezaproject/corteza-server/pkg/rbac"
	"github.com/cortezaproject/corteza-server/pkg/settings"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/importer"
	"github.com/cortezaproject/corteza-server/system/service"
	"github.com/cortezaproject/corteza-server/system/types"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"io"
	"time"
)

type (
	settingsService interface {
		FindByPrefix(context.Context, ...string) (types.SettingValueSet, error)
		BulkSet(context.Context, types.SettingValueSet) error
	}
)

// Check if any roels
func checkRoles(ctx context.Context, s store.Storer) (bool, error) {
	if set, _, err := store.SearchRoles(ctx, s, types.RoleFilter{}); err != nil {
		return false, err
	} else {
		return len(set) > 0, nil
	}
}

// Check if any RBAC rules exist
func checkRbacRules(ctx context.Context, s store.Storer) (bool, error) {
	if set, _, err := store.SearchRbacRules(ctx, s, rbac.RuleFilter{}); err != nil {
		return false, err
	} else {
		return len(set) > 0, nil
	}
}

func Provision(ctx context.Context, log *zap.Logger, s store.Storer) (err error) {
	var (
		hasRoles, hasRbacRules bool
		readers                []io.Reader
	)

	if hasRoles, err = checkRoles(ctx, s); err != nil {
		return err
	} else if !hasRoles {
		rr := types.RoleSet{
			&types.Role{ID: rbac.AdminsRoleID, Name: "Administrators", Handle: "admins"},
			&types.Role{ID: rbac.EveryoneRoleID, Name: "Everyone", Handle: "everyone"},
		}

		if err = rr.Walk(func(r *types.Role) error { return store.CreateRole(ctx, s, r) }); err != nil {
			return err
		}
	}

	if hasRbacRules, err = checkRbacRules(ctx, s); err != nil {
		return err
	}

	if !hasRbacRules {
		log.Info("provisioning system")
		readers, err = impAux.ReadStatic(Asset)
		if err != nil {
			return err
		}

		if err = importer.Import(ctx, readers...); err != nil {
			return err
		}

	} else {
		log.Info("provisioning system settings")
		// When already provisioned, make sure settings are re-provisioned
		readers, err = impAux.ReadStatic(Asset)
		if err != nil {
			return err
		}

		if err = partialImportSettings(ctx, service.DefaultSettings, readers...); err != nil {
			return err
		}
	}

	if err = makeDefaultApplications(ctx, log, s); err != nil {
		return
	}

	if err = authSettingsAutoDiscovery(ctx, log, service.DefaultSettings); err != nil {
		return
	}

	if err = authAddExternals(ctx, log); err != nil {
		return
	}

	if err = service.DefaultSettings.UpdateCurrent(ctx); err != nil {
		return
	}

	if err = oidcAutoDiscovery(ctx, log); err != nil {
		return
	}

	return nil
}

// Updates default application directly in the store
func makeDefaultApplications(ctx context.Context, log *zap.Logger, s store.Storer) error {
	var (
		now        = time.Now()
		aa, _, err = s.SearchApplications(ctx, types.ApplicationFilter{})
	)
	if err != nil {
		return err
	}

	// Update icon & logo on all apps
	const (
		oldIconUrl = "/applications/crust_favicon.png"
		oldLogoUrl = "/applications/crust.jpg"

		newIconUrl = "/applications/default_icon.png"
		newLogoUrl = "/applications/default_logo.jpg"
	)

	for a := 0; a < len(aa); a++ {
		var dirty bool

		if aa[a].Unify == nil {
			continue
		}

		if aa[a].Unify.Icon == oldIconUrl {
			aa[a].Unify.Icon = newIconUrl
			dirty = true
		}

		if aa[a].Unify.Logo == oldLogoUrl {
			aa[a].Unify.Logo = newLogoUrl
			dirty = true
		}

		if !dirty {
			continue
		}

		aa[a].UpdatedAt = &now

		if err = s.UpdateApplication(ctx, aa[a]); err != nil {
			return err
		}
	}

	// List of apps to create.
	//
	// We use Unify.Url field for matching,
	// so make sure it's always present!
	defApps := types.ApplicationSet{
		&types.Application{
			Name:    "Messaging",
			Enabled: true,
			Unify: &types.ApplicationUnify{
				Listed: true,
				Icon:   newIconUrl,
				Logo:   newLogoUrl,
				Url:    "/messaging",
			},
		},

		&types.Application{
			Name:    "Low Code",
			Enabled: true,
			Unify: &types.ApplicationUnify{
				Listed: true,
				Icon:   newIconUrl,
				Logo:   newLogoUrl,
				Url:    "/admin",
			},
		},

		&types.Application{
			Name:    "Corteza Admin Area",
			Enabled: true,
			Unify: &types.ApplicationUnify{
				Listed: true,
				Icon:   newIconUrl,
				Logo:   newLogoUrl,
				Url:    "/admin",
			},
		},

		&types.Application{
			Name:    "Corteza Jitsi Bridge",
			Enabled: true,
			Unify: &types.ApplicationUnify{
				Listed: true,
				Icon:   "/applications/jitsi_icon.png",
				Logo:   "/applications/jitsi.png",
				Url:    "/bridge/jitsi/",
			},
		},

		&types.Application{
			Name:    "Google Maps",
			Enabled: true,
			Unify: &types.ApplicationUnify{
				Listed: true,
				Icon:   "/applications/google_maps_icon.png",
				Logo:   "/applications/google_maps.png",
				Url:    "/bridge/google-maps/",
			},
		},

		&types.Application{
			Name:    "CRM",
			Enabled: true,
			Unify: &types.ApplicationUnify{
				Listed: true,
				Icon:   newIconUrl,
				Logo:   newLogoUrl,
				Url:    "/compose/ns/crm/pages",
			},
		},

		&types.Application{
			Name:    "Service Solution",
			Enabled: true,
			Unify: &types.ApplicationUnify{
				Listed: true,
				Icon:   newIconUrl,
				Logo:   newLogoUrl,
				Url:    "/compose/ns/service-solution/pages",
			},
		},
	}

	return defApps.Walk(func(defApp *types.Application) error {
		for _, a := range aa {
			if a.Unify != nil && a.Unify.Url == defApp.Unify.Url {
				// App already added.
				return nil
			}
		}

		defApp.ID = id.Next()
		defApp.CreatedAt = time.Now()

		err = s.CreateApplication(ctx, defApp)
		log.Info(
			"creating default application",
			zap.String("name", defApp.Name),
			zap.Uint64("name", defApp.ID),
			zap.Error(err),
		)
		return err
	})
}

// Partial import of settings from provision files
func partialImportSettings(ctx context.Context, ss settingsService, ff ...io.Reader) (err error) {
	var (
		// decoded content from YAML files
		aux interface{}

		si = settings.NewImporter()

		// importer w/o permissions & roles
		// we need only settings
		imp = importer.NewImporter(nil, si, nil)

		// current value
		current types.SettingValueSet

		// unexisting values
		unex types.SettingValueSet
	)

	for _, f := range ff {
		if err = yaml.NewDecoder(f).Decode(&aux); err != nil {
			return
		}

		err = imp.Cast(aux)
		if err != nil {
			return
		}
	}

	// Get all "current" settings storage
	current, err = ss.FindByPrefix(ctx)
	if err != nil {
		return
	}

	// Compare current settings with imported, get all that do not exist yet
	if unex = si.GetValues(); len(unex) > 0 {
		// Store non existing
		err = ss.BulkSet(ctx, current.New(unex))
		if err != nil {
			return
		}
	}

	return nil
}
