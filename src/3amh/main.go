package main

import (
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.Handle("/", http.FileServer(http.Dir("/vagrant/www")))
    http.ListenAndServe(":8080", r)
}
