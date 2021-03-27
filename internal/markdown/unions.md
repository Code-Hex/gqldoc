# Unions

### About unions

A [union](https://graphql.github.io/graphql-spec/June2018/#sec-Unions) is a type of object representing many objects.

{{- range $u := .Unions }}

### {{ $u.Name }}

{{ $u.Description }}

### Possible types

{{ range $v := $u.PossibleTypes }}
- [{{ $v.Name }}]({{ $v.Link }})
{{- end }}

---

{{- end }}