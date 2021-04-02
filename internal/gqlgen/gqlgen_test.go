package gqlgen_test

import (
	"context"
	"encoding/json"
	"path/filepath"
	"testing"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/Code-Hex/gqldoc/internal/gqlgen"
	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/Code-Hex/gqldoc/loader"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"
)

func TestGitHubSchema(t *testing.T) {
	schema := filepath.Join("..", "..", "example", "github", "schema.graphql")
	root, err := loader.LoadSchema(schema)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	root2, err := LoadSchema(schema)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(root2, root, cmpopts.IgnoreUnexported(introspection.Schema{})); diff != "" {
		t.Fatalf("-want, +got\n%s", diff)
	}
}

func TestStarwarsSchema(t *testing.T) {
	schema := filepath.Join("..", "..", "example", "starwars", "schema.graphql")
	root, err := loader.LoadSchema(schema)
	if err != nil {
		t.Fatal(err)
	}

	root2, err := LoadSchema(schema)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(root2, root, cmpopts.IgnoreUnexported(introspection.Schema{})); diff != "" {
		t.Fatalf("-want, +got\n%s", diff)
	}
}

func TestShopifySchema(t *testing.T) {
	schema := filepath.Join("..", "..", "example", "shopify", "schema.graphql")
	root, err := loader.LoadSchema(schema)
	if err != nil {
		t.Fatal(err)
	}

	root2, err := LoadSchema(schema)
	if err != nil {
		t.Fatal(err)
	}

	if diff := cmp.Diff(root2, root, cmpopts.IgnoreUnexported(introspection.Schema{})); diff != "" {
		t.Fatalf("-want, +got\n%s", diff)
	}
}

func LoadSchema(filenames ...string) (*introspection.Root, error) {
	es, err := gqlgen.Deprecated_NewExecutableSchema(filenames...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Set time for tracing
	ctx := graphql.StartOperationTrace(context.Background())

	exec := executor.New(es)
	oc, err := exec.CreateOperationContext(ctx, &graphql.RawParams{
		Query: introspection.Query,
	})
	// Need to do introspection
	oc.DisableIntrospection = false

	responses, ctx := exec.DispatchOperation(ctx, oc)
	resp := responses(ctx)
	if len(resp.Errors) > 0 {
		return nil, resp.Errors
	}

	var res introspection.Root
	if err := json.Unmarshal(resp.Data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
