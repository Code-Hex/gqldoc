# Scalars

### About scalars

[Scalars](https://graphql.github.io/graphql-spec/June2018/#sec-Scalars) are primitive values: `Int`, `Float`, `String`, `Boolean`, or `ID`.

When calling the GraphQL API, you must specify nested subfields until you return only scalars.

{{- range $s := .Scalars }}

### {{ $s.Name }}

{{ $s.Description }}

---

{{- end }}