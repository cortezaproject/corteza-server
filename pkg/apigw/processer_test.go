package apigw

import (
	"net/http"
	"testing"
)

func Test_processerWorkflow(t *testing.T) {
	type (
		tf struct {
			name    string
			expr    string
			err     string
			headers http.Header
		}
	)

	var (
		tcc = []tf{
			{
				name:    "matching simple",
				expr:    `{"expr":"foo == \"bar\""}`,
				headers: map[string][]string{"foo": {"bar"}},
			},
		}
	)

	for _, tc := range tcc {
		var (
		// ctx = context.Background()
		)

		t.Run(tc.name, func(t *testing.T) {
			// req := require.New(t)

			// r, err := http.NewRequest(http.MethodGet, "/foo", http.NoBody)
			// r.Header = tc.headers

			// req.NoError(err)

			// scope := &scp{"request": r}

			// h := NewValidatorHeader()
			// h.Merge([]byte(tc.expr))

			// err = h.Exec(ctx, scope)

			// if tc.err != "" {
			// 	req.EqualError(err, tc.err)
			// } else {
			// 	req.NoError(err)
			// }
		})
	}

	t.Fail()
}
