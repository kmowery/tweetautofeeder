package main

import (
  "net/http"
  "fmt"
  "github.com/hoisie/mustache"
  "github.com/mrjones/oauth"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {

  c := oauth.NewConsumer(
    API_CONSUMER_KEY,
    API_CONSUMER_SECRET,
    oauth.ServiceProvider{
      RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
      AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
      AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
    })

  c.Debug(true)

  requestToken, url, _ := c.GetRequestTokenAndUrl("http://callback.com")

  fmt.Printf("token: %s\n", requestToken)
  fmt.Printf("url: %s\n", url)

  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"thing":"places"})
  w.Write([]byte(data))
  return
}

