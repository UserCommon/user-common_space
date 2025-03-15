package main

import (
  "fmt"
  "log"
  "net/http"
  "usercommon_space/handlers"
)

func main() {
  mux := http.NewServeMux()
  mux.Handle("/", &handlers.Index{})

  s := &http.Server{
    Addr: "127.0.0.1:8080",
    Handler: mux,
  }

  log.Fatal(s.ListenAndServe())

  fmt.Printf("Hello, world!")
}

