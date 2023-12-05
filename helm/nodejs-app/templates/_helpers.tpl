{{- define "helpers.IMAGE" -}}
{{- printf "%s:%s" .Values.image.repository (.Values.image.tag| default "latest" ) -}}
{{- end -}}
