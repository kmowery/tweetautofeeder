package main

import (
  "encoding/base64"
  "crypto/rand"
  //"fmt"
  "log"
  "net/http"
  "regexp"
  "github.com/hoisie/mustache"
)

func generateRandomString() string {
  var r *regexp.Regexp;
  var err error;

  r,err = regexp.Compile( `[^\w]` )

  b := make([]byte, 32)
  _, err = rand.Read(b)
  if(err != nil ) {
    // not sure what to do here...
    log.Fatal("couldn't read random bytes...")
  }
  return r.ReplaceAllString(base64.StdEncoding.EncodeToString(b), "")
}



func loginHandler(s Services, w http.ResponseWriter, r *http.Request) {

  url,_ := s.router.Get("login_begin").URLPath()

  data := mustache.RenderFileInLayout(
    "/usr/share/tweetautofeeder/templates/main_page.must",
    "/usr/share/tweetautofeeder/templates/layout.must",
    map[string]string{"url":url.String()})
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

  log.Println("after authorize token")

  sessioncookie := generateRandomString()

  updateAccessToken(s, rt, atoken, sessioncookie)

  log.Println("updated access token")

  session, _ := s.sessions.Get(r, "login")
  session.Values["sessioncookie"] = sessioncookie
  err = session.Save(r, w)
  if(err != nil) {
    log.Fatal(err)
  }

  log.Println("redirecting...")

  redirect_url,err := s.router.Get("pending").URLPath()
  if err != nil {
    log.Fatal(err)
  }
  http.Redirect(w, r, redirect_url.String(), 307)
  return
}

