package loader

import (
	"context"
	"encoding/json"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Code-Hex/gqldoc/internal/gqlgen"
	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/parser"
	"github.com/vektah/gqlparser/v2/validator"
)

func LoadSchema(filenames ...string) (*introspection.Root, error) {
	es, err := gqlgen.NewExecutableSchema(filenames...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, err := CreateOperationContext(Params{
		Schema:    es.Schema(),
		Query:     introspection.Query,
		Variables: map[string]interface{}{},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp := es.Exec(ctx)(context.Background())
	if len(resp.Errors) > 0 {
		return nil, resp.Errors
	}

	var res introspection.Root
	if err := json.Unmarshal(resp.Data, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

type Params struct {
	Schema    *ast.Schema
	Query     string
	Variables map[string]interface{}
}

func CreateOperationContext(params Params) (context.Context, error) {
	doc, err := parseQuery(params.Schema, params.Query)
	if err != nil {
		return nil, err
	}
	operation := doc.Operations.ForName("")
	if operation == nil {
		return nil, errors.New("operation not found")
	}
	variables, verr := validator.VariableValues(params.Schema, operation, params.Variables)
	if verr != nil {
		return nil, errors.WithStack(verr)
	}

	opCtx := &graphql.OperationContext{
		RawQuery:  params.Query,
		Variables: variables,
		Doc:       doc,
		Operation: operation,
	}
	ctx := graphql.WithOperationContext(context.Background(), opCtx)
	return ctx, nil
}

func parseQuery(schema *ast.Schema, query string) (*ast.QueryDocument, error) {
	doc, err := parser.ParseQuery(&ast.Source{
		Input: query,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	listErr := validator.Validate(schema, doc)
	if len(listErr) != 0 {
		return nil, listErr
	}
	return doc, nil
}
