package markdown

import (
	_ "embed"
	"strings"
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
			implementedBy = append(implementedBy, &ImplementedBy{
				Type:     i.String(),
				TypeLink: "http://example.com",
			})
		}
		fields := make([]*InterfaceField, 0, len(typ.Fields))
		for _, f := range typ.Fields {
			var args []*InterfaceFieldArg
			if len(f.Args) > 0 {
				args = make([]*InterfaceFieldArg, 0, len(f.Args))
				for _, arg := range f.Args {
					args = append(args, &InterfaceFieldArg{
						Name:        arg.Name,
						Type:        arg.Type.String(),
						TypeLink:    "http://example.com",
						Description: renderHTML(arg.Description),
					})
				}
			}
			fields = append(fields, &InterfaceField{
				Name:        f.Name,
				Type:        f.Type.String(),
				TypeLink:    "http://example.com",
				Description: strings.ReplaceAll(f.Description, "\n", " "),
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
