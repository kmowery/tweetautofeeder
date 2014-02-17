package main

import (
  "net/http"
  //"fmt"
  "github.com/hoisie/mustache"
  //"github.com/mrjones/oauth"
  //"github.com/gorilla/sessions"
)

func loginHandler(s Services, w http.ResponseWriter, r *http.Request) {

  // TODO: add error handling
  _, url, _ := s.customer.GetRequestTokenAndUrl("http://127.0.0.1/login/oauth_redirect")

  //http.Redirect(w, r, url, 302)
  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"url":url})
  w.Write([]byte(data))
  return
}

