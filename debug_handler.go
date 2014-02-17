package main

import (
  "net/http"
  "fmt"
  "github.com/hoisie/mustache"
)

func debugHandler(w http.ResponseWriter, r *http.Request) {

  fmt.Printf("Headers: %s\n", r.Header)

  r.ParseForm()
  fmt.Printf("Post:    %s\n", r.Form)
  fmt.Printf("Post:    %s\n", r.Form)

  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"thing":"places"})
  w.Write([]byte(data))
  return
}

