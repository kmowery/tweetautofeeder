package main

import (
  "net/http"
  "github.com/hoisie/mustache"
)

func blogHandler(w http.ResponseWriter, r *http.Request) {
  data := mustache.RenderFile("/usr/share/3amh/templates/blog_main.must", map[string]string{"thing":"places"})
  w.Write([]byte(data))
  return
}

