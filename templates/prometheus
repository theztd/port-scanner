{{ range . }}
# HELP portscanner_opened_ports Information about opened ports
# TYPE portscanner_opened_ports gauge
portscanner_opened_ports{{"{"}}address="{{ .Host }}"{{"}"}} {{ .OpenCount }}
# HELP portscanner_closed_ports Information about closed ports
# TYPE portscanner_closed_ports gauge
portscanner_closed_ports{{"{"}}address="{{ .Host }}"{{"}"}} {{ .CloseCount }}
{{ end }}