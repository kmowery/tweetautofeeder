package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.Handle("/", http.FileServer(http.Dir("/usr/share/3amh/www")))
    http.ListenAndServe(":8080", r)
}
