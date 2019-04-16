package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/golang/gddo/httputil"
)

func main() {
	healthcheck := flag.Bool("healthcheck", false, "Do health check")
	flag.Parse()

	if *healthcheck {
		doHealthcheck()
	}

	addr := getAddr()

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Print(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	contentType := httputil.NegotiateContentType(r, []string{"text/html", "application/json", "text/plain"}, "text/plain")

	// Explicitly set the Content-Type header on non-HEAD requests
	// if the request "application/json". This is because
	// http.DetectContentType() is not able to detect it.
	if "application/json" == contentType && r.Method != http.MethodHead {
		w.Header().Set("Content-Type", contentType)
	}

	status, stateFormat := getStatusAndFormat(State())

	w.WriteHeader(status)

	// If this is a HEAD request we will not write a response body
	if r.Method == http.MethodHead {
		return
	}

	writeBody(w, contentType, stateFormat)
}

func getAddr() string {
	addr, present := os.LookupEnv("SYSTEMD_STATE_ADDR")

	if !present {
		addr = ":80"
	}

	return addr
}

// We do a health check simply by checking if we can get a systemd
// status or not. Notice the health check is not to check whether
// systemd is healthy or not but to check if this monitoring software
// itself is healthy.
func doHealthcheck() {
	_, err := State()

	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
