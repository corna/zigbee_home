{{- range .AdditionalContext }}
{{- if (ne .Size 0) }}
{{ .Name }}:
  address: {{formatHex .Address}}
  end_address: {{ formatHex .EndAddress }}
  region: {{ if .Region }}{{.Region}}{{else}}flash_primary{{end}}
  size: {{ formatHex .Size }}
{{- end}}
{{- end}}
