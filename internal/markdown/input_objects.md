# Input objects

### About input objects

[Input objects](https://graphql.github.io/graphql-spec/June2018/#sec-Input-Objects) can be described as "composable objects" because they include a set of input fields that define the object.


{{- range $i := .InputObjects }}

### {{ $i.Name }}

{{ $i.Description }}


#### Input fields

<table>
  <tr>
    <th>Name</th>
    <th>Description</th>
  </tr>
{{- range $f := $i.InputFields }}
  <tr>
    <td><strong>{{ $f.Name }}</strong> (<a href="{{ $f.TypeLink }}">{{ $f.Type }}</a>)</td>
    <td>{{ $f.Description }}</td>
  </tr>
{{- end }}
</table>

---

{{- end }}