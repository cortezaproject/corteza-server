package apigw

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	atypes "github.com/cortezaproject/corteza-server/automation/types"
	"github.com/cortezaproject/corteza-server/pkg/expr"
)

type (
	WfExecer interface {
		Exec(ctx context.Context, workflowID uint64, p atypes.WorkflowExecParams) (*expr.Vars, atypes.Stacktrace, error)
	}

	processerWorkflow struct {
		functionMeta
		d WfExecer

		params struct {
			Workflow uint64 `json:"workflow"`
		}
	}

	processerProxy struct {
		functionMeta
		t http.RoundTripper

		params struct {
			Location string `json:"location"`
		}
	}
)

func NewProcesserWorkflow(wf WfExecer) (p *processerWorkflow) {
	p = &processerWorkflow{}

	p.d = wf

	p.Step = 2
	p.Name = "processerWorkflow"
	p.Label = "Workflow processer"
	p.Kind = FunctionKindProcesser

	p.Args = []*functionMetaArg{
		{
			Type:    "workflow",
			Label:   "workflow",
			Options: map[string]interface{}{},
		},
	}

	return
}

func (h processerWorkflow) Meta() functionMeta {
	return h.functionMeta
}

func (f *processerWorkflow) Merge(params []byte) (Handler, error) {
	err := json.NewDecoder(bytes.NewBuffer(params)).Decode(&f.params)
	return f, err
}

func (h processerWorkflow) Exec(ctx context.Context, scope *scp) error {
	var (
		err error
	)

	// setup scope for workflow
	vv := map[string]interface{}{
		"request": scope.Request(),
	}

	// get the request data and put it into vars
	in, err := expr.NewVars(vv)

	if err != nil {
		return err
	}

	wp := atypes.WorkflowExecParams{
		Trace: false,
		// todo depending on settings per-route
		Async: false,
		// todo depending on settings per-route
		Wait:  true,
		Input: in,
	}

	out, _, err := h.d.Exec(ctx, uint64(h.params.Workflow), wp)

	if err != nil {
		return err
	}

	// merge out with scope
	merged, err := in.Merge(out)

	if err != nil {
		return err
	}

	mm, err := expr.CastToVars(merged)

	for k, v := range mm {
		scope.Set(k, v)
	}

	return err
}

func NewProcesserProxy(tr http.RoundTripper) (p *processerProxy) {
	p = &processerProxy{}

	p.t = tr

	p.Step = 2
	p.Name = "processerProxy"
	p.Label = "Proxy processer"
	p.Kind = FunctionKindProcesser

	p.Args = []*functionMetaArg{
		{
			Type:    "text",
			Label:   "location",
			Options: map[string]interface{}{},
		},
	}

	return
}

func (h processerProxy) Meta() functionMeta {
	return h.functionMeta
}

func (f *processerProxy) Merge(params []byte) (Handler, error) {
	err := json.NewDecoder(bytes.NewBuffer(params)).Decode(&f.params)
	return f, err
}

func (h processerProxy) Exec(ctx context.Context, scope *scp) (err error) {
	// do the request
	req := scope.Request()
	outreq := req.Clone(ctx)

	l, err := url.ParseRequestURI(h.params.Location)

	if err != nil {
		return fmt.Errorf("could not parse destination location: %s", err)
	}

	outreq.Header.Set("X-Forwarded-For", req.RemoteAddr)

	outreq.URL = l
	outreq.Method = req.Method
	outreq.Host = l.Hostname()

	resp, err := h.t.RoundTrip(outreq)

	if err != nil {
		return fmt.Errorf("could not proxy request: %s", err)
	}

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("could not read get body on proxy request: %s", err)
	}

	// add meta
	scope.Writer().Header().Set("Content-Type", resp.Header.Get("Content-Type"))
	scope.Writer().Header().Set("Content-Encoding", resp.Header.Get("Content-Encoding"))

	// add to writer
	scope.Writer().Write(b)

	return nil
}
