package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func main() {
    var err error
    err = nil

    r := mux.NewRouter()
    r.Handle("/", http.FileServer(http.Dir("/usr/share/3amh/www")))
    r.HandleFunc("/blog", blogHandler )
    err = http.ListenAndServe(":8080", r)
    if err != nil {
      log.Fatal(err)
    }
}
