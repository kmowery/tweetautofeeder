package main

import (
  "net/http"
  "fmt"
  "github.com/hoisie/mustache"
)

func loginHandler(s Services, w http.ResponseWriter, r *http.Request) {

  // TODO: add error handling
  requestToken, url, _ := s.customer.GetRequestTokenAndUrl("http://127.0.0.1/login/oauth_redirect")

  session,_ := s.sessions.Get(r, "session-name")
  session.Values["requestToken"] = requestToken.Token
  session.Values["requestTokenSecret"] = requestToken.Secret
  err := session.Save(r,w)

  if(err != nil) {
    fmt.Println(err)
  }

  //http.Redirect(w, r, url, 302)
  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"url":url})
  w.Write([]byte(data))
  return
}

