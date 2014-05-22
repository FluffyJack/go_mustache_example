package main

import (
  "fmt"
  "net/http"
  "os"
  "path"
  "github.com/hoisie/mustache"
)

func homePage(w http.ResponseWriter, r *http.Request) {
  filename := path.Join(path.Join(os.Getenv("PWD"), "views"), "home.mustache")
  data := map[string]string{"name": r.FormValue("name")}
  html := mustache.RenderFile(filename, data)
  fmt.Fprintf(w, html)
}

func main() {
  http.HandleFunc("/", homePage)
  http.ListenAndServe(":8080", nil)
}
