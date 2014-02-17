package main

import (
  "fmt"
  "errors"
  "log"
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


func TwitterPostTweet(s Services, user *User, tweetid string) error {
  tweet,err := getTweet(s, user, tweetid)
  if(err != nil || tweet == nil) {
    // TODO log maybe?
    return err
  }

  resp, err := s.customer.Post("https://api.twitter.com/1.1/statuses/update.json",
    map[string]string{"status": tweet.tweet},
    &user.atoken)

  if(err != nil) {
    return err
  }

  if(resp.StatusCode == 200) {
    // move this tweet to the "complete" stage

  } else {
    // show error
    return errors.New("Error communicating with Twitter!")
  }

  return nil
}


func postNewHandler(s Services, w http.ResponseWriter, r *http.Request) {
  log.Println("in postNewHandler")
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

  err = addTweet(s,*user, tweet[0])
  if(err != nil) {
    // TODO don't leak errors
    renderError(w, fmt.Sprintf("%s", err))
    return
  }

  redirect_url,err := s.router.Get("pending").URLPath()
  http.Redirect(w, r, redirect_url.String(), 302)
  return
}

func postNowHandler(s Services, w http.ResponseWriter, r *http.Request) {
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
    err = TwitterPostTweet(s, user, tweetid)
    if(err != nil) {
      // TODO don't leak errors
      renderError(w, fmt.Sprintf("%s", err))
      return
    }

    // Mark tweet as posted
    err = completeTweet(s, user, tweetid)
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

  tweets,err := getTweets(s,*user,false)
  if(err != nil) {
    renderError(w, fmt.Sprintf("%s", err))
    return
  }

  args := map[string]interface{}{
      "pending": "yes",
      "username": user.screenName,
      "tweets": tweets,
    }

  data := mustache.RenderFileInLayout(
    "/usr/share/tweetautofeeder/templates/pending_page.must",
    "/usr/share/tweetautofeeder/templates/layout.must",
    args)
  w.Write([]byte(data))

  return
}

func postedHandler(s Services, w http.ResponseWriter, r *http.Request) {
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

  tweets,err := getTweets(s,*user,true)
  if(err != nil) {
    renderError(w, fmt.Sprintf("%s", err))
    return
  }

  args := map[string]interface{}{
      "posted": "yes",
      "username": user.screenName,
      "tweets": tweets,
    }

  data := mustache.RenderFileInLayout(
    "/usr/share/tweetautofeeder/templates/posted_page.must",
    "/usr/share/tweetautofeeder/templates/layout.must",
    args)
  w.Write([]byte(data))

  return
}

