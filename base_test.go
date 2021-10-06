package document

import (
	"testing"

	"github.com/xeipuuv/gojsonschema"
)

func TestBase(t *testing.T) {
	testSchema(t, BaseType, NewBase())
}

func testSchema(t *testing.T, schemaURI string, doc Document) {
	t.Helper()

	schemaLoader := gojsonschema.NewReferenceLoader(schemaURI)
	schemas := gojsonschema.NewSchemaLoader()

	dataLoader := gojsonschema.NewGoLoader(doc)

	schema, err := schemas.Compile(schemaLoader)
	if err != nil {
		t.Fatal(err)
	}

	_, err = schema.Validate(dataLoader)
	if err != nil {
		t.Fatal(err)
	}
}
