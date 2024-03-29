package wrapper

import (
	"sort"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"
)

type WrapSchemaOption func(s *Schema)

func WithReservedTypes(want bool) WrapSchemaOption {
	return func(s *Schema) {
		s.wantReservedTypes = want
	}
}

type Schema struct {
	schema            *ast.Schema
	wantReservedTypes bool
}

func (s *Schema) Types() []Type {
	types := make([]Type, 0, len(s.schema.Types))
	for _, typ := range s.schema.Types {
		if !s.wantReservedTypes && strings.HasPrefix(typ.Name, "__") {
			continue
		}
		types = append(types, *WrapTypeFromDef(s.schema, typ))
	}
	sort.Slice(types, func(i, j int) bool {
		x := *types[i].Name()
		y := *types[j].Name()
		return strings.Compare(x, y) < 0
	})
	return types
}

func (s *Schema) QueryType() *Type {
	return WrapTypeFromDef(s.schema, s.schema.Query)
}

func (s *Schema) MutationType() *Type {
	return WrapTypeFromDef(s.schema, s.schema.Mutation)
}

func (s *Schema) SubscriptionType() *Type {
	return WrapTypeFromDef(s.schema, s.schema.Subscription)
}

func (s *Schema) Description() *string {
	return &s.schema.Description
}

func (s *Schema) Directives() []Directive {
	res := make([]Directive, 0, len(s.schema.Directives))

	for _, d := range s.schema.Directives {
		res = append(res, s.directiveFromDef(d))
	}
	sort.Slice(res, func(i, j int) bool {
		return strings.Compare(res[i].Name, res[j].Name) < 0
	})

	return res
}

func (s *Schema) directiveFromDef(d *ast.DirectiveDefinition) Directive {
	locs := make([]string, len(d.Locations))
	for i, loc := range d.Locations {
		locs[i] = string(loc)
	}

	args := make([]InputValue, len(d.Arguments))
	for i, arg := range d.Arguments {
		args[i] = InputValue{
			Name:         arg.Name,
			DefaultValue: defaultValue(arg.DefaultValue),
			Type:         WrapTypeFromType(s.schema, arg.Type),
			description:  arg.Description,
		}
	}

	return Directive{
		Name:         d.Name,
		Locations:    locs,
		Args:         args,
		IsRepeatable: d.IsRepeatable,
		description:  d.Description,
	}
}
