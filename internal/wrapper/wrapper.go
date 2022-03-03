package wrapper

import "github.com/vektah/gqlparser/v2/ast"

type (
	Directive struct {
		Name         string
		Locations    []string
		Args         []InputValue
		IsRepeatable bool
		description  string
	}

	EnumValue struct {
		Name        string
		description string
		deprecation *ast.Directive
	}

	Field struct {
		Name        string
		Type        *Type
		Args        []InputValue
		deprecation *ast.Directive
		description string
	}

	InputValue struct {
		Name         string
		DefaultValue *string
		Type         *Type
		description  string
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

func (f *EnumValue) Description() *string {
	if f.description == "" {
		return nil
	}
	return &f.description
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

func (f *Field) Description() *string {
	if f.description == "" {
		return nil
	}
	return &f.description
}

func (f *InputValue) Description() *string {
	if f.description == "" {
		return nil
	}
	return &f.description
}

func (f *Directive) Description() *string {
	if f.description == "" {
		return nil
	}
	return &f.description
}
