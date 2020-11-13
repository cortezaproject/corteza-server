package rdbms

// This file is an auto-generated file
//
// Template:    pkg/codegen/assets/store_rdbms.gen.go.tpl
// Definitions: store/applications.yaml
//
// Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated.

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/cortezaproject/corteza-server/pkg/errors"
	"github.com/cortezaproject/corteza-server/pkg/filter"
	"github.com/cortezaproject/corteza-server/store"
	"github.com/cortezaproject/corteza-server/system/types"
)

var _ = errors.Is

// SearchApplications returns all matching rows
//
// This function calls convertApplicationFilter with the given
// types.ApplicationFilter and expects to receive a working squirrel.SelectBuilder
func (s Store) SearchApplications(ctx context.Context, f types.ApplicationFilter) (types.ApplicationSet, types.ApplicationFilter, error) {
	var (
		err error
		set []*types.Application
		q   squirrel.SelectBuilder
	)

	return set, f, func() error {
		q, err = s.convertApplicationFilter(f)
		if err != nil {
			return err
		}

		// Paging enabled
		// {search: {enablePaging:true}}
		// Cleanup unwanted cursors (only relevant is f.PageCursor, next&prev are reset and returned)
		f.PrevPage, f.NextPage = nil, nil

		if f.PageCursor != nil {
			// Page cursor exists so we need to validate it against used sort
			if f.Sort, err = f.PageCursor.Sort(f.Sort); err != nil {
				return err
			}
		}

		if len(f.Sort) == 0 {
			f.Sort = filter.SortExprSet{}
		}

		// Make sure results are always sorted at least by primary keys
		if f.Sort.Get("id") == nil {
			f.Sort = append(f.Sort, &filter.SortExpr{Column: "id"})
		}

		sort := f.Sort.Clone()

		// When cursor for a previous page is used it's marked as reversed
		// This tells us to flip the descending flag on all used sort keys
		if f.PageCursor != nil && f.PageCursor.Reverse {
			sort.Reverse()
		}

		// Apply sorting expr from filter to query
		if q, err = setOrderBy(q, sort, s.sortableApplicationColumns()); err != nil {
			return err
		}

		set, f.PrevPage, f.NextPage, err = s.fetchFullPageOfApplications(ctx, q, sort.Columns(), sort.Reversed(), f.PageCursor, f.Limit, f.Check)
		if err != nil {
			return err
		}

		f.PageCursor = nil
		return nil
	}()
}

// fetchFullPageOfApplications collects all requested results.
//
// Function applies:
//  - cursor conditions (where ...)
//  - limit
//
// Main responsibility of this function is to perform additional sequential queries in case when not enough results
// are collected due to failed check on a specific row (by check fn).
//
// Function then moves cursor to the last item fetched
func (s Store) fetchFullPageOfApplications(
	ctx context.Context,
	q squirrel.SelectBuilder,
	sortColumns []string,
	sortDesc bool,
	cursor *filter.PagingCursor,
	limit uint,
	check func(*types.Application) (bool, error),
) (set []*types.Application, prev, next *filter.PagingCursor, err error) {
	var (
		aux []*types.Application

		// When cursor for a previous page is used it's marked as reversed
		// This tells us to flip the descending flag on all used sort keys
		reversedOrder = cursor != nil && cursor.Reverse

		// copy of the select builder
		tryQuery squirrel.SelectBuilder

		fetched uint
	)

	set = make([]*types.Application, 0, DefaultSliceCapacity)

	if cursor != nil {
		cursor.Reverse = sortDesc
	}

	for try := 0; try < MaxRefetches; try++ {
		tryQuery = setCursorCond(q, cursor)
		if limit > 0 {
			tryQuery = tryQuery.Limit(uint64(limit + 1))
		}

		if aux, err = s.QueryApplications(ctx, tryQuery, check); err != nil {
			return nil, nil, nil, err
		}

		fetched = uint(len(aux))
		if cursor != nil && prev == nil && fetched > 0 {
			// Cursor for previous page is calculated only when cursor is used (so, not on first page)
			prev = s.collectApplicationCursorValues(aux[0], sortColumns...)
		}

		// Point cursor to the last fetched element
		if fetched > limit && limit > 0 {
			next = s.collectApplicationCursorValues(aux[limit-1], sortColumns...)

			// we should use only as much as requested
			set = append(set, aux[:limit]...)
			break
		} else {
			set = append(set, aux...)
		}

		// if limit is not set or we've already collected enough items
		// we can break the loop right away
		if limit == 0 || fetched == 0 || fetched <= limit {
			break
		}

		// In case limit is set very low and we've missed records in the first fetch,
		// make sure next fetch limit is a bit higher
		if limit < MinEnsureFetchLimit {
			limit = MinEnsureFetchLimit
		}

		// @todo improve strategy for collecting next page with lower limit
	}

	if reversedOrder {
		// Fetched set needs to be reversed because we've forced a descending order to get the previous page
		for i, j := 0, len(set)-1; i < j; i, j = i+1, j-1 {
			set[i], set[j] = set[j], set[i]
		}

		// and flip prev/next cursors too
		prev, next = next, prev
	}

	if prev != nil {
		prev.Reverse = true
	}

	return set, prev, next, nil
}

