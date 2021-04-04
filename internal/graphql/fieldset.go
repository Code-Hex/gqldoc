package graphql

import (
	"io"
	"sync"

	gql "github.com/99designs/gqlgen/graphql"
)

type FieldSet struct {
	fields  []CollectedField
	Values  []gql.Marshaler
	delayed []delayedResult
}

type delayedResult struct {
	i int
	f func() gql.Marshaler
}

func NewFieldSet(fields []CollectedField) *FieldSet {
	return &FieldSet{
		fields: fields,
		Values: make([]gql.Marshaler, len(fields)),
	}
}

func (m *FieldSet) Concurrently(i int, f func() gql.Marshaler) {
	m.delayed = append(m.delayed, delayedResult{i: i, f: f})
}

func (m *FieldSet) Dispatch() {
	if len(m.delayed) == 1 {
		// only one concurrent task, no need to spawn a goroutine or deal create waitgroups
		d := m.delayed[0]
		m.Values[d.i] = d.f()
	} else if len(m.delayed) > 1 {
		// more than one concurrent task, use the main goroutine to do one, only spawn goroutines for the others

		var wg sync.WaitGroup
		for _, d := range m.delayed[1:] {
			wg.Add(1)
			go func(d delayedResult) {
				m.Values[d.i] = d.f()
				wg.Done()
			}(d)
		}

		m.Values[m.delayed[0].i] = m.delayed[0].f()
		wg.Wait()
	}
}

var openBrace = []byte(`{`)
var closeBrace = []byte(`}`)
var colon = []byte(`:`)
var comma = []byte(`,`)

func (m *FieldSet) MarshalGQL(writer io.Writer) {
	writer.Write(openBrace)
	for i, field := range m.fields {
		if i != 0 {
			writer.Write(comma)
		}
		writeQuotedString(writer, field.Alias)
		writer.Write(colon)
		m.Values[i].MarshalGQL(writer)
	}
	writer.Write(closeBrace)
}

const encodeHex = "0123456789ABCDEF"

func writeQuotedString(w io.Writer, s string) {
	start := 0
	io.WriteString(w, `"`)

	for i, c := range s {
		if c < 0x20 || c == '\\' || c == '"' {
			io.WriteString(w, s[start:i])

			switch c {
			case '\t':
				io.WriteString(w, `\t`)
			case '\r':
				io.WriteString(w, `\r`)
			case '\n':
				io.WriteString(w, `\n`)
			case '\\':
				io.WriteString(w, `\\`)
			case '"':
				io.WriteString(w, `\"`)
			default:
				io.WriteString(w, `\u00`)
				w.Write([]byte{encodeHex[c>>4], encodeHex[c&0xf]})
			}

			start = i + 1
		}
	}

	io.WriteString(w, s[start:])
	io.WriteString(w, `"`)
}
