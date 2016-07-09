package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var (
	fListen   string
	fTemplate string
	fAPI      string
)

func init() {
	flag.StringVar(&fListen, "l", "localhost:9997", "The `address` to listen on")
	flag.StringVar(&fTemplate, "t", "", "The `directory` containing the UI code")
	flag.StringVar(&fAPI, "a", "http://localhost:9998", "The `address` of the Tracer query API")
}

func main() {
	flag.Parse()
	if fListen == "" || fTemplate == "" || fAPI == "" {
		flag.Usage()
		os.Exit(1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".") {
			http.ServeFile(w, r, filepath.Join(fTemplate, r.URL.Path))
			return
		}
		http.ServeFile(w, r, filepath.Join(fTemplate, "index.html"))
	})
	proxy, err := url.Parse(fAPI)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid API address:", err)
		os.Exit(1)
	}
	http.Handle("/api/", httputil.NewSingleHostReverseProxy(proxy))
	if err := http.ListenAndServe(fListen, nil); err != nil {
		fmt.Fprintln(os.Stderr, "Error starting HTTP server:", err)
		os.Exit(2)
	}
}
