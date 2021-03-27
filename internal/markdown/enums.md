# Enums

### About enums

[Enums](https://graphql.github.io/graphql-spec/June2018/#sec-Enums) represent possible sets of values for a field.

{{- range $e := .Enums }}

### {{ $e.Name }}

{{ $e.Description }}

<table>
  <tr>
    <th>Value</th>
    <th>Description</th>
  </tr>
{{- range $v := $e.Values }}
  <tr>
    <td><strong>{{ $v.Name }}</strong></td>
    <td>{{ $v.Description }}</td>
  </tr>
{{- end }}
</table>

---

{{- end }}