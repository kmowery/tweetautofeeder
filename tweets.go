package main

import (
  "log"
  "net/http"
  "github.com/hoisie/mustache"
)

func listHandler(s Services, w http.ResponseWriter, r *http.Request) {
  session, err := s.sessions.Get(r, "login")
  if(err != nil) {
    log.Fatal(err)
  }

  sessioncookie,present := session.Values["sessioncookie"]
  if(!present) {
    // TODO redirect somewhere nice with an error message
    log.Println("cookie not existent, come back again")
  }

  user,err := getUser(s, sessioncookie.(string))
  if(err != nil) {
    log.Fatal(err)
  }
  if(user == nil) {
    // TODO redirect somewhere nice
    log.Println("user didn't exist")
  }


  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"url":"not a url"})
  w.Write([]byte(data))

  return
}


