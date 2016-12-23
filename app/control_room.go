package app

import (
  "net/http"
   "appengine"
  "appengine/capability"
  
  //"appengine/urlfetch"
  //"log"
  //"io/ioutil"

  //"strings"
  //"fmt"
)



//ROOM
func room(w http.ResponseWriter, r *http.Request) {

  //CHECK ADMIN
  checkAdmin(w,r)


  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET**************************/
  if r.Method == "GET"{
  
  

  //VARS
  checkDataRead  := "checkOnline"
  checkDataWrite := "checkOnline"
  checkCache := "checkOnline"
  checkMail := "checkOnline"
  checkBlob := "checkOnline"


  //CHECK DB READ
  if !capability.Enabled(c, "datastore_v3", "*") {
        http.Error(w, "The Database read capability is currently unavailable.", 503)
        checkDataRead = "checkOffline"
        return
  }

  //CHECK DB WRITE
  if !capability.Enabled(c, "datastore_v3", "write") {
        http.Error(w, "The Database write capability is currently unavailable.", 503)
        checkDataWrite = "checkOffline"
        return
  } 
  
  //CHECK MAIL
  if !capability.Enabled(c, "mail", "*") {
        http.Error(w, "The Mail capability is currently unavailable.", 503)
        checkMail = "checkOffline"
        return
  } 

    //CHECK MEMCACHE
  if !capability.Enabled(c, "memcache", "*") {
        http.Error(w, "The Memcache capability is currently unavailable.", 503)
        checkCache = "checkOffline"
        return
  } 

    //CHECK BLOBSTORE
  if !capability.Enabled(c, "blobstore", "*") {
        http.Error(w, "The blobstore capability is currently unavailable.", 503)
        checkBlob = "checkOffline"
        return
  } 


  //http://quotes.rest/qod.json?category=inspire

/**********************QUOTE OF THE DAY**********************/
/*  QOD_TEXT := ""
  //QOD_IMG := ""

    //SETUP URL FETCH
    client := urlfetch.Client(c)
      
      //GET URL
      resp, err := client.Get("http://quotes.rest/qod.json?category=inspire")

      //ERR
      if err != nil { 
        //DISPLAY WITH REDIRECT
        //log.Fatal(err.Error()) //http.Error(w, err.Error(), http.StatusInternalServerError)
          return
      }


    //READCLOSER -> BYTES
    img, err := ioutil.ReadAll(resp.Body)

    //CLOSE
    resp.Body.Close()
    
    //ERR
    if err != nil {
      log.Fatal(err)
    }

      QOD_TEXT = string(img)*/

 /**********************USER INFO**********************/     
currentUser := getCurrentUser(w,r)
//currentUser := "You gotta login"

/**********************DATA MAP**********************/

  data := map[string]string{
    "title": "Welcome",
             "checkDataRead":checkDataRead,
          "checkDataWrite":checkDataWrite,
          "checkMail":checkMail,
          "checkCache":checkCache,
          "checkBlob":checkBlob,
          "dataCenter":appengine.Datacenter(),
          "serverSoftware":appengine.ServerSoftware(),
          "versionID":appengine.VersionID(c),
          //"qodText": QOD_TEXT,
          "currentUser": currentUser,

   /* "checkDataRead":checkDataRead,
    "checkDataWrite":checkDataWrite,
    "checkMail":checkMail,
    "checkCache":checkCache,
    "checkBlob":checkBlob,
    "dataCenter":appengine.Datacenter(),
    "serverSoftware":appengine.ServerSoftware(),*/
    //"requestID":appengine.RequestID(c),

  }

  //fmt.Fprintln(w,data)
  //fmt.Fprintln(w,appengine.Datacenter())
  
  //RENDER
  renderControl(w, r, "control/control_room.html", data,nil)

  //END IF GET
  }

//END FUNC  
}



//MENUS
func menus(w http.ResponseWriter, r *http.Request) {

  checkAdmin(w,r)
    
  data := map[string]string{
    "title": "Menus",
  }

  renderControl(w, r, "control/control_menus.html", data,nil)
  
}


//SLUGS
func panels(w http.ResponseWriter, r *http.Request) {

  checkAdmin(w,r)
    
  data := map[string]string{
    "title": "Panels",
  } 

  renderControl(w, r, "control/control_slugs.html", data,nil)
  
}
