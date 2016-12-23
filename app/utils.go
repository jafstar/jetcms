package app

import (
  "net/http"
  "net/url"
  "strings"
  //"fmt"
  //"appengine"
  //"appengine/datastore"
  "time"
    //"crypto/sha256"
  //"bytes"
  //"encoding/base64"
)

func cleanURL(r *http.Request) string {
  realURL,_ := url.Parse(r.RequestURI)
  stringURL := realURL.String()

  if stringURL == "/" {
  	stringURL = "home"
  } else {
  	stringURL = strings.TrimLeft(stringURL,"/")
  }
  
  return stringURL
}

func hash(item1 string, item2 string) string {

/*
  h := sha256.New()
  h.Write([]byte(item1))
  h.Write([]byte(item2))
  
  sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
  
  return sha
*/

  return item1 + item2
}

func token(r *http.Request) string {

  ip := r.RemoteAddr
  date := time.Now().String()
  token := hash(ip, date)
  
  
  return token
}


