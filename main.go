package main

import (
	"fmt"
	"net/http"
	"os"

	"bitbucket.org/ww/goautoneg"
)

func main() {
	addr := os.Getenv("SYSTEMD_STATE_ADDR")

	if "" == addr {
		addr = ":80"
	}

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Print(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	contentType := goautoneg.Negotiate(r.Header.Get("Accept"), []string{"text/html", "application/json", "text/plain"})

	// Explicitly set the Content-Type header on non-HEAD requests
	// if the request "application/json". This is because
	// http.DetectContentType() is not able to detect it.
	if "application/json" == contentType && r.Method != http.MethodHead {
		w.Header().Set("Content-Type", contentType)
	}

	status, stateFormat := getStatusAndFormat()

	w.WriteHeader(status)

	// If this is a HEAD request we will not write a response body
	if r.Method == http.MethodHead {
		return
	}

	writeBody(w, contentType, stateFormat)
}
