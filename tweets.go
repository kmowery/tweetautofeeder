package main

import (
  "fmt"
  //"log"
  "net/http"
  "github.com/hoisie/mustache"
)

type Tweet struct {
  tweet string
}
func (tw Tweet) Get() string {
  return tw.tweet
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

  args := map[string]interface{}{
      "pending": "yes",
      "username": user.screenName,
      "tweets": []Tweet{ Tweet{"one"}, Tweet{"two"}},
    }

  // TODO get actual pending tweets


  data := mustache.RenderFileInLayout(
    "/usr/share/tweetautofeeder/templates/pending_page.must",
    "/usr/share/tweetautofeeder/templates/layout.must",
    args)
  w.Write([]byte(data))

  return
}


