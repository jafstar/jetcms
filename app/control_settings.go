package app


import (
  "net/http"
  "fmt"
  "appengine"
  "appengine/datastore"
   // "appengine/capability"

  //"appengine/blobstore"
  //"encoding/csv"

  //"strings"
  //"strconv"
  "encoding/json"
  "time"

)


//var queryType = "data"


/****************************GET LIST PAGES/POST ADD PAGES*********************************/
func control_settings(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET**************************/
  if r.Method == "GET"{
  
/*
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

*/
  /*
  	 uploadURL, err := blobstore.UploadURL(c, "/"+AdminSlug+"/files/", nil)

  	 if err != nil {
  	 	fmt.Fprintln(w,"error with upload:" + err.Error())
  	 	return
  	 }*/

    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Settings",
       //"uploadURL":uploadURL.String(), 
     } 


    //QUERY
    q := datastore.NewQuery("settings").Order("-Last_Update").Limit(1)


    //DB GET ALL
    var db []*Settings 
    keys,err := q.GetAll(c,&db)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting items: " + err.Error())
      return
    }

    //VAR
    var dbData []map[string]string

    //fmt.Fprintln(w,keys)

    //CHECK IF DATA EXISTS
    if db != nil {

    //FOR DB ITEMS THAT EXIST
    for i := range db {
      
      //KEYS ENCODE
      k := keys[i].Encode()

      dbData = append(dbData,
        map[string]string {
         //"title": db[i].Filename,
         "key":k,
        
         "SiteTitle":db[i].Site_Title,
         "SiteClosed":db[i].Site_Closed,
         "SiteAnalytics":db[i].Site_Analytics,
        
        // "HomeTitle":db[i].Home_Title,
         //"HomeDesc":db[i].Home_Desc,
         //"HomeKeyW":db[i].Home_KeyW,
         //"SoonTitle":db[i].Soon_Title,
         //"SoonDesc":db[i].Soon_Desc,
        // "SoonKeyW":db[i].Soon_KeyW,
        "SoonURL":db[i].Soon_URL,
        
         "AdminEmail":db[i].Admin_Email,
        // "FromEmail":db[i].From_Email,
        // "ToEmail":db[i].To_Email,
        
         "CacheControl":db[i].Cache_Control,
         "CacheDays":db[i].Cache_Days,
         "CacheMonths":db[i].Cache_Months,

         /*Fields: map[string]string{
            "Name":"text",
            "Email":"text",
            "Phone":"text",
            "Message":"textarea", 
          },*/
         // "checkDataRead":checkDataRead,
          //"checkDataWrite":checkDataWrite,
          //"checkMail":checkMail,
         // "checkCache":checkCache,
         // "checkBlob":checkBlob,
          //"dataCenter":appengine.Datacenter(),
          //"serverSoftware":appengine.ServerSoftware(),


       },

      //END APPEND
      )

    //END FOR
    } 

    //ELSE NO DATA
    } else {

      dbData = append(dbData,
        map[string]string {
         //"title": db[i].Filename,
         //"key":k,
         "SiteTitle":"",
         "SiteClosed":"",
         "SiteAnalytics":"",
        // "HomeTitle":"",
         //"HomeDesc":"",
         //"HomeKeyW":"",
         //"SoonTitle":"",
         //"SoonDesc":"",
         //"SoonKeyW":"",
         "SoonURL":"",
         "AdminEmail":"",
         //"FromEmail":"",
         //"ToEmail":"",
         "CacheControl":"",
         "CacheDays":"",
         "CacheMonths":"",
         

       },
      )

    //END CHECK DATA
    }


  renderControl(w, r, "/control/settings.html", data, dbData)
  


/********************************POST*********************************/
} else {

      //fmt.Fprintln(w,"posted ajax:")
//GET FORM VALS
  formVal := func(val string)string{
    return r.FormValue(val)
  }
  
  //MAP FORM VALS
  dbEntry := Settings {
    Site_Title: formVal("site_title"),
    Site_Closed: formVal("site_closed"),
    Site_Analytics:formVal("site_analytics"),
   // Home_Title: formVal("home_title"),
    //Home_Desc: formVal("home_desc"),
    //Home_KeyW: formVal("home_keyw"),
    //Soon_Title: formVal("soon_title"),
   // Soon_Desc: formVal("soon_desc"),
    //Soon_KeyW: formVal("soon_keyw"),
    Soon_URL:formVal("soon_url"),
    Cache_Control:formVal("cache_control"),
    Cache_Days:formVal("cache_days"),
    Cache_Months:formVal("cache_months"),

    Admin_Email: formVal("admin_email"),
    //From_Email: formVal("from_email"),
    //To_Email: formVal("to_email"),
    Last_Update: time.Now(),
  }
  

  //DB PUT
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "settings", nil), &dbEntry)
    
    //IF ERRORS
    if err != nil {
        fmt.Fprint(w,"error adding")
        return

    //NO ERRORS
    } else {

  

      //FLUSH LIST
      list := []string {
         "SiteTitle",
         "SiteClosed",
         "SiteAnalytics",

        // "HomeTitle",
        // "HomeDesc",
        // "HomeKeyW",

         //"SoonTitle",
         //"SoonDesc",
         //"SoonKeyW",
         "SoonURL",
         "CacheControl",
         "CacheDays",
         "CacheMonths",


         "AdminEmail",
        // "FromEmail",
         //"ToEmail",
      }

      //FLUSH CACHE
      cacheFlushMulti(list,r)


      //PREP JSON
      m := map[string]string{
        "message":"new settings saved",
        "key":key.Encode(),
        "title":dbEntry.Site_Title,
        "adminSlug":AdminSlug,
      } 

      //MARSHAL JSON
      j,errJSON := json.Marshal(m)
      if errJSON != nil {
        fmt.Fprintln(w,"error with JSON")
      }

      //DISPLAY JSON
      w.Header().Set("Content-Type", "application/json")
      fmt.Fprint(w,string(j))
      return

    //END ERRORS
    }

//END POST
}

//END FUNC
}
