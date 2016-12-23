package app


import (
  "net/http"
  "fmt"
  "appengine"
  "appengine/datastore"
  //"appengine/blobstore"
  "time"
  "strings"
  //"strconv"
  "encoding/json"
  //"net/url"
  "regexp"
)

/****************************GET DATA ITEM DETAILS********************************
func control_data_item_details(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //PARAMETERS
  itemID := strings.Split(r.RequestURI,"/")

  //fmt.Fprintln(w,itemID[5])

  //CONTEXT
  c := appengine.NewContext(r)

  //DECODE KEY
  key,err := datastore.DecodeKey(itemID[5])

  //KEY ERR
  if err != nil {
    fmt.Fprintln(w, "error decoding")
    return
  }


  //GET
  if r.Method == "GET"{


    //QUERY
    q := datastore.NewQuery("DataItem").Filter("__key__ =", key).Limit(1).Limit(100)


    //DB GET ALL
    var db []*DataItem
    _,err := q.GetAll(c,&db)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting items")
      return
    }

    //VAR
    dbData := []DataItem{}

    //FOR LOOP 1 DB ITEMS
    for i := range db {


      //APPEND
      dbData = append(dbData,
        DataItem {
            Content: db[i].Content,
            Type: db[i].Type,
            Tags: db[i].Tags,
            Title: db[i].Title,
            URL: db[i].URL,
         //Title: db[i].Title,
         //"key": k,
         //Fields: TheFields,
       },
      )

    //END FOR 1
    }


    //fmt.Fprintln(w,data)
    //fmt.Fprintln(w,dbData)
    //fmt.Fprintln(w,r.Header.Get("X-Requested-With")) //"X-Requested-With"

    if r.Header.Get("X-Requested-With") != "" {

        //MARSHAL JSON
        j,errJSON := json.Marshal(dbData)
        if errJSON != nil {
          fmt.Fprintln(w,"error with JSON")
        }

        //SET CONTENT-TYPE
        w.Header().Set("Content-Type", "application/json")

        //DISPLAY JSON
        fmt.Fprint(w,string(j))
      
    } 

  //END GET    
  }

//END FUNC
}*/



/****************************GET LIST PAGES/POST ADD PAGES*********************************/
func control_data_item(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)


  /*******************************GET LIST/ADD**************************/
  if r.Method == "GET"{
  
    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Content",
     } 


    //QUERY
    q := datastore.NewQuery("DataItem").Limit(100)


    //DB GET ALL
    var db []*DataItem
    keys,err := q.GetAll(c,&db)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting items...")
      return
    }




    //VAR
    var dbData []map[string]string

    //FOR DB ITEMS
    for i := range db {
      
      //KEYS ENCODE
      k := keys[i].Encode()

      //TIME PARSE
      //itemTime,_ := time.Parse("2006-01-02 00:00:00 +0000 UTC",db[i].DateOrder.String())
      //itemTime.Format("Jan 2, 2006")

      dbData = append(dbData,
        map[string]string {
         "title": db[i].Title,
         "url": db[i].URL,
         "type": db[i].Type,
         "date": db[i].DateOrder.Format(ItemDateDisplay),
         "typeurl":db[i].TypeURL,
         "key":k,
         /*Fields: map[string]string{
            "Name":"text",
            "Email":"text",
            "Phone":"text",
            "Message":"textarea", 
          },*/


       },
      )
    }


  renderControl(w, r, "/control/data_item.html", data, dbData)
  


