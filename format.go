package main

import (
	"encoding/json"
	htmlTemplate "html/template"
	"io"
	textTemplate "text/template"
)

// StateFormat is the data to format in the response.
type StateFormat struct {
	Label string `json:"label"`
	State string `json:"state"`
}

func writeBody(w io.Writer, contentType string, stateFormat StateFormat) {
	switch contentType {
	case "text/plain":
		writeTextBody(w, stateFormat)

	case "application/json":
		writeJSONBody(w, stateFormat)

	case "text/html":
		writeHTMLBody(w, stateFormat)

	default:
		writeHTMLBody(w, stateFormat)
	}
}

func writeHTMLBody(w io.Writer, stateFormat StateFormat) {
	tpl := `<!DOCTYPE html>
<html>
	<head>
		<title>{{.Label}}: {{.State}}</title>
	</head>
	<body>
		<h2>{{.Label}}</h2>
		<h1 style="text-transform: capitalize;">{{.State}}<h1>
	</body>
</html>`

	t, _ := htmlTemplate.New("systemd-state").Parse(tpl)
	_ = t.Execute(w, stateFormat)
}

func writeTextBody(w io.Writer, stateFormat StateFormat) {
	tpl := `{{.Label}}: {{.State}}`

	t, _ := textTemplate.New("systemd-state").Parse(tpl)
	_ = t.Execute(w, stateFormat)
}

func writeJSONBody(w io.Writer, stateFormat StateFormat) {
	_ = json.NewEncoder(w).Encode(stateFormat)
}
