package markdown

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

func (m *Config) Create(filename string) (*os.File, error) {
	dest := filepath.Join(m.BaseDir, filename)
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
	if s.QueryType != nil {
		if err := m.renderQuery(s); err != nil {
			return errors.WithStack(err)
		}
	}
	if s.MutationType != nil {
		if err := m.renderMutation(s); err != nil {
			return errors.WithStack(err)
		}
	}
	if len(s.Types) > 0 {
		types := make(map[introspection.TypeKind][]*introspection.Types, len(s.Types))
		for _, typ := range s.Types {
			types[typ.Kind] = append(types[typ.Kind], typ)
		}
		if err := m.renderObjects(types[introspection.OBJECT]); err != nil {
			return errors.WithStack(err)
		}
		if err := m.renderInterfaces(types[introspection.INTERFACE]); err != nil {
			return errors.WithStack(err)
		}
		if err := m.renderEnums(types[introspection.ENUM]); err != nil {
			return errors.WithStack(err)
		}
		if err := m.renderUnions(types[introspection.UNION]); err != nil {
			return errors.WithStack(err)
		}
		if err := m.renderInputObjects(types[introspection.INPUT_OBJECT]); err != nil {
			return errors.WithStack(err)
		}
		if err := m.renderScalars(types[introspection.SCALAR]); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
