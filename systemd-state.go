package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/coreos/go-systemd/dbus"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	code, state, err := state()
	if err != nil {
		w.WriteHeader(code)
		fmt.Fprintf(w, "Error %s!", err)
		return
	}

	w.WriteHeader(code)

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
	t.Execute(w, state)
}

func state() (int, string, error) {
	systemd, err := dbus.NewSystemdConnection()
	if err != nil {
		return http.StatusGone, "", err
	}
	defer systemd.Close()

	p, err := systemd.SystemState()
	if err != nil {
		return http.StatusServiceUnavailable, "", err
	}

	status := strings.Trim((&p.Value).String(), "\"")

	if "running" == status {
		return http.StatusOK, status, nil
	}

	return http.StatusInternalServerError, status, nil
}
