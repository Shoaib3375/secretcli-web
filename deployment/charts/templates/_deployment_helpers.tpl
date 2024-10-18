{{- define "secretcli-web.labels" -}}
labels:
  chart: {{ template "secretcli-web.chart" . }}
  release: {{ .Release.Name }}
  heritage: {{ .Release.Service }}
{{- end -}}

{{- define "secretcli-web.metadata" -}}
metadata:
  labels:
    release: {{ .Release.Name }}
{{- end -}}

{{- define "secretcli-web.selector" -}}
selector:
  matchLabels:
    release: {{ .Release.Name }}
{{- end -}}

{{- define "secretcli-web.envFrom" -}}
envFrom:
- configMapRef:
    name: {{ template "secretcli-web.name" . }}-configmap
{{- end -}}