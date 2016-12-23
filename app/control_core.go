package app

import (
  "net/http"
  "fmt"
  "appengine"
  "appengine/datastore"
  //"appengine/blobstore"
  //"encoding/csv"

  "strings"
  //"strconv"
  "encoding/json"
  "time"

)


/****************************CORE*********************************/
func control_core(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET LIST/ADD**************************/
  if r.Method == "GET"{
  
  /*
  	 uploadURL, err := blobstore.UploadURL(c, "/"+AdminSlug+"/files/", nil)

  	 if err != nil {
  	 	fmt.Fprintln(w,"error with upload:" + err.Error())
  	 	return
  	 }*/

    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Core",
       //"uploadURL":uploadURL.String(), 
     } 


    //QUERY
    q := datastore.NewQuery("core").Order("-Last_Update").Limit(1)


    //DB GET ALL
    var db []*Core 
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
         "CoreHeader":db[i].Core_Header,
         "CoreFooter":db[i].Core_Footer,
         "CoreVersion":db[i].Core_Version,

         /*Fields: map[string]string{
            "Name":"text",
            "Email":"text",
            "Phone":"text",
            "Message":"textarea", 
          },*/


       },

      //END APPEND
      )

    //END FOR
    } 

    //ELSE NO DATA
    } else {

             //FLUSH CACHE
        cacheFlush("CoreHeader",r)
        cacheFlush("CoreFooter",r)

        //MAP FORM VALS
        dbEntry := Core {
          Core_Header:  DEFAULT_CSS_HEADER,
          Core_Footer:  DEFAULT_CSS_FOOTER,
          Core_Version: "default",
          Last_Update: time.Now(),
        }
        

        //DB PUT
        key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "core", nil), &dbEntry)
          
          //IF ERRORS
          if err != nil {
              fmt.Fprint(w,"error adding default")
              return

          //NO ERRORS
          } else {

            dbData = append(dbData,
              map[string]string {
               //"title": db[i].Filename,
               "key":key.Encode(),
               "CoreHeader":DEFAULT_CSS_HEADER,
               "CoreFooter":DEFAULT_CSS_FOOTER,
               "CoreVersion":"default",
          
             },
            )

          //END NO ERRORS
          }
    
    //END CHECK DATA
    }


  renderControl(w, r, "control/control_core.html", data, dbData)
  


/********************************POST ADD*********************************/
} else {



  //PARAMETERS
  itemID := strings.Split(r.RequestURI,"/")

  //DECODE KEY
  key,err := datastore.DecodeKey(itemID[3])
    
  //KEY ERR
  if err != nil {
    fmt.Fprintln(w, "error decoding")
    return
  }

      //fmt.Fprintln(w,"posted ajax:")
//GET FORM VALS
  formVal := func(val string)string{
    return r.FormValue(val)
  }
  
  //MAP FORM VALS
  dbEntry := Core {
    Core_Header: formVal("core_header"),
    Core_Footer: formVal("core_footer"),
    Core_Version:formVal("core_version"),
    Last_Update: time.Now(),
  }
  

  //DB PUT
  _,errPut := datastore.Put(c, key, &dbEntry)
    
    //IF ERRORS
    if errPut != nil {
        fmt.Fprint(w,"error adding")
        return

    //NO ERRORS
    } else {


        //FLUSH CACHE
        cacheFlush("CoreHeader",r)
        cacheFlush("CoreFooter",r)

  

      //PREP JSON
      m := map[string]string{
        "message":"new core saved",
        "key":key.Encode(),
        "title":dbEntry.Core_Version,
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