/********************************POST ADD*********************************/
} else {

/*blobs, _, err := blobstore.ParseUpload(r)

if err != nil {
  fmt.Fprintln(w,"error with upload:"+err.Error())
  //serveError(c, w, err)
    return
}

file := blobs["file"]

if len(file) == 0 {
    fmt.Fprintln(w,"file is nil")  
    //http.Redirect(w, r, "/admin/files/", http.StatusFound)
    return
}*/


  //GET FORM VALS
  formVal := func(val string)string{
    return r.FormValue(val)
  }

  //newFields := map[string]string {}
  //fmt.Fprintln(w,r)

  fieldNames := strings.Split(formVal("fieldNames"),",")

  //fmt.Fprintln(w,fieldNames)
  //fmt.Fprintln(w,len(fieldNames))



  //TheFields := []FieldItem{}
    //TheFields = make(FieldItem)

  var TheFields map[string]string

  TheFields = make(map[string]string)

  FieldCount := len(fieldNames)-1

    
    //FOR LOOP
    for i := 0; i <= FieldCount; i++ {
     

      TheFields[fieldNames[i]] = formVal(fieldNames[i])
      //TheFields["thing"] = "whatever"
     
      /*
      TheFields = append(TheFields, map[string]string {
      //string(fieldNames[i]) : formVal(fieldNames[i]),
       
        fieldNames[i]: formVal(fieldNames[i]),
      //Value:  map[string]string{
          //        fieldNames[i]:formVal(fieldNames[i]),
            //    },
      })*/
      

      //fmt.Fprintln(w,fieldNames[i])
      //fmt.Fprintln(w,formVal(fieldNames[i]))
    
    //END FOR LOOP
    }

    //fmt.Fprintln(w,TheFields)

      //let's make pretty urls from title
    reg, err := regexp.Compile("[^A-Za-z0-9]+")
    
    if err != nil {
        fmt.Fprintln(w,"error with RegX")
    }
    
    prettyurl := reg.ReplaceAllString(formVal("title"), "-")
    prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))



    //fmt.Fprintln(w,prettyurl)
    //fmt.Fprintln(w,TheFields["title"])
  
    
    /*
      {
        Name: formVal("fieldName1"),
        Label: formVal("fieldLabel1"),
        UI: formVal("fieldUI1"),
        Errors: formVal("fieldErrors1"),
      },
      {
        Name: formVal("fieldName2"),
        Label: formVal("fieldLabel2"),
        UI: formVal("fieldUI2"),
        Errors: formVal("fieldErrors2"),
      },
    }*/

  //MARSHAL JSON
  j,errJSON := json.Marshal(TheFields)

  if errJSON != nil {
    fmt.Fprintln(w,"error with JSON")
  }

    //loc, _ := time.LoadLocation("America/New_York")
  
  //FORMAT DATE
  const shortForm = "2006-01-02 15:04:05 +0000 UTC"
  timeNow := time.Now()

  //dateOrder,_ := time.Parse(shortForm, formVal("fieldDateOrder"))
  //dateOrder,_ := time.ParseInLocation(shortForm, timeNow.String(),loc)

  dateOrder,_ := time.Parse(shortForm, timeNow.String())




