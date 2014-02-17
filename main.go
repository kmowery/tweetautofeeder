package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

func makeIndexHandler(router mux.Router) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    s,err := router.Get("login").URLPath()
    if err != nil {
      log.Fatal(err)
    }

    http.Redirect(w, r, s.String(), 307)
    return
  }
}

func main() {
    var err error
    err = nil

    r := mux.NewRouter()
    r.HandleFunc("/", makeIndexHandler(*r) ).Name("root")
    r.HandleFunc("/login", loginHandler ).Name("login")
    r.HandleFunc("/debug", debugHandler ).Name("debug")
    r.HandleFunc("/blog", blogHandler ).Name("blog")

    r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/usr/share/tweetautofeeder/www"))))

    err = http.ListenAndServe(":8080", r)
    if err != nil {
      log.Fatal(err)
    }
}
