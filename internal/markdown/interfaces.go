package markdown

import (
	_ "embed"
	"text/template"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

//go:embed interfaces.md
var interfaces string

type (
	Interfaces struct {
		Interfaces []*Interface
	}

	Interface struct {
		Name        string
		Description string
		Implemented []*ImplementedBy
		Fields      []*InterfaceField
	}

	ImplementedBy struct {
		Type     string
		TypeLink string
	}

	InterfaceField struct {
		Name        string
		Type        string
		TypeLink    string
		Description string
		Args        []*InterfaceFieldArg
	}

	InterfaceFieldArg struct {
		Name        string
		Type        string
		TypeLink    string
		Description string
	}
)

func (m *Config) renderInterfaces(interfaceTypes []*introspection.Types) error {
	interfaceList := make([]*Interface, 0, len(interfaceTypes))
	for _, typ := range interfaceTypes {
		implementedBy := make([]*ImplementedBy, 0, len(typ.Interfaces))
		for _, i := range typ.PossibleTypes {
			link, err := m.MakeLinkFromType(i)
			if err != nil {
				return errors.Wrapf(err, "possible type %q has caused error", i.UnderlyingName())
			}
			implementedBy = append(implementedBy, &ImplementedBy{
				Type:     i.String(),
				TypeLink: link,
			})
		}
		fields := make([]*InterfaceField, 0, len(typ.Fields))
		for _, f := range typ.Fields {
			var args []*InterfaceFieldArg
			if len(f.Args) > 0 {
				args = make([]*InterfaceFieldArg, 0, len(f.Args))
				for _, arg := range f.Args {
					link, err := m.MakeLinkFromType(arg.Type)
					if err != nil {
						return errors.Wrapf(err, "argument type %q has caused error", arg.Type.UnderlyingName())
					}
					args = append(args, &InterfaceFieldArg{
						Name:        arg.Name,
						Type:        arg.Type.String(),
						TypeLink:    link,
						Description: renderHTML(arg.Description),
					})
				}
			}
			link, err := m.MakeLinkFromType(f.Type)
			if err != nil {
				return errors.Wrapf(err, "field type %q has caused error", f.Type.UnderlyingName())
			}
			fields = append(fields, &InterfaceField{
				Name:        f.Name,
				Type:        f.Type.String(),
				TypeLink:    link,
				Description: renderHTML(f.Description),
				Args:        args,
			})
		}
		interfaceList = append(interfaceList, &Interface{
			Name:        typ.Name,
			Description: renderHTML(typ.Description),
			Implemented: implementedBy,
			Fields:      fields,
		})
	}

	f, err := m.Create("interfaces.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	t, err := template.New("interfaces").Parse(interfaces)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := t.Execute(f, &Interfaces{
		Interfaces: interfaceList,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
