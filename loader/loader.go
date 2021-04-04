package loader

import (
	"encoding/json"

	"github.com/Code-Hex/gqldoc/internal/gqlgen"
	"github.com/Code-Hex/gqldoc/internal/graphql"
	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/Code-Hex/gqlparser/v2/ast"
	"github.com/Code-Hex/gqlparser/v2/parser"
	"github.com/Code-Hex/gqlparser/v2/validator"
	"github.com/pkg/errors"
)

func LoadSchema(filenames ...string) (*introspection.Root, error) {
	es, err := gqlgen.NewExecutableSchema(filenames...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, err := CreateOperationContext(Params{
		Schema:    es.ParsedSchema,
		Query:     introspection.Query,
		Variables: map[string]interface{}{},
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp := es.Exec(ctx)

	var res introspection.Root
	if err := json.NewDecoder(resp).Decode(&res); err != nil {
		return nil, err
	}
	return &res, nil
}

type Params struct {
	Schema    *ast.Schema
	Query     string
	Variables map[string]interface{}
}

func CreateOperationContext(params Params) (*graphql.OperationContext, error) {
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

	return &graphql.OperationContext{
		RawQuery:  params.Query,
		Variables: variables,
		Doc:       doc,
		Operation: operation,
	}, nil
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
