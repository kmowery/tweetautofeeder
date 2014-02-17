package main

import (
  "log"
  "net/http"
)

func loginHandler(s Services, w http.ResponseWriter, r *http.Request) {

  // TODO: add error handling
  requestToken, url, _ := s.customer.GetRequestTokenAndUrl("http://127.0.0.1:8080//login/oauth_redirect")

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

