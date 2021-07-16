package apigw

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type (
	mockExecer struct {
		exec func(context.Context, *scp) (err error)
	}

	mockErrorExecer struct {
		exec func(context.Context, *scp, error)
	}
)

func Test_pipelineAdd(t *testing.T) {
	var (
		req = require.New(t)
	)

	p := &pl{}
	p.Add(mockExecer{})

	req.Len(p.w, 1)
}

func Test_pipelineExec(t *testing.T) {
	var (
		ctx   = context.Background()
		req   = require.New(t)
		scope = &scp{"foo": 1}
	)

	p := &pl{}
	p.Add(mockExecer{
		exec: func(c context.Context, s *scp) (err error) {
			s.Set("foo", 2)
			return nil
		},
	})

	err := p.Exec(ctx, scope)

	req.NoError(err)

	foo, err := scope.Get("foo")

	req.NoError(err)
	req.Equal(2, foo)
}

func Test_pipelineExecErr(t *testing.T) {
	var (
		ctx   = context.Background()
		req   = require.New(t)
		scope = &scp{"foo": 1}
	)

	p := &pl{}
	p.Add(mockExecer{
		exec: func(c context.Context, s *scp) (err error) {
			return fmt.Errorf("error returned")
		},
	})

	err := p.Exec(ctx, scope)

	req.Error(err, "error returned")
}

func Test_pipelineExecErrHandler(t *testing.T) {
	var (
		rc    = httptest.NewRecorder()
		ctx   = context.Background()
		req   = require.New(t)
		scope = &scp{"foo": 1, "writer": rc}
	)

	p := &pl{}
	p.Add(mockExecer{
		exec: func(c context.Context, s *scp) (err error) {
			return fmt.Errorf("error returned")
		},
	})

	p.ErrorHandler(mockErrorExecer{
		exec: func(c context.Context, s *scp, e error) {
			s.Writer().WriteHeader(666)
			s.Writer().Write([]byte(`_body_`))
		},
	})

	err := p.Exec(ctx, scope)

	req.Error(err, "error returned")
	req.Equal("_body_", rc.Body.String())
	req.Equal(666, rc.Result().StatusCode)
}

func (me mockExecer) Exec(ctx context.Context, s *scp) (err error) {
	return me.exec(ctx, s)
}

func (me mockErrorExecer) Exec(ctx context.Context, s *scp, e error) {
	me.exec(ctx, s, e)
	return
}