//PUBLISH CHECK
        publishAddBool := true

  if formVal("info_publish") == "on" {
    publishAddBool = true
  } else {
    publishAddBool = false
  }


  //MAP FORM VALS
  newItem := DataItem {
    Title: formVal("title"),
    DateOrder:  dateOrder,
    Content: j,
    Type: strings.ToUpper(formVal("fieldType")),
    //Type string
    //Tags []string
    //Title string
    Published: publishAddBool,
    URL:  prettyurl,
    TypeURL: formVal("fieldTypeUrl"),
    Description: formVal("info_description"),
    Keywords: formVal("info_keywords"),
  }

  //DEBUG
  //fmt.Fprintln(w,j)
  //var stuff = map[string]string{}
  //json.Unmarshal(newItem.Content, &stuff)
  //fmt.Fprintln(w,newItem)
  //fmt.Fprintln(w,stuff["location"])

  //DB PUT
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "DataItem", nil), &newItem)
    
    //IF ERRORS
    if err != nil {
        fmt.Fprintln(w,"error adding")
        fmt.Fprintln(w,err)
        return

    //NO ERRORS
    } else {

      //FLUSH CACHE
      cacheFlush("dataURLs",r)

     // fmt.Fprintln(w,"added successfully")
     // fmt.Fprintln(w,"key: " + key.Encode())

      
      //PREP JSON
      m := map[string]string{
        "message":"new item added",
        "key":key.Encode(),
        "title":newItem.Title,
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



//FUNC CONTROL DATA ITEMS EDIT
func control_data_item_edit(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //PARAMETERS
  itemID := strings.Split(r.RequestURI,"/")

  //CONTEXT
  c := appengine.NewContext(r)


//fmt.Fprintln(w,itemID[5])

  //DECODE KEY
  key,err := datastore.DecodeKey(itemID[5])
    
  //KEY ERR
  if err != nil {
    fmt.Fprintln(w, "error decoding")
    //fmt.Fprintln(w,itemID[4])
    return
  }



  /*******************************GET EDIT**************************/
  if r.Method == "GET"{


    //DB QUERY
    q := datastore.NewQuery("DataItem").Filter("__key__ =", key).Limit(1)

    //VARS
    var bb []*DataItem
    data := map[string]map[string]string{}


    //DB GET ALL
    _,er := q.GetAll(c,&bb)
    
    //DB ERR
    if er != nil {    
      fmt.Fprintln(w, "error getting data items")
      return
    }

    //UNMARSHALL
    var jsonContent = map[string]string{}
    json.Unmarshal(bb[0].Content, &jsonContent)


    /*
    //CREATE UPLOAD URL
    uploadURL, errURL := blobstore.UploadURL(c, "/"+AdminSlug+"/files/", nil)

     if errURL != nil {
      fmt.Fprintln(w,"error with making upload URL on GET:" + errURL.Error())
      return
     }
     */
     publishStringBool := ""

     if bb[0].Published {
      publishStringBool = "True"
     } else {
      publishStringBool = "False"
     }

    //DATA
    data = map[string]map[string]string{
      "info":map[string]string{
          "title":bb[0].Title, 
          "url":bb[0].URL,
          "description":bb[0].Description,
          "keywords":bb[0].Keywords,
          "published":publishStringBool,
          //"uploadURL": uploadURL.String(),
          //"tags":bb[0].Tags,
          "typeurl":bb[0].TypeURL,
          "type":bb[0].Type,
        },
      "content": jsonContent,

    //END DATA
    }


    //MARSHAL JSON
    j,errJSON := json.Marshal(data)
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
  

  //PREPARE FIELDS
  fieldNames := strings.Split(formVal("fieldNames"),",")

  var TheFields map[string]string
      TheFields = make(map[string]string)

  FieldCount := len(fieldNames)-1

    
  //FOR LOOP
  for i := 0; i <= FieldCount; i++ {
     
      //CAPTURE FIELDS
      TheFields[fieldNames[i]] = formVal(fieldNames[i])
    
  //END FOR LOOP
  }


  //PRETTY URLS
  reg, err := regexp.Compile("[^A-Za-z0-9]+")
    
  if err != nil {
    fmt.Fprintln(w,"error with RegX")
  }
    
  prettyurl := reg.ReplaceAllString(formVal("title"), "-")
  prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))


  //MARSHAL JSON
  binaryContent,errJSON := json.Marshal(TheFields)

  if errJSON != nil {
    fmt.Fprintln(w,"error with JSON")
  }

  //DATE ORDER
  const shortForm = "2006-01-02"
  dateOrder,_ := time.Parse(shortForm, formVal("fieldDateOrder"))

      publishBool := true


//PUBLISH CHECK
  if formVal("info_publish") == "on" {
    publishBool = true
  } else {
    publishBool = false
  }



  //MAP FORM VALS
  newItem := DataItem {
    Title: formVal("title"),
    DateOrder:  dateOrder,
        Published: publishBool,

    Content: binaryContent,
    Type: strings.ToUpper(formVal("fieldType")),
    URL:  prettyurl,
    TypeURL: formVal("fieldTypeUrl"),
        Description: formVal("info_description"),
    Keywords: formVal("info_keywords"),

  }

 
    //DB PUT
    _,errPut := datastore.Put(c, key, &newItem)
    
    //DB ERR
    if errPut != nil {
        fmt.Fprintln(w,"error Putting Data Item")
        return


   //NO ERRORS
    } else {

        //SETUP JSON
        m := map[string]string{
          "message":"update success",
          "title": newItem.Title,
        }

        //MARSHAL JSON
        j,errJSON2 := json.Marshal(m)
        if errJSON2 != nil {
          fmt.Fprintln(w,"error with JSON")
        }

        //SET CONTENT-TYPE
        w.Header().Set("Content-Type", "application/json")


        //DISPLAY JSON
        fmt.Fprint(w,string(j))

    //END NO ERRORS
    }
    

  //END POST
  }
  
  
//END FUNC
}



//FUNC CONTROL PAGES
func control_data_item_delete(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  /********************************POST DELETE*********************************/
  if r.Method == "POST"{
  
    //CONTEXT
    c := appengine.NewContext(r)

    //GET KEY
    itemID := strings.Split(r.RequestURI,"/") 

    
    //KEY DECODE
    key,err := datastore.DecodeKey(itemID[5])
    
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
      fmt.Fprintln(w,"delete successful")
      return
    }


  //END POST
  } else {
    fmt.Fprintln(w, "say whaaa")
  }
  
  
//END FUNC
}

