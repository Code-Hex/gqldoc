# API Reference

View reference documentation to learn about the data types available in your GraphQL API schema.

{{ range $i := .Items }}
- [{{ $i.Type }}]({{ $i.Link }})
{{- end }}