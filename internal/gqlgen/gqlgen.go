package gqlgen

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"sync"

	gql "github.com/99designs/gqlgen/graphql"
	"github.com/Code-Hex/gqldoc/internal/graphql"
	"github.com/Code-Hex/gqldoc/internal/wrapper"
	gqlparser "github.com/Code-Hex/gqlparser/v2"
	"github.com/Code-Hex/gqlparser/v2/ast"
	"github.com/Code-Hex/gqlparser/v2/parser"
	"github.com/Code-Hex/gqlparser/v2/validator"
	"github.com/pkg/errors"
)

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

func NewExecutableSchema(filenames ...string) (*ExecutableSchema, error) {
	sources := make([]*ast.Source, len(filenames))
	for i, filename := range filenames {
		body, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		sources[i] = &ast.Source{
			Name:  filename,
			Input: string(body),
		}
	}
	AST, err := gqlparser.LoadSchema(sources...)
	if err != nil {
		return nil, err
	}
	return &ExecutableSchema{
		ParsedSchema: AST,
	}, nil
}

type ExecutableSchema struct {
	ParsedSchema *ast.Schema
}

func (e *ExecutableSchema) Exec(oc *graphql.OperationContext) *bytes.Buffer {
	ec := executionContext{oc, e}
	data := ec._Query(context.Background(), oc.Operation.SelectionSet)
	var buf bytes.Buffer
	data.MarshalGQL(&buf)
	return &buf
}

type executionContext struct {
	*graphql.OperationContext
	*ExecutableSchema
}

func (ec *executionContext) introspectSchema() *wrapper.Schema {
	return wrapper.WrapSchema(ec.ParsedSchema)
}

func (ec *executionContext) introspectType(name string) *wrapper.Type {
	return wrapper.WrapTypeFromDef(ec.ParsedSchema, ec.ParsedSchema.Types[name])
}

