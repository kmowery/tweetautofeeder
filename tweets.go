package main

import (
  "fmt"
  //"log"
  "net/http"
  "github.com/hoisie/mustache"
  "github.com/gorilla/mux"
)

type Tweet struct {
  id string
  tweet string
}
func (tw Tweet) Get() string {
  return tw.tweet
}
func (tw Tweet) ID() string {
  return tw.id
}


func postNewHandler(s Services, w http.ResponseWriter, r *http.Request) {
  user,err := getUser(s,w,r)
  if(err != nil) {
    // TODO don't leak errors
    renderError(w, fmt.Sprintf("%s", err))
    return
  }
  if(user == nil) {
    renderError(w, "You aren't logged in!")
    return
  }

  r.ParseForm()
  tweet,exists := r.Form["tweet"]
  if(!exists) {
     renderError(w, "Unknown error")
     return
  }

  addTweet(s,*user, tweet[0])

  redirect_url,err := s.router.Get("pending").URLPath()
  http.Redirect(w, r, redirect_url.String(), 302)
  return
}


func deleteHandler(s Services, w http.ResponseWriter, r *http.Request) {
  user,err := getUser(s,w,r)
  if(err != nil) {
    // TODO don't leak errors
    renderError(w, fmt.Sprintf("%s", err))
    return
  }
  if(user == nil) {
    renderError(w, "You aren't logged in!")
    return
  }

  vars := mux.Vars(r)
  tweetid,exists := vars["tweetid"]

  if(exists) {
    err = deleteTweet(s, *user, tweetid)
    if(err != nil) {
      // TODO don't leak errors
      renderError(w, fmt.Sprintf("%s", err))
      return
    }
  }

  redirect_url,err := s.router.Get("pending").URLPath()
  http.Redirect(w, r, redirect_url.String(), 302)
  return
}



func pendingHandler(s Services, w http.ResponseWriter, r *http.Request) {
  user,err := getUser(s,w,r)
  if(err != nil) {
    // TODO don't leak errors
    renderError(w, fmt.Sprintf("%s", err))
    return
  }
  if(user == nil) {
    renderError(w, "You aren't logged in!")
    return
  }

  tweets,err := getTweets(s,*user)
  if(err != nil) {
    renderError(w, fmt.Sprintf("%s", err))
    return
  }

  args := map[string]interface{}{
      "pending": "yes",
      "username": user.screenName,
      "tweets": tweets,
    }

  // TODO get actual pending tweets
  data := mustache.RenderFileInLayout(
    "/usr/share/tweetautofeeder/templates/pending_page.must",
    "/usr/share/tweetautofeeder/templates/layout.must",
    args)
  w.Write([]byte(data))

  return
}


