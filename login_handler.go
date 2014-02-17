package main

import (
  "net/http"
  //"fmt"
  //"github.com/hoisie/mustache"
  "github.com/mrjones/oauth"
)

type LoginHandler struct { customer *oauth.Consumer }

func NewLoginHandler( customer *oauth.Consumer ) LoginHandler {
  lh := LoginHandler {}
  lh.customer = customer
  return lh
}

func (lh LoginHandler) ServeHTTP(
   w http.ResponseWriter,
   r *http.Request) {

  // TODO: add error handling
  _, url, _ := lh.customer.GetRequestTokenAndUrl("http://callback.com")

  http.Redirect(w, r, url, 307)
  return
}

