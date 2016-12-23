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


/****************************DESIGNS*********************************/
func control_design(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET LIST/ADD**************************/
  if r.Method == "GET"{


    //QUERY
    queryCore := datastore.NewQuery("core").Order("-Last_Update").Limit(1)


    //DB GET ALL
    var dbCore []*Core
    kCore := ""
    keysCore,err := queryCore.GetAll(c,&dbCore)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting items: " + err.Error())
      return
    }

    //VAR
   // var dbCoreData []map[string]string

    //fmt.Fprintln(w,keys)

    //CHECK IF DATA EXISTS
    if dbCore != nil {

          //KEYS ENCODE
      kCore = keysCore[0].Encode()
/*
    //FOR DB ITEMS THAT EXIST
    for i := range dbCore {
      
  

      dbCoreData = append(dbCoreData,
        map[string]string {
         //"title": db[i].Filename,
         "key":kCore,
         "CoreHeader":dbCore[i].Core_Header,
         "CoreFooter":dbCore[i].Core_Footer,
         "CoreVersion":dbCore[i].Core_Version,



       },

      //END APPEND
      )

    //END FOR
    } */
    
    //END IF NIL
    }

  
  /*
  	 uploadURL, err := blobstore.UploadURL(c, "/"+AdminSlug+"/files/", nil)

  	 if err != nil {
  	 	fmt.Fprintln(w,"error with upload:" + err.Error())
  	 	return
  	 }*/

    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Design",
       //"uploadURL":uploadURL.String(), 
     } 


    //QUERY
    q := datastore.NewQuery("style").Order("-Last_Update").Limit(100)


    //DB GET ALL
    var db []*Style 
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
         "StyleName":db[i].Style_Name,
         "StyleSheet":db[i].Style_Sheet,
         "keyCore":kCore,
         "CoreHeader":dbCore[0].Core_Header,
         "CoreFooter":dbCore[0].Core_Footer,
         "CoreVersion":dbCore[0].Core_Version,
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

    //ELSE NO DATA INSERT DEFAULT
    } else {


          //MAP FORM VALS
          dbEntry := Style {
            Style_Name: "default",
            Style_Sheet: DEFAULT_CSS,
            Last_Update: time.Now(),
          }
          

          //DB PUT
          key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "style", nil), &dbEntry)
            
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
                 "StyleName":"default",
                 "StyleSheet":DEFAULT_CSS,
                  
               },
              )

            }



    //END CHECK DATA
    }


  renderControl(w, r, "control/control_design.html", data, dbData)
  


/********************************POST ADD*********************************/
} else {

      //fmt.Fprintln(w,"posted ajax:")
//GET FORM VALS
  formVal := func(val string)string{
    return r.FormValue(val)
  }
  
  //MAP FORM VALS
  dbEntry := Style {
    Style_Name: formVal("style_name"),
    Style_Sheet: formVal("style_sheet"),
    Last_Update: time.Now(),
  }
  

  //DB PUT
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "style", nil), &dbEntry)
    
    //IF ERRORS
    if err != nil {
        fmt.Fprint(w,"error adding")
        return

    //NO ERRORS
    } else {

    
      //PREP JSON
      m := map[string]string{
        "message":"new settings saved",
        "key":key.Encode(),
        "title":dbEntry.Style_Name,
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


//FUNC EDIT
func control_design_edit(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //PARAMETERS
  itemID := strings.Split(r.RequestURI,"/")

  //CONTEXT
  c := appengine.NewContext(r)

  //DECODE KEY
  key,err := datastore.DecodeKey(itemID[4])
    
  //KEY ERR
  if err != nil {
    fmt.Fprintln(w, "error decoding")
    return
  }

  /*******************************GET EDIT**************************/
  if r.Method == "GET"{

    //DB QUERY
    q := datastore.NewQuery("style").Filter("__key__ =", key).Limit(1)

    //VAR
    var bb []*Style

    //DB GET ALL
    _,er := q.GetAll(c,&bb)
    
    //DB ERR
    if er != nil {    
      fmt.Fprintln(w, "error getting style")
      return
    }


    //PREP JSON
    m := Style {
      bb[0].Style_Name,
      bb[0].Style_Sheet,
      bb[0].Last_Update,
    }

    //MARSHAL JSON
    j,errJSON := json.Marshal(m)
    if errJSON != nil {
      fmt.Fprintln(w,"error with JSON")
    }

    //SET CONTENT-TYPE
    w.Header().Set("Content-Type", "application/json")

    //DISPLAY JSON
    fmt.Fprint(w,string(j))


  /********************************POST EDIT*********************************/
  } else {
  
  
    //GET FORM VALS
    formVal := func(val string)string{
      return r.FormValue(val)
    }

    
    //MAP FORM VALS
    newStyle := &Style {
      Style_Name: formVal("style_name"),
      Style_Sheet: formVal("style_sheet"),
      Last_Update: time.Now(),
    }

    sName := string("StyleName_" + formVal("style_name"))
    sSheet := string("StyleSheet_" + formVal("style_name"))
 
    //DB PUT
    _,errPut := datastore.Put(c, key,newStyle)
    
    //DB ERR
    if errPut != nil {
        fmt.Fprintln(w,"error Put")
        return
    } else {

        //string("StyleName" + formVal("style_name"))

        //FLUSH CACHE
        cacheFlush(sName,r)
        cacheFlush(sSheet,r)
    }

    //SETUP JSON
    m := map[string]string{
      "message": "update success",
      "title": newStyle.Style_Name,
    }

    //MARSHAL JSON
    j,errJSON := json.Marshal(m)
    if errJSON != nil {
      fmt.Fprintln(w,"error with JSON")
    }

    //SET CONTENT-TYPE
    w.Header().Set("Content-Type", "application/json")


    //DISPLAY JSON
    fmt.Fprint(w,string(j))

  //END POST
  }
  
  
//END FUNC
}



//FUNC DELETE
func control_design_delete(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  /********************************POST DELETE*********************************/
  if r.Method == "POST"{
  
    //CONTEXT
    c := appengine.NewContext(r)

    formVal := func(val string)string{
      return r.FormValue(val)
    }

    //GET KEY
    itemID := strings.Split(r.RequestURI,"/") 

    
    //KEY DECODE
    key,err := datastore.DecodeKey(itemID[4])
    
    //KEY ERR
    if err != nil {
      fmt.Fprintln(w, "error decoding")
      return
    }
   
    //DB DELETE
    errDelete := datastore.Delete(c, key)
    
    //DB ERR
    if errDelete != nil {
      fmt.Fprintln(w,"error deleting")
      return
    
    //NO ERR
    } else {

      //sName := formVal('styleName')
      //FLUSH CACHE
      cacheFlush(string("StyleName_" + formVal("styleName")),r)
      cacheFlush(string("StyleSheet_" + formVal("styleName")),r) 
 
      //cacheFlush(sSheet,r)
      fmt.Fprintln(w,"delete successful")
      return
    }


  //END POST
  } else {
    fmt.Fprintln(w, "say whaaa")
  }
  
  
//END FUNC
}


