package markdown

import (
	_ "embed"
	"text/template"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

//go:embed queries.md
var queries string

type (
	Query struct {
		Description string
		Fields      []*QueryField
	}

	QueryField struct {
		Name        string
		Description string
		Args        []*QueryFieldArg
		Return      *QueryFieldReturn
	}

	QueryFieldArg struct {
		Name        string
		Description string
		Type        string
		TypeLink    string
	}

	QueryFieldReturn struct {
		Type     string
		TypeLink string
	}
)

func (m *Config) renderQuery(s *introspection.Schema) error {
	t, err := template.New("queries").Parse(queries)
	if err != nil {
		return errors.WithStack(err)
	}

	types, err := s.FindTypes(s.QueryType.Name)
	if err != nil {
		return errors.WithStack(err)
	}

	mf := make([]*QueryField, 0, len(types.Fields))
	for _, f := range types.Fields {
		mfa := make([]*QueryFieldArg, 0, len(f.Args))
		for _, arg := range f.Args {
			link, err := m.MakeLinkFromType(arg.Type)
			if err != nil {
				return errors.Wrapf(err, "argument %q has caused error", arg.Type.UnderlyingName())
			}
			mfa = append(mfa, &QueryFieldArg{
				Name:        arg.Name,
				Description: renderHTML(arg.Description),
				Type:        arg.Type.String(),
				TypeLink:    link,
			})
		}

		fieldTypeLink, err := m.MakeLinkFromType(f.Type)
		if err != nil {
			return errors.Wrapf(err, "field %q has caused error", f.Type.UnderlyingName())
		}
		mf = append(mf, &QueryField{
			Name:        f.Name,
			Description: f.Description,
			Args:        mfa,
			Return: &QueryFieldReturn{
				Type:     f.Type.String(),
				TypeLink: fieldTypeLink,
			},
		})
	}

	f, err := m.Create("queries.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	if err := t.Execute(f, &Query{
		Description: types.Description,
		Fields:      mf,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
