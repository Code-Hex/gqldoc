package loader

import (
	"context"
	"encoding/json"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/Code-Hex/gqldoc/internal/gqlgen"
	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

func LoadSchema(filenames ...string) (*introspection.Root, error) {
	es, err := gqlgen.NewExecutableSchema(filenames...)
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
