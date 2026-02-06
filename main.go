package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/elnormous/contenttype"
)

func main() {
	ctx := context.Background()

	healthcheck := flag.Bool("healthcheck", false, "Do health check")
	flag.Parse()

	if *healthcheck {
		doHealthcheck(ctx)
	}

	server := &http.Server{
		Addr:              getAddr(),
		ReadHeaderTimeout: 3 * time.Second,
	}

	http.HandleFunc("/", handler)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Print(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	supportedContentTypes := []contenttype.MediaType{
		contenttype.NewMediaType("text/plain"),
		contenttype.NewMediaType("application/json"),
		contenttype.NewMediaType("text/html"),
	}

	contentType, _, err := contenttype.GetAcceptableMediaType(r, supportedContentTypes)
	if err != nil {
		contentType = contenttype.NewMediaType("text/plain")
	}

	// Explicitly set the Content-Type header on non-HEAD requests
	// if the request "application/json". This is because
	// http.DetectContentType() is not able to detect it.
	if contentType.String() == "application/json" && r.Method != http.MethodHead {
		w.Header().Set("Content-Type", contentType.String())
	}

	status, stateFormat := getStatusAndFormat(State(r.Context()))

	w.WriteHeader(status)

	// If this is a HEAD request we will not write a response body
	if r.Method == http.MethodHead {
		return
	}

	writeBody(w, contentType.String(), stateFormat)
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
func doHealthcheck(ctx context.Context) {
	_, err := State(ctx)
	if err != nil {
		os.Exit(1)
	}

	os.Exit(0)
}
