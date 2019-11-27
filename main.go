package main

import (
	"fmt"
	"net/http"

	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
)

func main() {
	port := "6060"
	dir := "public"

	fmt.Printf("Running HTTP server in port: %s", port)
	http.HandleFunc("/markdown", generateMarkdown)
	http.Handle("/", http.FileServer(http.Dir(dir)))
	http.ListenAndServe(":"+port, nil)
}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	unsafe := blackfriday.Run([]byte(r.FormValue("body")))
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	rw.Write(html)
}
