package main

import (
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "strings"
  "time"
)

func myHandler(c http.ResponseWriter, req *http.Request) {

  log.Print(req.URL.RequestURI())

  if req.URL.RequestURI() == "/" {

    http.ServeFile(c, req, "./index.html")

  } else if !(strings.HasPrefix(req.URL.RequestURI(), "/ebauche/Telechargements/ical")) {

    http.Error(c, "Tu veux du pain ?", 403)

  } else {

    resp, _ := http.Get("http://agenda.ecp.fr" + req.URL.RequestURI())

    defer resp.Body.Close()
    contents, _ := ioutil.ReadAll(resp.Body)

    // This is needed for Google Calendar
    c.Header().Set("Content-Type", "text/calendar; charset=utf-8")

    io.WriteString(c, string(contents))
  }

}

func main() {

  s := &http.Server{
    Addr:           ":" + os.Getenv("PORT"),
    Handler:        http.HandlerFunc(myHandler),
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }
  log.Print("listening...")
  log.Fatal(s.ListenAndServe())
}
