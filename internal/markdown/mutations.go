package markdown

import (
	_ "embed"
	"text/template"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
)

//go:embed mutations.md
var mutations string

type (
	Mutation struct {
		Description string
		Fields      []*MutationField
	}

	MutationField struct {
		Name        string
		Description string
		Args        []*MutationFieldArg
		Returns     []*MutationFieldReturn
	}

	MutationFieldArg struct {
		Name     string
		Type     string
		TypeLink string
	}

	MutationFieldReturn struct {
		Name        string
		Description string
		Type        string
		TypeLink    string
	}
)

func (m *Config) renderMutation(s *introspection.Schema) error {
	t, err := template.New("mutations").Parse(mutations)
	if err != nil {
		return errors.WithStack(err)
	}

	types, err := s.FindTypes(s.MutationType.Name)
	if err != nil {
		return errors.WithStack(err)
	}

	mf := make([]*MutationField, 0, len(types.Fields))
	for _, f := range types.Fields {
		mfa := make([]*MutationFieldArg, 0, len(f.Args))
		for _, arg := range f.Args {
			link, err := m.MakeLinkFromType(arg.Type)
			if err != nil {
				return errors.Wrapf(err, "argument type %q has caused error", arg.Type.UnderlyingName())
			}
			mfa = append(mfa, &MutationFieldArg{
				Name:     arg.Name,
				Type:     arg.Type.String(),
				TypeLink: link,
			})
		}

		retType, err := s.FindTypes(f.Type.UnderlyingName())
		if err != nil {
			return errors.Wrapf(err, "failed to find types: %q", f.Name)
		}
		rfs := make([]*MutationFieldReturn, 0, len(retType.Fields))
		for _, rf := range retType.Fields {
			link, err := m.MakeLinkFromType(rf.Type)
			if err != nil {
				return errors.Wrapf(err, "return type %q has caused error", rf.Type.UnderlyingName())
			}
			rfs = append(rfs, &MutationFieldReturn{
				Name:        rf.Name,
				Description: renderHTML(rf.Description),
				Type:        rf.Type.String(),
				TypeLink:    link,
			})
		}

		mf = append(mf, &MutationField{
			Name:        f.Name,
			Description: renderHTML(f.Description),
			Args:        mfa,
			Returns:     rfs,
		})
	}

	f, err := m.Create("mutations.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	if err := t.Execute(f, &Mutation{
		Description: types.Description,
		Fields:      mf,
	}); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
