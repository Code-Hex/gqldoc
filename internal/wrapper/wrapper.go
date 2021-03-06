package wrapper

import "github.com/gqlgo/gqlparser/v2/ast"

type (
	Directive struct {
		Name         string
		Description  string
		Locations    []string
		Args         []InputValue
		IsRepeatable bool
	}

	EnumValue struct {
		Name        string
		Description string
		deprecation *ast.Directive
	}

	Field struct {
		Name        string
		Description string
		Type        *Type
		Args        []InputValue
		deprecation *ast.Directive
	}

	InputValue struct {
		Name         string
		Description  string
		DefaultValue *string
		Type         *Type
	}
)

func WrapSchema(schema *ast.Schema, opts ...WrapSchemaOption) *Schema {
	s := &Schema{schema: schema}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func (f *EnumValue) IsDeprecated() bool {
	return f.deprecation != nil
}

func (f *EnumValue) DeprecationReason() *string {
	if f.deprecation == nil {
		return nil
	}

	reason := f.deprecation.Arguments.ForName("reason")
	if reason == nil {
		return nil
	}

	return &reason.Value.Raw
}

func (f *Field) IsDeprecated() bool {
	return f.deprecation != nil
}

func (f *Field) DeprecationReason() *string {
	if f.deprecation == nil {
		return nil
	}

	reason := f.deprecation.Arguments.ForName("reason")
	if reason == nil {
		return nil
	}

	return &reason.Value.Raw
}
