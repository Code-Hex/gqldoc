package markdown

import (
	_ "embed"
	"text/template"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

//go:embed enums.md
var enums string

type (
	Enums struct {
		Enums []*Enum
	}

	Enum struct {
		Name        string
		Description string
		Values      []*EnumValue
	}

	EnumValue struct {
		Name        string
		Description string
	}
)

func (m *Config) renderEnums(enumTypes []*introspection.Types) error {
	enumList := make([]*Enum, 0, len(enumTypes))
	for _, typ := range enumTypes {
		values := make([]*EnumValue, 0, len(typ.EnumValues))
		for _, v := range typ.EnumValues {
			values = append(values, &EnumValue{
				Name:        v.Name,
				Description: renderHTML(v.Description),
			})
		}
		enumList = append(enumList, &Enum{
			Name:        typ.Name,
			Description: renderHTML(typ.Description),
			Values:      values,
		})
	}
	f, err := m.Create("enums.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	t, err := template.New("enums").Parse(enums)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := t.Execute(f, &Enums{
		Enums: enumList,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
