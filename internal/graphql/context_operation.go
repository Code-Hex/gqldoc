package graphql

import (
	"context"
	"errors"

	"github.com/Code-Hex/gqlparser/v2/ast"
)

type OperationContext struct {
	RawQuery      string
	Variables     map[string]interface{}
	OperationName string
	Doc           *ast.QueryDocument

	Operation *ast.OperationDefinition
}

func (c *OperationContext) Validate(ctx context.Context) error {
	if c.Doc == nil {
		return errors.New("field 'Doc'is required")
	}
	if c.RawQuery == "" {
		return errors.New("field 'RawQuery' is required")
	}
	if c.Variables == nil {
		c.Variables = make(map[string]interface{})
	}
	return nil
}

type operationCtx struct{}

func GetOperationContext(ctx context.Context) *OperationContext {
	if val, ok := ctx.Value(operationCtx{}).(*OperationContext); ok && val != nil {
		return val
	}
	panic("missing operation context")
}

func WithOperationContext(ctx context.Context, rc *OperationContext) context.Context {
	return context.WithValue(ctx, operationCtx{}, rc)
}
