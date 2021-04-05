package loader

import (
	"encoding/json"

	"github.com/Code-Hex/gqldoc/internal/gqlgen"
	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

func LoadSchema(filenames ...string) (*introspection.Root, error) {
	es, err := gqlgen.NewExecutableSchema(filenames...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, err := gqlgen.CreateOperationContext(gqlgen.Params{
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
