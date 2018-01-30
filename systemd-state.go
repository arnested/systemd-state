package main

import (
	"fmt"
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
	fmt.Fprintf(w, "<h1 style=\"text-transform: capitalize;\">%s<h1>", state)
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
