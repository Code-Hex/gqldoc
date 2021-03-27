# Queries

### About queries

{{ .Description }}

{{- range $f := .Fields }}

### {{ $f.Name }}

#### Type: [{{ $f.Return.Type }}]({{ $f.Return.TypeLink }})

{{ $f.Description }}

{{- $length := len $f.Args }} {{ if ne $length 0 }}

#### Arguments

| Name | Description |
|------|-------------|
{{- range $i, $a := $f.Args }}
| {{ $a.Name }} ([{{ $a.Type }}]({{ $a.TypeLink }})) | {{ $a.Description }} |
{{- end }}
{{- end }}

---

{{- end }}