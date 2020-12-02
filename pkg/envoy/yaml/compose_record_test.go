package yaml

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestComposeRecord_UnmarshalYAML(t *testing.T) {
	var (
		parseString = func(src string) (*composeRecord, error) {
			w := &composeRecord{
				values: make(map[string]string),
			}
			return w, yaml.Unmarshal([]byte(src), w)
		}
	)

	t.Run("empty", func(t *testing.T) {
		req := require.New(t)

		w, err := parseString(``)
		req.NoError(err)
		req.NotNil(w)
		req.Empty(w.values)
	})

	t.Run("empty", func(t *testing.T) {
		req := require.New(t)

		w, err := parseString(`{ values: { foo: bar }, createdBy: foo, updatedAt: 2020-10-10T10:10:00Z }`)
		req.NoError(err)
		req.NotNil(w)
		req.NotEmpty(w.values)
		req.NotNil(w.ts)
		req.NotNil(w.us)
		req.Equal("bar", w.values["foo"])
	})

	t.Run("compose record file 1", func(t *testing.T) {
		req := require.New(t)

		doc, err := parseDocument("compose_record_1")
		req.NoError(err)
		req.NotNil(doc)
		req.NotNil(doc.compose)
		req.Len(doc.compose.Records, 3)

		req.NotEmpty(doc.compose.Records[0].values)
		req.Equal("Department", doc.compose.Records[0].refModule)
		v := doc.compose.Records[0].values
		req.Equal("Service", v["Name"])
		req.Equal("50", v["HourCost"])

		req.NotEmpty(doc.compose.Records[1].values)
		req.Equal("EmailTemplate", doc.compose.Records[1].refModule)

		req.NotEmpty(doc.compose.Records[2].values)
		req.Equal("Settings", doc.compose.Records[2].refModule)

		//req.NotNil(doc.compose.records[0].rbac)
		//req.NotEmpty(doc.compose.records[0].rbac.rules)
	})
}