// QueryApplications queries the database, converts and checks each row and
// returns collected set
//
// Fn also returns total number of fetched items and last fetched item so that the caller can construct cursor
// for next page of results
func (s Store) QueryApplications(
	ctx context.Context,
	q squirrel.Sqlizer,
	check func(*types.Application) (bool, error),
) ([]*types.Application, error) {
	var (
		set = make([]*types.Application, 0, DefaultSliceCapacity)
		res *types.Application

		// Query rows with
		rows, err = s.Query(ctx, q)
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		if err = rows.Err(); err == nil {
			res, err = s.internalApplicationRowScanner(rows)
		}

		if err != nil {
			return nil, err
		}

		// check fn set, call it and see if it passed the test
		// if not, skip the item
		if check != nil {
			if chk, err := check(res); err != nil {
				return nil, err
			} else if !chk {
				continue
			}
		}

		set = append(set, res)
	}

	return set, rows.Err()
}

// LookupApplicationByID searches for application by ID
//
// It returns application even if deleted
func (s Store) LookupApplicationByID(ctx context.Context, id uint64) (*types.Application, error) {
	return s.execLookupApplication(ctx, squirrel.Eq{
		s.preprocessColumn("app.id", ""): store.PreprocessValue(id, ""),
	})
}

// CreateApplication creates one or more rows in applications table
func (s Store) CreateApplication(ctx context.Context, rr ...*types.Application) (err error) {
	for _, res := range rr {
		err = s.checkApplicationConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.execCreateApplications(ctx, s.internalApplicationEncoder(res))
		if err != nil {
			return err
		}
	}

	return
}

// UpdateApplication updates one or more existing rows in applications
func (s Store) UpdateApplication(ctx context.Context, rr ...*types.Application) error {
	return s.partialApplicationUpdate(ctx, nil, rr...)
}

// partialApplicationUpdate updates one or more existing rows in applications
func (s Store) partialApplicationUpdate(ctx context.Context, onlyColumns []string, rr ...*types.Application) (err error) {
	for _, res := range rr {
		err = s.checkApplicationConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.execUpdateApplications(
			ctx,
			squirrel.Eq{
				s.preprocessColumn("app.id", ""): store.PreprocessValue(res.ID, ""),
			},
			s.internalApplicationEncoder(res).Skip("id").Only(onlyColumns...))
		if err != nil {
			return err
		}
	}

	return
}

