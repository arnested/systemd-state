package main

import (
	"net/http"

	"bitbucket.org/ww/goautoneg"
)

func main() {
	http.HandleFunc("/", handler)
	_ = http.ListenAndServe(":80", nil)
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
