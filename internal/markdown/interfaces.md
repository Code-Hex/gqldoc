# Interfaces

### About interfaces

[Interfaces](https://graphql.github.io/graphql-spec/June2018/#sec-Interfaces) serve as parent objects from which other objects can inherit.

{{- range $i := .Interfaces }}

### {{ $i.Name }}

{{ $i.Description }}

{{- $length := len $i.Implemented }} {{ if ne $length 0 }}

#### Implemented by

{{ range $i, $a := $i.Implemented }}
- [{{ $a.Type }}]({{ $a.TypeLink }})
{{- end }}
{{- end }}

{{- $length := len $i.Fields }} {{ if ne $length 0 }}

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
{{- range $f := $i.Fields }}
  <tr>
    <td><strong>{{ $f.Name }}</strong> (<a href="{{ $f.TypeLink }}">{{ $f.Type }}</a>)</td>
  {{- $length := len $f.Args }} {{ if eq $length 0 }}
    <td>{{ $f.Description }}</td>
  {{- else }}
    <td>
      <p>{{ $f.Description }}</p>
      <table>
        <tr>
          <th><strong>Arguments</strong></th>
        </tr>
    {{- range $a := $f.Args }}
        <tr>
          <td>
            <p>{{ $a.Name }} (<a href="{{ $a.TypeLink }}">{{ $a.Type }}</a>)</p>
            <p>{{ $a.Description }}</p>
          </td>
        </tr>
    {{- end }}
      </table>
    </td>
  {{- end }}
  </tr>
{{- end }}
</table>
{{- end }}

---

{{- end }}