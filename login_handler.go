package main

import (
  //"fmt"
  "log"
  "net/http"
  "github.com/hoisie/mustache"
)

func loginHandler(s Services, w http.ResponseWriter, r *http.Request) {

  url,_ := s.router.Get("begin_login").URLPath()

  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"url":url.String()})
  w.Write([]byte(data))
  return
}


func beginLoginHandler(s Services, w http.ResponseWriter, r *http.Request) {

  // TODO: add error handling
  requestToken, url, _ := s.customer.GetRequestTokenAndUrl("http://127.0.0.1:8080/login/oauth_callback")

  session,_ := s.sessions.Get(r, "session-name")
  session.Values["requestToken"] = requestToken.Token
  session.Values["requestTokenSecret"] = requestToken.Secret
  err := session.Save(r,w)

  err = addRequestTokens(s, requestToken)
  if(err != nil) {
    log.Fatal(err)
  }

  http.Redirect(w, r, url, 302)
  return
}

func oauthCallbackHandler(s Services, w http.ResponseWriter, r *http.Request) {
  token,present := r.URL.Query()["oauth_token"]
  if(!present) {
    log.Fatal("couldn't find oauth_token")
  }
  rt,err := getRequestToken(s, token[0])
  if(err != nil) {
    log.Fatal(err)
  }

  oauth_verifier,present := r.URL.Query()["oauth_verifier"]
  if(!present) {
    log.Fatal("couldn't find oauth_verifier")
  }

  atoken, err := s.customer.AuthorizeToken(rt, oauth_verifier[0])
  if(err != nil) {
    log.Fatal(err)
  }

  updateAccessToken(s, rt, atoken)


  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"url":"not a url"})
  w.Write([]byte(data))

  return
}