// UpsertApplication updates one or more existing rows in applications
func (s Store) UpsertApplication(ctx context.Context, rr ...*types.Application) (err error) {
	for _, res := range rr {
		err = s.checkApplicationConstraints(ctx, res)
		if err != nil {
			return err
		}

		err = s.execUpsertApplications(ctx, s.internalApplicationEncoder(res))
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteApplication Deletes one or more rows from applications table
func (s Store) DeleteApplication(ctx context.Context, rr ...*types.Application) (err error) {
	for _, res := range rr {

		err = s.execDeleteApplications(ctx, squirrel.Eq{
			s.preprocessColumn("app.id", ""): store.PreprocessValue(res.ID, ""),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteApplicationByID Deletes row from the applications table
func (s Store) DeleteApplicationByID(ctx context.Context, ID uint64) error {
	return s.execDeleteApplications(ctx, squirrel.Eq{
		s.preprocessColumn("app.id", ""): store.PreprocessValue(ID, ""),
	})
}

// TruncateApplications Deletes all rows from the applications table
func (s Store) TruncateApplications(ctx context.Context) error {
	return s.Truncate(ctx, s.applicationTable())
}

// execLookupApplication prepares Application query and executes it,
// returning types.Application (or error)
func (s Store) execLookupApplication(ctx context.Context, cnd squirrel.Sqlizer) (res *types.Application, err error) {
	var (
		row rowScanner
	)

	row, err = s.QueryRow(ctx, s.applicationsSelectBuilder().Where(cnd))
	if err != nil {
		return
	}

	res, err = s.internalApplicationRowScanner(row)
	if err != nil {
		return
	}

	return res, nil
}

// execCreateApplications updates all matched (by cnd) rows in applications with given data
func (s Store) execCreateApplications(ctx context.Context, payload store.Payload) error {
	return s.Exec(ctx, s.InsertBuilder(s.applicationTable()).SetMap(payload))
}

// execUpdateApplications updates all matched (by cnd) rows in applications with given data
func (s Store) execUpdateApplications(ctx context.Context, cnd squirrel.Sqlizer, set store.Payload) error {
	return s.Exec(ctx, s.UpdateBuilder(s.applicationTable("app")).Where(cnd).SetMap(set))
}

// execUpsertApplications inserts new or updates matching (by-primary-key) rows in applications with given data
func (s Store) execUpsertApplications(ctx context.Context, set store.Payload) error {
	upsert, err := s.config.UpsertBuilder(
		s.config,
		s.applicationTable(),
		set,
		"id",
	)

	if err != nil {
		return err
	}

	return s.Exec(ctx, upsert)
}

// execDeleteApplications Deletes all matched (by cnd) rows in applications with given data
func (s Store) execDeleteApplications(ctx context.Context, cnd squirrel.Sqlizer) error {
	return s.Exec(ctx, s.DeleteBuilder(s.applicationTable("app")).Where(cnd))
}

func (s Store) internalApplicationRowScanner(row rowScanner) (res *types.Application, err error) {
	res = &types.Application{}

	if _, has := s.config.RowScanners["application"]; has {
		scanner := s.config.RowScanners["application"].(func(_ rowScanner, _ *types.Application) error)
		err = scanner(row, res)
	} else {
		err = row.Scan(
			&res.ID,
			&res.Name,
			&res.OwnerID,
			&res.Enabled,
			&res.Unify,
			&res.CreatedAt,
			&res.UpdatedAt,
			&res.DeletedAt,
		)
	}

	if err == sql.ErrNoRows {
		return nil, store.ErrNotFound.Stack(1)
	}

	if err != nil {
		return nil, errors.Store("could not scan application db row").Wrap(err)
	} else {
		return res, nil
	}
}

// QueryApplications returns squirrel.SelectBuilder with set table and all columns
func (s Store) applicationsSelectBuilder() squirrel.SelectBuilder {
	return s.SelectBuilder(s.applicationTable("app"), s.applicationColumns("app")...)
}

// applicationTable name of the db table
func (Store) applicationTable(aa ...string) string {
	var alias string
	if len(aa) > 0 {
		alias = " AS " + aa[0]
	}

	return "applications" + alias
}

// ApplicationColumns returns all defined table columns
//
// With optional string arg, all columns are returned aliased
func (Store) applicationColumns(aa ...string) []string {
	var alias string
	if len(aa) > 0 {
		alias = aa[0] + "."
	}

	return []string{
		alias + "id",
		alias + "name",
		alias + "rel_owner",
		alias + "enabled",
		alias + "unify",
		alias + "created_at",
		alias + "updated_at",
		alias + "deleted_at",
	}
}

// {true true false true true true}

// sortableApplicationColumns returns all Application columns flagged as sortable
//
// With optional string arg, all columns are returned aliased
func (Store) sortableApplicationColumns() map[string]string {
	return map[string]string{
		"id": "id", "name": "name", "created_at": "created_at",
		"createdat":  "created_at",
		"updated_at": "updated_at",
		"updatedat":  "updated_at",
		"deleted_at": "deleted_at",
		"deletedat":  "deleted_at",
	}
}

// internalApplicationEncoder encodes fields from types.Application to store.Payload (map)
//
// Encoding is done by using generic approach or by calling encodeApplication
// func when rdbms.customEncoder=true
func (s Store) internalApplicationEncoder(res *types.Application) store.Payload {
	return store.Payload{
		"id":         res.ID,
		"name":       res.Name,
		"rel_owner":  res.OwnerID,
		"enabled":    res.Enabled,
		"unify":      res.Unify,
		"created_at": res.CreatedAt,
		"updated_at": res.UpdatedAt,
		"deleted_at": res.DeletedAt,
	}
}

// collectApplicationCursorValues collects values from the given resource that and sets them to the cursor
// to be used for pagination
//
// Values that are collected must come from sortable, unique or primary columns/fields
// At least one of the collected columns must be flagged as unique, otherwise fn appends primary keys at the end
//
// Known issue:
//   when collecting cursor values for query that sorts by unique column with partial index (ie: unique handle on
//   undeleted items)
func (s Store) collectApplicationCursorValues(res *types.Application, cc ...string) *filter.PagingCursor {
	var (
		cursor = &filter.PagingCursor{}

		hasUnique bool

		// All known primary key columns

		pkId bool

		collect = func(cc ...string) {
			for _, c := range cc {
				switch c {
				case "id":
					cursor.Set(c, res.ID, false)

					pkId = true
				case "name":
					cursor.Set(c, res.Name, false)

				case "created_at":
					cursor.Set(c, res.CreatedAt, false)

				case "updated_at":
					cursor.Set(c, res.UpdatedAt, false)

				case "deleted_at":
					cursor.Set(c, res.DeletedAt, false)

				}
			}
		}
	)

	collect(cc...)
	if !hasUnique || !(pkId && true) {
		collect("id")
	}

	return cursor
}

// checkApplicationConstraints performs lookups (on valid) resource to check if any of the values on unique fields
// already exists in the store
//
// Using built-in constraint checking would be more performant but unfortunately we can not rely
// on the full support (MySQL does not support conditional indexes)
func (s *Store) checkApplicationConstraints(ctx context.Context, res *types.Application) error {
	// Consider resource valid when all fields in unique constraint check lookups
	// have valid (non-empty) value
	//
	// Only string and uint64 are supported for now
	// feel free to add additional types if needed
	var valid = true

	if !valid {
		return nil
	}

	return nil
}
