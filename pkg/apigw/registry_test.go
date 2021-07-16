package apigw

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type (
	mockHandler struct {
		Foo string `json:"foo"`
	}
)

func Test_registryAddGet(t *testing.T) {
	var (
		req = require.New(t)
		r   = NewRegistry()
	)

	r.Add("mockHandler", mockHandler{})

	h, err := r.Get("mockHandler")

	req.NoError(err)
	req.Len(r.h, 1)
	req.IsType(mockHandler{}, h)
}

func Test_registryAddGetErr(t *testing.T) {
	var (
		req = require.New(t)
		r   = NewRegistry()
	)

	r.Add("mockHandler", mockHandler{})

	h, err := r.Get("foo")

	req.EqualError(err, "could not get element from registry: foo")
	req.Len(r.h, 1)
	req.Nil(h)
}

func Test_registryMerge(t *testing.T) {
	type (
		tf struct {
			name   string
			err    string
			params string
			exp    string
		}
	)

	var (
		tcc = []tf{
			{
				name:   "set params",
				params: `{"foo":"bar"}`,
				exp:    "bar",
			},
			{
				name:   "set invalid params",
				params: `{"foo1":"bar"}`,
				exp:    "",
			},
			{
				name:   "set invalid params err",
				params: `{"foo1":"bar"`,
				exp:    "",
				err:    "unexpected EOF",
			},
		}
	)

	for _, tc := range tcc {
		var (
			req = require.New(t)
			r   = NewRegistry()
		)

		m, err := r.Merge(mockHandler{}, []byte(tc.params))

		if tc.err != "" {
			req.EqualError(err, tc.err)
		} else {
			req.Equal(m.(mockHandler).Foo, tc.exp)
			req.NoError(err)
		}
	}

}

func Test_registryAll(t *testing.T) {
	var (
		req = require.New(t)
		r   = NewRegistry()
	)

	r.Add("mockHandler", mockHandler{})

	list := r.All()

	req.Len(list, 1)
	req.NotEmpty(list[0].Name)
}

func (h mockHandler) Exec(_ context.Context, _ *scp) error {
	panic("not implemented") // TODO: Implement
}

func (h mockHandler) Merge(params []byte) (Handler, error) {
	err := json.NewDecoder(bytes.NewBuffer(params)).Decode(&h)
	return h, err
}

func (h mockHandler) Meta() functionMeta {
	return functionMeta{
		Name: "return mocked function metadata",
	}
}
