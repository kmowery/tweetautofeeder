package main

import (
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/gorilla/sessions"
  "github.com/mrjones/oauth"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
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



type Services struct {
  customer *oauth.Consumer
  sessions *sessions.CookieStore
  storage *sql.DB
  router *mux.Router
}
type ServicesHandler func(s Services, w http.ResponseWriter, r *http.Request)

func makeServicesHandler( s Services, sh ServicesHandler ) http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request) {
    sh(s, w, r)
  }
}

type User struct {
  sessionCookie string
  userId string
  screenName string
  atoken oauth.AccessToken
}


func main() {
    var err error
    err = nil

    services := Services{}

    services.customer = oauth.NewConsumer(
      API_CONSUMER_KEY,
      API_CONSUMER_SECRET,
      oauth.ServiceProvider{
        RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
        AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
        AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
      })
    services.customer.Debug(true)

    // TODO parameterize
    services.sessions = sessions.NewCookieStore([]byte(COOKIE_KEY))

    // TODO paramterize
    services.storage, err = sql.Open("sqlite3", "/home/vagrant/storage.sqlite")
    if(err != nil) {
      log.Fatal(err)
    }
    defer services.storage.Close()

    r := mux.NewRouter()
    services.router = r
    r.HandleFunc("/",                     makeIndexHandler(*r) ).Name("root")
    r.HandleFunc("/login",                makeServicesHandler(services, loginHandler)  ).Name("login")
    r.HandleFunc("/login/begin_login",    makeServicesHandler(services, beginLoginHandler)  ).Name("begin_login")
    r.HandleFunc("/login/oauth_callback", makeServicesHandler(services, oauthCallbackHandler)  ).Name("oauth_callback")
    r.HandleFunc("/list",                 makeServicesHandler(services, listHandler)  ).Name("list")
    r.HandleFunc("/debug", debugHandler ).Name("debug")
    r.HandleFunc("/blog", blogHandler ).Name("blog")
    // TODO: robots.txt, humans.txt

    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/usr/share/tweetautofeeder/www"))))
    r.PathPrefix("/css"    ).Handler(http.StripPrefix("/css/",    http.FileServer(http.Dir("/usr/share/tweetautofeeder/css"))))
    r.PathPrefix("/js"     ).Handler(http.StripPrefix("/js/",     http.FileServer(http.Dir("/usr/share/tweetautofeeder/js"))))

    err = http.ListenAndServe(":8080", r)
    if err != nil {
      log.Fatal(err)
    }
}
