package renderer

import (
	"context"
	"errors"
	"io"

	"github.com/cortezaproject/corteza-server/pkg/options"
	"github.com/davecgh/go-spew/spew"
)

type (
	renderer struct {
		factories []driverFactory
	}
)

func Renderer(cfg options.TemplateOpt) *renderer {
	ff := make([]driverFactory, 0, 3)
	ff = append(ff, newGenericTXT(), newGenericHTML())
	if cfg.RendererGotenbergEnabled {
		ff = append(ff, newGotenbergPDF(cfg.RendererGotenbergAddress))
	}

	return &renderer{
		factories: []driverFactory{},
	}
}

func (r *renderer) Render(ctx context.Context, pl *RendererPayload) (io.ReadSeeker, error) {
	spew.Dump(pl)
	for _, f := range r.factories {
		if f.CanRender(pl.TemplateType) && f.CanProduce(pl.TargetType) {
			pp := make(map[string]io.Reader)
			for _, prt := range pl.Partials {
				pp[prt.Handle] = prt.Template
			}
			dpl := &driverPayload{
				Template:    pl.Template,
				Variables:   pl.Variables,
				Options:     pl.Options,
				Partials:    pp,
				Attachments: pl.Attachments,
			}

			return f.Driver().Render(ctx, dpl)
		}
	}

	return nil, errors.New("rendering failed: driver not found")
}
