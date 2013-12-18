package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/hoisie/mustache"
)

func templateHandler(w http.ResponseWriter, r *http.Request) {
  data := mustache.RenderFile("/usr/share/3amh/templates/template.must", map[string]string{"thing":"places"})
  w.Write([]byte(data))
  return
}

func main() {
    var err error
    err = nil

    r := mux.NewRouter()
    r.Handle("/", http.FileServer(http.Dir("/usr/share/3amh/www")))
    r.HandleFunc("/template", templateHandler)
    err = http.ListenAndServe(":8080", r)
    if err != nil {
      log.Fatal(err)
    }
}
