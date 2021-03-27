# Objects

### About objects

[Objects](https://graphql.github.io/graphql-spec/June2018/#sec-Objects) in GraphQL represent the resources you can access. An object can contain a list of fields, which are specifically typed.

{{- range $o := .Objects }}

### {{ $o.Name }}

{{ $o.Description }}

{{- $length := len $o.Implements }} {{ if ne $length 0 }}

#### Implements

{{ range $i, $a := $o.Implements }}
- [{{ $a.Type }}]({{ $a.TypeLink }})
{{- end }}
{{- end }}

{{- $length := len $o.Fields }} {{ if ne $length 0 }}

#### Fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
{{- range $f := $o.Fields }}
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