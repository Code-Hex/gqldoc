package wrapper

import (
	"testing"

	"github.com/Code-Hex/gqlparser/v2/ast"
)

func TestType(t *testing.T) {
	schemaType := Type{
		def: &ast.Definition{
			Name:        "Query",
			Description: "test description",
			Fields: ast.FieldList{
				&ast.FieldDefinition{Name: "__schema"},
				&ast.FieldDefinition{Name: "test"},
				&ast.FieldDefinition{Name: "deprecated", Directives: ast.DirectiveList{
					&ast.Directive{Name: "deprecated"},
				}},
			},
			Kind: ast.Object,
		},
	}

	t.Run("name", func(t *testing.T) {
		got := *schemaType.Name()
		want := "Query"
		if got != want {
			t.Fatalf("want %q but got %q", want, got)
		}
	})

	t.Run("description", func(t *testing.T) {
		got := schemaType.Description()
		want := "test description"
		if got != want {
			t.Fatalf("want %q but got %q", want, got)
		}
	})

	t.Run("fields", func(t *testing.T) {
		fields := schemaType.Fields(false)
		if len(fields) != 1 {
			t.Fatalf("len(fields) = %d", len(fields))
		}
		if fields[0].Name != "test" {
			t.Fatalf("fields[0].Name = %q", fields[0].Name)
		}
	})

	t.Run("fields includeDepricated", func(t *testing.T) {
		fields := schemaType.Fields(true)
		if len(fields) != 2 {
			t.Fatalf("len(fields) = %d", len(fields))
		}
		if fields[0].Name != "test" {
			t.Fatalf("fields[0].Name = %q", fields[0].Name)
		}
		if fields[1].Name != "deprecated" {
			t.Fatalf("fields[0].Name = %q", fields[1].Name)
		}
	})
}
