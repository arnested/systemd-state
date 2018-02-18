package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	state, err := State()
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error %s!", err)
		return
	}

	if state.IsRunning() {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	// If this is a HEAD request we will not write a response body
	if r.Method == http.MethodHead {
		return
	}

	tpl := `<!DOCTYPE html>
<html>
	<head>
		<title>Systemd state: {{.}}</title>
	</head>
	<body>
		<h1 style="text-transform: capitalize;">{{.}}<h1>
	</body>
</html>`

	t, _ := template.New("systemd-state").Parse(tpl)
	_ = t.Execute(w, state.String())
}
