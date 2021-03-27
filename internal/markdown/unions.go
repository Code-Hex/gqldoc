package markdown

import (
	_ "embed"
	"text/template"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

//go:embed unions.md
var unions string

type (
	Unions struct {
		Unions []*Union
	}

	Union struct {
		Name          string
		Description   string
		PossibleTypes []*UnionPossibleType
	}

	UnionPossibleType struct {
		Name string
		Link string
	}
)

func (m *Config) renderUnions(unionTypes []*introspection.Types) error {
	unionList := make([]*Union, 0, len(unionTypes))
	for _, typ := range unionTypes {
		possibleTypes := make([]*UnionPossibleType, 0, len(typ.PossibleTypes))
		for _, v := range typ.PossibleTypes {
			possibleTypes = append(possibleTypes, &UnionPossibleType{
				Name: v.String(),
				Link: "http://example.com",
			})
		}
		unionList = append(unionList, &Union{
			Name:          typ.Name,
			Description:   renderHTML(typ.Description),
			PossibleTypes: possibleTypes,
		})
	}
	f, err := m.Create("unions.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	t, err := template.New("unions").Parse(unions)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := t.Execute(f, &Unions{
		Unions: unionList,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
