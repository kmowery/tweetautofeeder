package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/mrjones/oauth"
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

    c := oauth.NewConsumer(
      API_CONSUMER_KEY,
      API_CONSUMER_SECRET,
      oauth.ServiceProvider{
        RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
        AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
        AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
      })

    c.Debug(true)


    r := mux.NewRouter()
    r.HandleFunc("/", makeIndexHandler(*r) ).Name("root")
    r.Handle(    "/login", NewLoginHandler(c)  ).Name("login")
    r.HandleFunc("/debug", debugHandler ).Name("debug")
    r.HandleFunc("/blog", blogHandler ).Name("blog")

    r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/usr/share/tweetautofeeder/www"))))

    err = http.ListenAndServe(":8080", r)
    if err != nil {
      log.Fatal(err)
    }
}
