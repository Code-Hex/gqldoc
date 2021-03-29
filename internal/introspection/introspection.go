package introspection

import (
	"encoding/json"
	"errors"
	"sort"
	"strings"

	"github.com/Code-Hex/gqldoc/internal/pool"
)

const Query = `
query IntrospectionQuery {
  __schema {
    queryType {
        name
	    description
	    kind
    }
    mutationType {
        name
	    description
	    kind
    }
    subscriptionType {
        name
	    description
	    kind
    }
    types {
      ...FullType
    }
    directives {
      name
      description
      locations
      args {
        ...InputValue
      }
    }
  }
}
fragment FullType on __Type {
  kind
  name
  description
  fields(includeDeprecated: true) {
    name
    description
    args {
      ...InputValue
    }
    type {
      ...TypeRef
    }
    isDeprecated
    deprecationReason
  }
  inputFields {
    ...InputValue
  }
  interfaces {
    ...TypeRef
  }
  enumValues(includeDeprecated: true) {
    name
    description
    isDeprecated
    deprecationReason
  }
  possibleTypes {
    ...TypeRef
  }
}
fragment InputValue on __InputValue {
  name
  description
  type {
    ...TypeRef
  }
  defaultValue
}
fragment TypeRef on __Type {
  kind
  name
  ofType {
    kind
    name
    ofType {
      kind
      name
      ofType {
        kind
        name
        ofType {
          kind
          name
          ofType {
            kind
            name
            ofType {
              kind
              name
              ofType {
                kind
                name
              }
            }
          }
        }
      }
    }
  }
}
`

// https://github.com/graphql/graphql-js/blob/540f33605ee047ab86e0d9666f4a60526d56ff0f/src/type/introspection.js#L427-L436
type TypeKind string

const (
	SCALAR       TypeKind = "SCALAR"
	OBJECT       TypeKind = "OBJECT"
	INTERFACE    TypeKind = "INTERFACE"
	UNION        TypeKind = "UNION"
	ENUM         TypeKind = "ENUM"
	INPUT_OBJECT TypeKind = "INPUT_OBJECT"
	LIST         TypeKind = "LIST"
	NON_NULL     TypeKind = "NON_NULL"
)

type Root struct {
	Schema *Schema `json:"__schema"`
}

type Schema struct {
	QueryType        *OperationType `json:"queryType"`
	MutationType     *OperationType `json:"mutationType"`
	SubscriptionType *OperationType `json:"subscriptionType"`
	Types            []*Types       `json:"types"`
	Directives       []*Directives  `json:"directives"`

	memoizeTypes map[string]*Types
}

var _ json.Unmarshaler = (*Schema)(nil)

var ErrNotFoundTypes = errors.New("not found types")

func (s *Schema) FindTypes(typName string) (*Types, error) {
	memo, ok := s.memoizeTypes[typName]
	if ok {
		return memo, nil
	}
	return nil, ErrNotFoundTypes
}

func (s *Schema) UnmarshalJSON(data []byte) error {
	// HACK(codehex): To avoid stackoverflow, I assigned to tmp type.
	type HookSchema Schema
	var schema HookSchema
	if err := json.Unmarshal(data, &schema); err != nil {
		return err
	}
	sort.Slice(schema.Types, func(i, j int) bool {
		return strings.Compare(schema.Types[i].Name, schema.Types[j].Name) < 0
	})
	sort.Slice(schema.Directives, func(i, j int) bool {
		return strings.Compare(schema.Directives[i].Name, schema.Directives[j].Name) < 0
	})
	*s = (Schema(schema))

	// Make a memo
	memoizeTypes := make(map[string]*Types, len(s.Types))
	mid := len(s.Types) / 2
	for i, j := 0, len(schema.Types)-1; i < mid; i, j = i+1, j-1 {
		memoizeTypes[schema.Types[i].Name] = schema.Types[i]
		memoizeTypes[schema.Types[j].Name] = schema.Types[j]
	}

	s.memoizeTypes = memoizeTypes

	return nil
}

