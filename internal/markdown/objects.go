package markdown

import (
	_ "embed"
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
			implementedTypeLink, err := m.MakeLinkFromType(i)
			if err != nil {
				return errors.Wrapf(err, "implemented type %q has caused error", i.UnderlyingName())
			}
			implements = append(implements, &ObjectImplements{
				Type:     i.String(),
				TypeLink: implementedTypeLink,
			})
		}
		fields := make([]*ObjectField, 0, len(typ.Fields))
		for _, f := range typ.Fields {
			var args []*ObjectFieldArg
			if len(f.Args) > 0 {
				args = make([]*ObjectFieldArg, 0, len(f.Args))
				for _, arg := range f.Args {
					argLink, err := m.MakeLinkFromType(arg.Type)
					if err != nil {
						return errors.Wrapf(err, "argument type %q has caused error", arg.Type.UnderlyingName())
					}
					args = append(args, &ObjectFieldArg{
						Name:        arg.Name,
						Type:        arg.Type.String(),
						TypeLink:    argLink,
						Description: renderHTML(arg.Description),
					})
				}
			}
			fieldLink, err := m.MakeLinkFromType(f.Type)
			if err != nil {
				return errors.Wrapf(err, "field type %q has caused error", f.Type.UnderlyingName())
			}
			fields = append(fields, &ObjectField{
				Name:        f.Name,
				Type:        f.Type.String(),
				TypeLink:    fieldLink,
				Description: renderHTML(f.Description),
				Args:        args,
			})
		}
		objectList = append(objectList, &Object{
			Name:        typ.Name,
			Description: renderHTML(typ.Description),
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
