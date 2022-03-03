package gqlgen_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/Code-Hex/gqldoc/internal/gqlgen"
	"github.com/Code-Hex/gqldoc/loader"
	"github.com/goccy/go-yaml"
	"github.com/google/go-cmp/cmp"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

func TestExampleSchema(t *testing.T) {
	cases := []string{
		"github",
		"shopify",
		"starwars",
	}
	for _, example := range cases {
		t.Run("example "+example, func(t *testing.T) {
			schema := filepath.Join("..", "..", "example", example, "schema.graphql")
			_, err := loader.LoadSchema(schema)
			if err != nil {
				t.Fatalf("%+v", err)
			}
		})
	}
}

func TestExec(t *testing.T) {
	RunSpec(t, "gqlgen_test.yml", func(t *testing.T, schema, query string) ([]byte, error) {
		source := &ast.Source{
			Name:  "test.gql",
			Input: schema,
		}
		AST, gerr := gqlparser.LoadSchema(source)
		if gerr != nil {
			return nil, gerr
		}

		oc, err := gqlgen.CreateOperationContext(
			gqlgen.Params{
				Schema:    AST,
				Query:     query,
				Variables: map[string]interface{}{},
			},
		)
		if err != nil {
			return nil, err
		}

		es := &gqlgen.ExecutableSchema{
			ParsedSchema: AST,
		}

		resp := es.Exec(oc, gqlgen.WithReservedTypes(true))

		return resp.Bytes(), nil
	})
}

type Features struct {
	Tests map[string][]Spec
}

type Spec struct {
	Name   string
	Schema string
	Query  string
	Error  error
	JSON   string
}

func RunSpec(t *testing.T, filename string, f func(t *testing.T, schema, query string) ([]byte, error)) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var features Features
	err = yaml.Unmarshal(b, &features)
	if err != nil {
		t.Errorf("unable to load %s: %s", filename, err.Error())
		return
	}
	for name, specs := range features.Tests {
		t.Run(name, func(t *testing.T) {
			for _, spec := range specs {
				t.Run(spec.Name, func(t *testing.T) {
					gotJSON, err := f(t, spec.Schema, spec.Query)
					if spec.Error != nil {
						if diff := cmp.Diff(spec.Error, err); diff != "" {
							t.Fatalf("error (-want, +got)\n%s", diff)
						}
					}
					if err != nil {
						t.Fatalf("unexpected error: %v", err)
					}
					if spec.JSON != "" {
						var got, want interface{}
						if err := json.Unmarshal(gotJSON, &got); err != nil {
							t.Fatal(err)
						}
						if err := json.Unmarshal([]byte(spec.JSON), &want); err != nil {
							t.Fatal(err)
						}
						if diff := cmp.Diff(want, got); diff != "" {
							t.Fatalf("json (-want, +got)\n%s", diff)
						}
					}
				})
			}
		})
	}
}
