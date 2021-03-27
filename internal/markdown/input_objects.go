package markdown

import (
	_ "embed"
	"text/template"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

//go:embed input_objects.md
var inputObjects string

type (
	InputObjects struct {
		InputObjects []*InputObject
	}

	InputObject struct {
		Name        string
		Description string
		InputFields []*InputObjectField
	}

	InputObjectField struct {
		Name        string
		Type        string
		TypeLink    string
		Description string
	}
)

func (m *Config) renderInputObjects(inputObjectTypes []*introspection.Types) error {
	inputObjectList := make([]*InputObject, 0, len(inputObjectTypes))
	for _, typ := range inputObjectTypes {

		fields := make([]*InputObjectField, 0, len(typ.InputFields))
		for _, f := range typ.InputFields {
			fields = append(fields, &InputObjectField{
				Name:        f.Name,
				Type:        f.Type.String(),
				TypeLink:    "http://example.com",
				Description: renderHTML(f.Description),
			})
		}
		inputObjectList = append(inputObjectList, &InputObject{
			Name:        typ.Name,
			Description: renderHTML(typ.Description),
			InputFields: fields,
		})
	}

	f, err := m.Create("input_objects.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	t, err := template.New("input_objects").Parse(inputObjects)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := t.Execute(f, &InputObjects{
		InputObjects: inputObjectList,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
