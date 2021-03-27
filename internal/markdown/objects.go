package markdown

import (
	_ "embed"
	"strings"
	"text/template"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

//go:embed objects.md
var objects string

type (
	Objects struct {
		Objects []*Object
	}

	Object struct {
		Name        string
		Description string
		Implements  []*ObjectImplements
		Fields      []*ObjectField
	}

	ObjectImplements struct {
		Type     string
		TypeLink string
	}

	ObjectField struct {
		Name        string
		Type        string
		TypeLink    string
		Description string
		Args        []*ObjectFieldArg
	}

	ObjectFieldArg struct {
		Name        string
		Type        string
		TypeLink    string
		Description string
	}
)

func (m *Config) renderObjects(objectTypes []*introspection.Types) error {
	objectList := make([]*Object, 0, len(objectTypes))
	for _, typ := range objectTypes {
		if typ.IsOperationType() {
			continue
		}
		implements := make([]*ObjectImplements, 0, len(typ.Interfaces))
		for _, i := range typ.Interfaces {
			implements = append(implements, &ObjectImplements{
				Type:     i.String(),
				TypeLink: "http://example.com",
			})
		}
		fields := make([]*ObjectField, 0, len(typ.Fields))
		for _, f := range typ.Fields {
			var args []*ObjectFieldArg
			if len(f.Args) > 0 {
				args = make([]*ObjectFieldArg, 0, len(f.Args))
				for _, arg := range f.Args {
					args = append(args, &ObjectFieldArg{
						Name:        arg.Name,
						Type:        arg.Type.String(),
						TypeLink:    "http://example.com",
						Description: renderHTML(arg.Description),
					})
				}
			}
			fields = append(fields, &ObjectField{
				Name:        f.Name,
				Type:        f.Type.String(),
				TypeLink:    "http://example.com",
				Description: strings.ReplaceAll(f.Description, "\n", " "),
				Args:        args,
			})
		}
		objectList = append(objectList, &Object{
			Name:        typ.Name,
			Description: strings.ReplaceAll(typ.Description, "\n", " "),
			Implements:  implements,
			Fields:      fields,
		})
	}

	f, err := m.Create("objects.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	t, err := template.New("objects").Parse(objects)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := t.Execute(f, &Objects{
		Objects: objectList,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
