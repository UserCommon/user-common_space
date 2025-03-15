package handlers

import (
	"fmt"
	"net/http"
)

type Index struct { }

func (h Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>This is main page!</h1>")
}