func (s *Schema) String() string {
	b := pool.GetBuilder()
	defer pool.PutBuilder(b)

	for _, t := range s.Types {
		b.WriteString(t.String())
		b.WriteString("\n\n")
	}
	for _, d := range s.Directives {
		b.WriteString(d.String())
		b.WriteString("\n")
	}
	return b.String()
}

type OperationType struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Kind        string `json:"kind"`
}

type Types struct {
	Kind          TypeKind      `json:"kind"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Fields        []*Field      `json:"fields"`
	InputFields   []*InputValue `json:"inputFields"`
	Interfaces    []*Type       `json:"interfaces"`
	EnumValues    []*EnumValue  `json:"enumValues"`
	PossibleTypes []*Type       `json:"possibleTypes"`
}

var _ json.Unmarshaler = (*Types)(nil)

func (t *Types) UnmarshalJSON(data []byte) error {
	// HACK(codehex): To avoid stackoverflow, I assigned to tmp type.
	type HookType Types

	var types HookType
	if err := json.Unmarshal(data, &types); err != nil {
		return err
	}

	sort.Slice(types.Fields, func(i, j int) bool {
		return strings.Compare(types.Fields[i].Name, types.Fields[j].Name) < 0
	})
	sort.Slice(types.InputFields, func(i, j int) bool {
		return strings.Compare(types.InputFields[i].Name, types.InputFields[j].Name) < 0
	})
	sort.Slice(types.Interfaces, func(i, j int) bool {
		return strings.Compare(types.Interfaces[i].UnderlyingName(), types.Interfaces[j].UnderlyingName()) < 0
	})
	sort.Slice(types.EnumValues, func(i, j int) bool {
		return strings.Compare(types.EnumValues[i].Name, types.EnumValues[j].Name) < 0
	})
	sort.Slice(types.PossibleTypes, func(i, j int) bool {
		return strings.Compare(types.PossibleTypes[i].UnderlyingName(), types.PossibleTypes[j].UnderlyingName()) < 0
	})

	*t = Types(types)
	return nil
}

const indent = "  "

func (t *Types) IsOperationType() bool {
	if t == nil {
		return false
	}
	return t.Kind == OBJECT && (t.Name == "Query" || t.Name == "Mutation" || t.Name == "Subscription")
}

func (t *Types) String() string {
	if t == nil {
		return ""
	}
	b := pool.GetBuilder()
	defer pool.PutBuilder(b)

	switch t.Kind {
	case SCALAR:
		return "scalar " + t.Name
	case OBJECT:
		b.WriteString("type ")
		b.WriteString(t.Name)
		if len(t.Interfaces) > 0 {
			b.WriteString(" implements ")
			implements := make([]string, len(t.Interfaces))
			for i, v := range t.Interfaces {
				implements[i] = v.String()
			}
			b.WriteString(strings.Join(implements, " & "))
		}
		b.WriteString(" {\n")
		for _, f := range t.Fields {
			b.WriteString(indent)
			b.WriteString(f.String())
			b.WriteString("\n")
		}
		b.WriteString("}")
		return b.String()
	case INTERFACE:
		b.WriteString("interface ")
		b.WriteString(t.Name)
		b.WriteString(" {\n")
		for _, f := range t.Fields {
			b.WriteString(indent)
			b.WriteString(f.String())
			b.WriteString("\n")
		}
		b.WriteString("}")
		return b.String()
	case UNION:
		b.WriteString("union ")
		b.WriteString(t.Name)
		b.WriteString(" = ")
		elems := make([]string, len(t.PossibleTypes))
		for i, v := range t.PossibleTypes {
			elems[i] = v.String()
		}
		b.WriteString(strings.Join(elems, " | "))
		return b.String()
	case ENUM:
		b.WriteString("enum ")
		b.WriteString(t.Name)
		b.WriteString(" {\n")
		for _, ev := range t.EnumValues {
			b.WriteString(indent)
			b.WriteString(ev.String())
			b.WriteString("\n")
		}
		b.WriteString("}")
		return b.String()
	case INPUT_OBJECT:
		b.WriteString("input ")
		b.WriteString(t.Name)
		b.WriteString(" {\n")
		for _, f := range t.InputFields {
			b.WriteString(indent)
			b.WriteString(f.String())
			b.WriteString("\n")
		}
		b.WriteString("}")
		return b.String()
	}
	return ""
}

type Type struct {
	Kind   TypeKind `json:"kind"`
	Name   *string  `json:"name"`
	OfType *Type    `json:"ofType"`
}

func (t *Type) UnderlyingName() string {
	if t == nil {
		return ""
	}
	switch t.Kind {
	case NON_NULL, LIST:
		return t.OfType.UnderlyingName()
	}
	if t.Name != nil {
		return *t.Name
	}
	return ""
}

func (t *Type) UnderlyingKind() TypeKind {
	if t == nil {
		return ""
	}
	switch t.Kind {
	case NON_NULL, LIST:
		return t.OfType.UnderlyingKind()
	}
	return t.Kind
}

func (t *Type) String() string {
	if t == nil {
		return ""
	}

	switch t.Kind {
	case NON_NULL:
		return t.OfType.String() + "!"
	case LIST:
		return "[" + t.OfType.String() + "]"
	default:
		return t.UnderlyingName()
	}
}

type Args []*Arg

// String concatenates the elements of its first argument to create a single element. The separator
func (a Args) String() string {
	const sep = ", "
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0].String()
	}

	b := pool.GetBuilder()
	defer pool.PutBuilder(b)

	b.WriteString(a[0].String())
	for _, s := range a[1:] {
		b.WriteString(sep)
		b.WriteString(s.String())
	}
	return b.String()
}

type Arg struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Type         *Type   `json:"type"`
	DefaultValue *string `json:"defaultValue"`
}

func (a *Arg) String() string {
	if a == nil {
		return ""
	}
	ret := a.Name + ": " + a.Type.String()
	if a.DefaultValue != nil {
		return ret + " = " + *a.DefaultValue
	}
	return ret
}

type Directives struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Locations   []string `json:"locations"`
	Args        Args     `json:"args"`
}

func (d *Directives) String() string {
	if d == nil {
		return ""
	}

	b := pool.GetBuilder()
	defer pool.PutBuilder(b)

	b.WriteString("directive @")
	b.WriteString(d.Name)
	b.WriteString("(")
	b.WriteString(d.Args.String())
	b.WriteString(") on ")
	b.WriteString(strings.Join(d.Locations, " | "))
	return b.String()
}

type Field struct {
	Args              Args   `json:"args"`
	DeprecationReason string `json:"deprecationReason"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	Name              string `json:"name"`
	Type              *Type  `json:"type"`
}

func (f *Field) String() string {
	if f == nil {
		return ""
	}

	b := pool.GetBuilder()
	defer pool.PutBuilder(b)

	if len(f.Args) == 0 {
		b.WriteString(f.Name)
		b.WriteString(": ")
		b.WriteString(f.Type.String())
		return b.String()
	}
	b.WriteString(f.Name)
	b.WriteString("(")
	b.WriteString(f.Args.String())
	b.WriteString("): ")
	b.WriteString(f.Type.String())
	return b.String()
}

type InputValue struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Type         *Type   `json:"type"`
	DefaultValue *string `json:"defaultValue"`
}

func (i *InputValue) String() string {
	if i == nil {
		return ""
	}

	b := pool.GetBuilder()
	defer pool.PutBuilder(b)

	b.WriteString(i.Name)
	b.WriteString(": ")
	b.WriteString(i.Type.String())
	return b.String()
}

type EnumValue struct {
	Name              string `json:"name"`
	Description       string `json:"description"`
	IsDeprecated      bool   `json:"isDeprecated"`
	DeprecationReason string `json:"deprecationReason"`
}

func (e *EnumValue) String() string {
	return e.Name
}