func (ec *executionContext) field_Query___type_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["name"]; ok {
		arg0, err = gql.UnmarshalString(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["name"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_enumValues_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = gql.UnmarshalBoolean(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (ec *executionContext) field___Type_fields_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 bool
	if tmp, ok := rawArgs["includeDeprecated"]; ok {
		arg0, err = gql.UnmarshalBoolean(tmp)
		if err != nil {
			return nil, err
		}
	}
	args["includeDeprecated"] = arg0
	return args, nil
}

func (ec *executionContext) _Query___type(ctx context.Context, field graphql.CollectedField) (ret gql.Marshaler) {
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field_Query___type_args(ctx, rawArgs)
	if err != nil {
		return gql.Null
	}

	res := ec.introspectType(args["name"].(string))
	return ec.marshalType(ctx, field.Selections, res)
}

func (ec *executionContext) _Query___schema(ctx context.Context, field graphql.CollectedField) (ret gql.Marshaler) {
	res := ec.introspectSchema()
	return ec.___Schema(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_name(ctx context.Context, field graphql.CollectedField, obj *wrapper.Directive) (ret gql.Marshaler) {
	res := obj.Name
	return gql.MarshalString(res)
}

func (ec *executionContext) ___Directive_description(ctx context.Context, field graphql.CollectedField, obj *wrapper.Directive) (ret gql.Marshaler) {
	res := obj.Description
	return gql.MarshalString(res)
}

func (ec *executionContext) ___Directive_locations(ctx context.Context, field graphql.CollectedField, obj *wrapper.Directive) (ret gql.Marshaler) {
	res := obj.Locations
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Directive_args(ctx context.Context, field graphql.CollectedField, obj *wrapper.Directive) (ret gql.Marshaler) {
	res := obj.Args
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___EnumValue_name(ctx context.Context, field graphql.CollectedField, obj *wrapper.EnumValue) (ret gql.Marshaler) {
	res := obj.Name
	return gql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_description(ctx context.Context, field graphql.CollectedField, obj *wrapper.EnumValue) (ret gql.Marshaler) {
	res := obj.Description
	return gql.MarshalString(res)
}

func (ec *executionContext) ___EnumValue_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *wrapper.EnumValue) (ret gql.Marshaler) {
	res := obj.IsDeprecated()
	return gql.MarshalBoolean(res)
}

func (ec *executionContext) ___EnumValue_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *wrapper.EnumValue) (ret gql.Marshaler) {
	res := obj.DeprecationReason()
	if res == nil {
		return gql.Null
	}
	return gql.MarshalString(*res)
}

func (ec *executionContext) ___Field_name(ctx context.Context, field graphql.CollectedField, obj *wrapper.Field) (ret gql.Marshaler) {
	res := obj.Name
	return gql.MarshalString(res)
}

func (ec *executionContext) ___Field_description(ctx context.Context, field graphql.CollectedField, obj *wrapper.Field) (ret gql.Marshaler) {
	res := obj.Description
	return gql.MarshalString(res)
}

func (ec *executionContext) ___Field_args(ctx context.Context, field graphql.CollectedField, obj *wrapper.Field) (ret gql.Marshaler) {
	res := obj.Args
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_type(ctx context.Context, field graphql.CollectedField, obj *wrapper.Field) (ret gql.Marshaler) {
	res := obj.Type
	if res == nil {
		return gql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Field_isDeprecated(ctx context.Context, field graphql.CollectedField, obj *wrapper.Field) (ret gql.Marshaler) {
	res := obj.IsDeprecated()
	return gql.MarshalBoolean(res)
}

func (ec *executionContext) ___Field_deprecationReason(ctx context.Context, field graphql.CollectedField, obj *wrapper.Field) (ret gql.Marshaler) {
	res := obj.DeprecationReason()
	if res == nil {
		return gql.Null
	}
	return gql.MarshalString(*res)
}

func (ec *executionContext) ___InputValue_name(ctx context.Context, field graphql.CollectedField, obj *wrapper.InputValue) (ret gql.Marshaler) {
	res := obj.Name
	return gql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_description(ctx context.Context, field graphql.CollectedField, obj *wrapper.InputValue) (ret gql.Marshaler) {
	res := obj.Description
	return gql.MarshalString(res)
}

func (ec *executionContext) ___InputValue_type(ctx context.Context, field graphql.CollectedField, obj *wrapper.InputValue) (ret gql.Marshaler) {
	res := obj.Type
	if res == nil {
		return gql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___InputValue_defaultValue(ctx context.Context, field graphql.CollectedField, obj *wrapper.InputValue) (ret gql.Marshaler) {
	res := obj.DefaultValue
	if res == nil {
		return gql.Null
	}
	return gql.MarshalString(*res)
}

func (ec *executionContext) ___Schema_types(ctx context.Context, field graphql.CollectedField, obj *wrapper.Schema) (ret gql.Marshaler) {
	res := obj.Types()
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_queryType(ctx context.Context, field graphql.CollectedField, obj *wrapper.Schema) (ret gql.Marshaler) {
	res := obj.QueryType()
	if res == nil {
		return gql.Null
	}
	return ec.___Type(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_mutationType(ctx context.Context, field graphql.CollectedField, obj *wrapper.Schema) (ret gql.Marshaler) {
	res := obj.MutationType()
	return ec.marshalType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_subscriptionType(ctx context.Context, field graphql.CollectedField, obj *wrapper.Schema) (ret gql.Marshaler) {
	res := obj.SubscriptionType()
	return ec.marshalType(ctx, field.Selections, res)
}

func (ec *executionContext) ___Schema_directives(ctx context.Context, field graphql.CollectedField, obj *wrapper.Schema) (ret gql.Marshaler) {
	res := obj.Directives()
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_kind(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	res := obj.Kind()
	return gql.MarshalString(res)
}

func (ec *executionContext) ___Type_name(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	res := obj.Name()
	if res == nil {
		return gql.Null
	}
	return gql.MarshalString(*res)
}

func (ec *executionContext) ___Type_description(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	res := obj.Description()
	return gql.MarshalString(res)
}

func (ec *executionContext) ___Type_fields(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_fields_args(ctx, rawArgs)
	if err != nil {
		return gql.Null
	}

	res := obj.Fields(args["includeDeprecated"].(bool))
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_interfaces(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	res := obj.Interfaces()
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_possibleTypes(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	res := obj.PossibleTypes()
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_enumValues(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	rawArgs := field.ArgumentMap(ec.Variables)
	args, err := ec.field___Type_enumValues_args(ctx, rawArgs)
	if err != nil {
		return gql.Null
	}

	res := obj.EnumValues(args["includeDeprecated"].(bool))
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_inputFields(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	res := obj.InputFields()
	return ec.marshalArray(ctx, field.Selections, res)
}

func (ec *executionContext) ___Type_ofType(ctx context.Context, field graphql.CollectedField, obj *wrapper.Type) (ret gql.Marshaler) {
	res := obj.OfType()
	return ec.marshalType(ctx, field.Selections, res)
}

func (ec *executionContext) _Query(ctx context.Context, sel ast.SelectionSet) gql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, queryImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = gql.MarshalString("Query")
		case "__type":
			out.Values[i] = ec._Query___type(ctx, field)
		case "__schema":
			out.Values[i] = ec._Query___schema(ctx, field)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return gql.Null
	}
	return out
}

func (ec *executionContext) ___Directive(ctx context.Context, sel ast.SelectionSet, obj *wrapper.Directive) gql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __DirectiveImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = gql.MarshalString("__Directive")
		case "name":
			out.Values[i] = ec.___Directive_name(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Directive_description(ctx, field, obj)
		case "locations":
			out.Values[i] = ec.___Directive_locations(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "args":
			out.Values[i] = ec.___Directive_args(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return gql.Null
	}
	return out
}

func (ec *executionContext) ___EnumValue(ctx context.Context, sel ast.SelectionSet, obj *wrapper.EnumValue) gql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __EnumValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = gql.MarshalString("__EnumValue")
		case "name":
			out.Values[i] = ec.___EnumValue_name(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___EnumValue_description(ctx, field, obj)
		case "isDeprecated":
			out.Values[i] = ec.___EnumValue_isDeprecated(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___EnumValue_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return gql.Null
	}
	return out
}

func (ec *executionContext) ___Field(ctx context.Context, sel ast.SelectionSet, obj *wrapper.Field) gql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __FieldImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = gql.MarshalString("__Field")
		case "name":
			out.Values[i] = ec.___Field_name(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___Field_description(ctx, field, obj)
		case "args":
			out.Values[i] = ec.___Field_args(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "type":
			out.Values[i] = ec.___Field_type(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "isDeprecated":
			out.Values[i] = ec.___Field_isDeprecated(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "deprecationReason":
			out.Values[i] = ec.___Field_deprecationReason(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return gql.Null
	}
	return out
}

func (ec *executionContext) ___InputValue(ctx context.Context, sel ast.SelectionSet, obj *wrapper.InputValue) gql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __InputValueImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = gql.MarshalString("__InputValue")
		case "name":
			out.Values[i] = ec.___InputValue_name(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "description":
			out.Values[i] = ec.___InputValue_description(ctx, field, obj)
		case "type":
			out.Values[i] = ec.___InputValue_type(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "defaultValue":
			out.Values[i] = ec.___InputValue_defaultValue(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return gql.Null
	}
	return out
}

func (ec *executionContext) ___Schema(ctx context.Context, sel ast.SelectionSet, obj *wrapper.Schema) gql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __SchemaImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = gql.MarshalString("__Schema")
		case "types":
			out.Values[i] = ec.___Schema_types(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "queryType":
			out.Values[i] = ec.___Schema_queryType(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "mutationType":
			out.Values[i] = ec.___Schema_mutationType(ctx, field, obj)
		case "subscriptionType":
			out.Values[i] = ec.___Schema_subscriptionType(ctx, field, obj)
		case "directives":
			out.Values[i] = ec.___Schema_directives(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return gql.Null
	}
	return out
}

func (ec *executionContext) ___Type(ctx context.Context, sel ast.SelectionSet, obj *wrapper.Type) gql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, __TypeImplementors)

	out := graphql.NewFieldSet(fields)
	var invalids uint32
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = gql.MarshalString("__Type")
		case "kind":
			out.Values[i] = ec.___Type_kind(ctx, field, obj)
			if out.Values[i] == gql.Null {
				invalids++
			}
		case "name":
			out.Values[i] = ec.___Type_name(ctx, field, obj)
		case "description":
			out.Values[i] = ec.___Type_description(ctx, field, obj)
		case "fields":
			out.Values[i] = ec.___Type_fields(ctx, field, obj)
		case "interfaces":
			out.Values[i] = ec.___Type_interfaces(ctx, field, obj)
		case "possibleTypes":
			out.Values[i] = ec.___Type_possibleTypes(ctx, field, obj)
		case "enumValues":
			out.Values[i] = ec.___Type_enumValues(ctx, field, obj)
		case "inputFields":
			out.Values[i] = ec.___Type_inputFields(ctx, field, obj)
		case "ofType":
			out.Values[i] = ec.___Type_ofType(ctx, field, obj)
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch()
	if invalids > 0 {
		return gql.Null
	}
	return out
}

func (ec *executionContext) marshalArray(ctx context.Context, sel ast.SelectionSet, v interface{}) gql.Marshaler {
	if reflect.TypeOf(v).Kind() != reflect.Slice {
		panic("expected slice type")
	}

	slice := reflect.ValueOf(v)
	len := slice.Len()
	if len == 0 {
		return gql.Array{}
	}

	ret := make(gql.Array, len)
	var wg sync.WaitGroup
	isLen1 := len == 1
	if !isLen1 {
		wg.Add(len)
	}
	for i := 0; i < len; i++ {
		i := i
		v := slice.Index(i).Interface()
		f := func(i int) {
			if !isLen1 {
				defer wg.Done()
			}
			switch typ := v.(type) {
			case wrapper.Type:
				ret[i] = ec.___Type(ctx, sel, &typ)
			case wrapper.Directive:
				ret[i] = ec.___Directive(ctx, sel, &typ)
			case wrapper.InputValue:
				ret[i] = ec.___InputValue(ctx, sel, &typ)
			case wrapper.Field:
				ret[i] = ec.___Field(ctx, sel, &typ)
			case wrapper.EnumValue:
				ret[i] = ec.___EnumValue(ctx, sel, &typ)
			case string:
				ret[i] = gql.MarshalString(typ)
			default:
				panic(fmt.Sprintf("unknown type %T", typ))
			}
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()
	return ret
}

func (ec *executionContext) marshalType(ctx context.Context, sel ast.SelectionSet, v *wrapper.Type) gql.Marshaler {
	if v == nil {
		return gql.Null
	}
	return ec.___Type(ctx, sel, v)
}
