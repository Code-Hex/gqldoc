package markdown

import (
	_ "embed"
	"text/template"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

//go:embed scalars.md
var scalars string

type (
	Scalars struct {
		Scalars []*Scalar
	}

	Scalar struct {
		Name        string
		Description string
	}
)

func (m *Config) renderScalars(scalarTypes []*introspection.Types) error {
	scalarList := make([]*Scalar, 0, len(scalarTypes))
	for _, typ := range scalarTypes {
		scalarList = append(scalarList, &Scalar{
			Name:        typ.Name,
			Description: renderHTML(typ.Description),
		})
	}
	f, err := m.Create("scalars.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	t, err := template.New("scalars").Parse(scalars)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := t.Execute(f, &Scalars{
		Scalars: scalarList,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
