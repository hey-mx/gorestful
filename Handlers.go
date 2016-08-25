package main

import (
	"fmt"
	"html"
	"net/http"
)

//Index Handler for root
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
