package app


import (
  "net/http"
  "fmt"
  "appengine"
  "appengine/datastore"
  "strings"
  "strconv"
  "encoding/json"
  "regexp"

)

//var queryType = "DataType"

/****************************GET DATA TYPE DETAILS*********************************/
func control_data_type_details(w http.ResponseWriter, r *http.Request) {
  
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

  /*******************************GET**************************/
  if r.Method == "GET"{


    //QUERY
    q := datastore.NewQuery("DataType").Filter("__key__ =", key).Limit(1).Limit(100)


    //DB GET ALL
    var db []*DataType
    _,err := q.GetAll(c,&db)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting items")
      return
    }

    //VAR
    dbData := []DataType{}

    //FOR LOOP 1 DB ITEMS
    for i := range db {
      
      //KEYS ENCODE
      //k := keys[i].Encode()


      TheFields := []FieldType {}

      //DEBUG
      //fmt.Fprintln(w,k)
      //fmt.Fprintln(w,db[i].Fields[0].Name)
      //fmt.Fprintln(w,len(db[i].Fields)-1)
      
      //FOR LOOP 2 FIELDS
      for j := 0; j <= (len(db[i].Fields)-1); j++ {

        TheFields = append(TheFields, FieldType {
          Name:  db[i].Fields[j].Name ,
          UI:    db[i].Fields[j].UI ,
        })

      //END FOR 2
      }

      //APPEND
      dbData = append(dbData,
        DataType {
         Title: db[i].Title,
         URL: db[i].URL,
         //"key": k,
         Fields: TheFields,
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
}

  



/****************************GET LIST PAGES/POST ADD PAGES*********************************/
func control_data_type(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET LIST/ADD**************************/
  if r.Method == "GET"{
  
    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Data",
     } 


    //QUERY
    q := datastore.NewQuery("DataType").Limit(100)


    //DB GET ALL
    var db []*DataType
    keys,err := q.GetAll(c,&db)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting items")
      return
    }

    //VAR
    var dbData []map[string]string

    //FOR DB ITEMS
    for i := range db {
      
      //KEYS ENCODE
      k := keys[i].Encode()

      dbData = append(dbData,
        map[string]string {
         "title": db[i].Title,
         "key": k,
         //"fieldNames": db[i].Fields[1].Name,
         //"fieldUI":db[i].UI,
         /*Fields: map[string]string{
            "Name":"text",
            "Email":"text",
            "Phone":"text",
            "Message":"textarea", 
          },*/


       },
      )
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
  
} else {
  renderControl(w, r, "/control/data.html", data, dbData)
}
  


/********************************POST ADD*********************************/
} else {


  //GET FORM VALS
  formVal := func(val string)string{
    return r.FormValue(val)
  }

  //newFields := map[string]string {}
  
  //fmt.Fprintln(w,r)

  TheFields := []FieldType {}
  FieldCount,_ := strconv.Atoi(formVal("field_count"))
  FieldCount = FieldCount - 1
 
    for i := 0; i <= FieldCount; i++ {

      iCount := strconv.Itoa(i)

      TheFields = append(TheFields, FieldType {
        Name:   formVal("fieldName" + iCount),
        Order:  formVal("fieldOrder" + iCount),
        UI:     formVal("fieldUI" + iCount),
        //Errors: formVal("fieldErrors" + iCount),
      })

      //fmt.Fprintln(w,formVal("fieldName" + iCount))
      //fmt.Fprintln(w,formVal("fieldName0"))
      //fmt.Fprintln(w, FieldCount)
      //fmt.Fprintln(w,iCount)
    }



  
    
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


    //PRETTIFY URL
    reg, err := regexp.Compile("[^A-Za-z0-9]+")
    
    if err != nil {
        fmt.Fprintln(w,"error with RegX")
    }
    
    prettyurl := reg.ReplaceAllString(formVal("slug"), "-")
    prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))


    //fmt.Fprintln(w,prettyurl)



  //MAP FORM VALS
  newType := DataType {
    Title: formVal("typeName"),
    TemplateList: formVal("data_type_template_list"),
    TemplateItem: formVal("data_type_template_item"),
    Description: formVal("description"),
    Keywords: formVal("keywords"),
    URL: prettyurl,
    Fields:  TheFields,
    
  }

  //fmt.Fprintln(w,newType)


  //DB PUT
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "DataType", nil), &newType)
    
    //IF ERRORS
    if err != nil {
        fmt.Fprint(w,"error adding")
        return

    //NO ERRORS
    } else {

      
      cacheFlush("types",r)

      //DEBUG
      //fmt.Fprintln(w,"added successfully")
      //fmt.Fprintln(w,"key: " + key.Encode())

      
      //PREP JSON
      m := map[string]string{
        "message":"new type added",
        "key":key.Encode(),
        "title":newType.Title,
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



//FUNC CONTROL PAGES
func control_data_edit(w http.ResponseWriter, r *http.Request) {
  
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
    //fmt.Fprintln(w,itemID[4])
    return
  }

  /*******************************GET EDIT**************************/
  if r.Method == "GET"{

    //DB QUERY
    q := datastore.NewQuery("DataType").Filter("__key__ =", key).Limit(1)

    //VAR
    var bb []*DataType

    //DB GET ALL
    _,er := q.GetAll(c,&bb)
    
    //DB ERR
    if er != nil {    
      fmt.Fprintln(w, "error getting data type")
      return
    }

    /*
    //PREP JSON
    m := DataType {
      Title: bb[0].Title, //"Person" , 
 

    }*/

    //MARSHAL JSON
    j,errJSON := json.Marshal(bb)
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
    

  //INIT VAR
  TheFields := []FieldType {}

  //GET FIELD COUNT
  FieldCount,_ := strconv.Atoi(formVal("field_count")) 
  FieldCount = FieldCount - 1


    //FOR LOOP
    for i := 0; i <= FieldCount; i++ {

      //INT TO STR
      iCount := strconv.Itoa(i)

      //THE FIELDS
      TheFields = append(TheFields, FieldType {
        Name:   formVal("fieldName" + iCount),
        Order:  formVal("fieldOrder" + iCount),
        UI:     formVal("fieldUI" + iCount),
        //Errors: formVal("fieldErrors" + iCount),
      })

      //DEBUG
      //fmt.Fprintln(w,formVal("fieldName" + iCount))
      //fmt.Fprintln(w,formVal("fieldName0"))
      //fmt.Fprintln(w, FieldCount)
      //fmt.Fprintln(w,iCount)

    //END FOR  
    }

        //PRETTIFY URL
    reg, err := regexp.Compile("[^A-Za-z0-9]+")
    
    if err != nil {
        fmt.Fprintln(w,"error with RegX")
    }
    
    prettyurl := reg.ReplaceAllString(formVal("slug"), "-")
    prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))


    //fmt.Fprintln(w,prettyurl)


  //MAP FORM VALS
  newType := DataType {
    Title: formVal("typeName"),
    TemplateList: formVal("data_type_template_list"),
    TemplateItem: formVal("data_type_template_item"),
    Fields:  TheFields,
    Description: formVal("description"),
    Keywords: formVal("keywords"),
    URL: prettyurl,
  }

  //fmt.Fprintln(w,newType)


  //DB PUT
  //key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "DataType", nil), &newType)
  _,errPut := datastore.Put(c, key,&newType)

    //IF ERRORS
    if errPut != nil {
        fmt.Fprint(w,"error editing")
        return

    //NO ERRORS
    } else {


      //FLUSH CACHE
      cacheFlush("types",r)
      
      //fmt.Fprintln(w,"added successfully")
      //fmt.Fprintln(w,"key: " + key.Encode())


      //SETUP JSON
      m := map[string]string{
        "message":"updated type successfully",
        "title": newType.Title,
        //"key": key.Encode(),
        "adminSlug":AdminSlug,

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
      return

    //END NO ERRORS
    }


    /*
      //fmt.Fprintln(w,r)
      fmt.Fprintln(w,formVal("title"))
      //fmt.Fprintln(w,formVal("fieldlabel"))
      //fmt.Fprintln(w,r.FormValue("fieldlabel"))
     // fmt.Fprintln(w,r.PostForm)

      count := formVal("field_count")

      //num,_ := strconv.ParseInt(count,10,0)

      //for i := 0; i <= 2; i++ {
        //label := "fieldlabel" + string(i)
        //fmt.Fprintln(w, string(i))
        //fmt.Fprintln(w,formVal(label))
      //}


      fmt.Fprintln(w,count)

    //MAP FORM VALS
    newPage := &Data {
      Title: formVal("title"),
    }
 
    //DB PUT
    _,errPut := datastore.Put(c, key,newPage)
    
    //DB ERR
    if errPut != nil {
        fmt.Fprintln(w,"error Put")
        return
    }

    //SETUP JSON
    m := map[string]string{
      "message":"update success",
      "title": newPage.Title,
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

    */

  //END POST
  }
  
  
//END FUNC
}



//FUNC CONTROL PAGES
func control_data_delete(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  /********************************POST DELETE*********************************/
  if r.Method == "POST"{
  
    //CONTEXT
    c := appengine.NewContext(r)

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
      fmt.Fprintln(w,"delete successful")
      return
    }


  //END POST
  } else {
    fmt.Fprintln(w, "say whaaa")
  }
  
  
//END FUNC
}

