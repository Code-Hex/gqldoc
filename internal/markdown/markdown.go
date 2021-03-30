package markdown

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unsafe"

	"github.com/Code-Hex/gqldoc/internal/introspection"
	"github.com/pkg/errors"
	"github.com/russross/blackfriday/v2"
)

func renderHTML(input string) string {
	result := blackfriday.Run([]byte(input))
	return strings.TrimSpace(*(*string)(unsafe.Pointer(&result)))
}

type Config struct {
	BaseDir string
}

func (m *Config) MakePath(filename string) string {
	return filepath.Join(m.BaseDir, filename)
}

func (m *Config) Create(filename string) (*os.File, error) {
	dest := m.MakePath(filename)
	return os.Create(dest)
}

func (m *Config) MakeLinkFromType(typ *introspection.Type) (string, error) {
	if typ == nil {
		return "", errors.New("type is nil")
	}
	kind := typ.UnderlyingKind()
	switch kind {
	case introspection.SCALAR,
		introspection.OBJECT,
		introspection.INTERFACE,
		introspection.UNION,
		introspection.ENUM,
		introspection.INPUT_OBJECT:
		link := fmt.Sprintf(
			"%ss.md#%s",
			kind,
			typ.UnderlyingName(),
		)
		return strings.ToLower(link), nil
	}
	return "", errors.Errorf("unexpected kind: %q", kind)
}

func (m *Config) Render(s *introspection.Schema) error {
	readmeContent := &README{}

	if s.QueryType != nil {
		if err := m.renderQuery(s); err != nil {
			return errors.WithStack(err)
		}
		readmeContent.Items = append(readmeContent.Items, &READMEItem{
			Type: "Query",
			Link: m.MakePath("queries.md"),
		})
	}
	if s.MutationType != nil {
		if err := m.renderMutation(s); err != nil {
			return errors.WithStack(err)
		}
		readmeContent.Items = append(readmeContent.Items, &READMEItem{
			Type: "Mutations",
			Link: m.MakePath("mutations.md"),
		})
	}
	if len(s.Types) > 0 {
		types := make(map[introspection.TypeKind][]*introspection.Types, len(s.Types))
		for _, typ := range s.Types {
			types[typ.Kind] = append(types[typ.Kind], typ)
		}
		if typs := types[introspection.OBJECT]; len(typs) > 0 {
			if err := m.renderObjects(typs); err != nil {
				return errors.WithStack(err)
			}
			readmeContent.Items = append(readmeContent.Items, &READMEItem{
				Type: "Objects",
				Link: m.MakePath("objects.md"),
			})
		}
		if typs := types[introspection.INTERFACE]; len(typs) > 0 {
			if err := m.renderInterfaces(typs); err != nil {
				return errors.WithStack(err)
			}
			readmeContent.Items = append(readmeContent.Items, &READMEItem{
				Type: "Interfaces",
				Link: m.MakePath("interfaces.md"),
			})
		}
		if typs := types[introspection.ENUM]; len(typs) > 0 {
			if err := m.renderEnums(typs); err != nil {
				return errors.WithStack(err)
			}
			readmeContent.Items = append(readmeContent.Items, &READMEItem{
				Type: "Enums",
				Link: m.MakePath("enums.md"),
			})
		}
		if typs := types[introspection.UNION]; len(typs) > 0 {
			if err := m.renderUnions(typs); err != nil {
				return errors.WithStack(err)
			}
			readmeContent.Items = append(readmeContent.Items, &READMEItem{
				Type: "Unions",
				Link: m.MakePath("unions.md"),
			})
		}
		if typs := types[introspection.INPUT_OBJECT]; len(typs) > 0 {
			if err := m.renderInputObjects(typs); err != nil {
				return errors.WithStack(err)
			}
			readmeContent.Items = append(readmeContent.Items, &READMEItem{
				Type: "Input objects",
				Link: m.MakePath("input_objects.md"),
			})
		}
		if typs := types[introspection.SCALAR]; len(typs) > 0 {
			if err := m.renderScalars(typs); err != nil {
				return errors.WithStack(err)
			}
			readmeContent.Items = append(readmeContent.Items, &READMEItem{
				Type: "Scalars",
				Link: m.MakePath("scalars.md"),
			})
		}
	}

	if err := m.renderREADME(readmeContent); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

//go:embed README.md
var readme string

type README struct {
	Items []*READMEItem
}

type READMEItem struct {
	Type string
	Link string
}

func (m *Config) renderREADME(r *README) error {
	f, err := m.Create("README.md")
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	t, err := template.New("README").Parse(readme)
	if err != nil {
		return errors.WithStack(err)
	}

	if err := t.Execute(f, r); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
