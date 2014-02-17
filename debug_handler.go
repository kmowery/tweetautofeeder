package main

import (
  "net/http"
  //"net/http/httputil"
  "io/ioutil"
  "fmt"
  "github.com/hoisie/mustache"
)

func debugHandler(w http.ResponseWriter, r *http.Request) {

  r.URL.Scheme = "http"
  r.URL.Host = "localhost"

  //fmt.Printf("Headers: %s\n", r.Header)

  if(r.Body != nil) {
    body, _ := ioutil.ReadAll(r.Body)
    fmt.Printf("Here's the body: %s\n\n", body)

  }

  //dump, err := httputil.DumpRequestOut(r, false)
  //if(err != nil) {
  //  fmt.Printf("debug, some error: ")
  //  fmt.Println(err)
  //}
  //fmt.Printf("DEBUG: %s", string(dump))
  //fmt.Printf("DEBUG END\n\n")

  fmt.Printf("Headers: %s\n", r.Header)

  //fmt.Printf("Body:    %s\n", body,_ := ioutil.ReadAll(resp.Body)
  r.ParseForm()
  fmt.Printf("Post:    %s\n", r.Form)
  fmt.Printf("PostForm:%s\n", r.PostForm)

  data := mustache.RenderFile("/usr/share/tweetautofeeder/templates/blog_main.must", map[string]string{"thing":"places"})
  w.Write([]byte(data))
  return
}

