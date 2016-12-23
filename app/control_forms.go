package app


import (
  "net/http"
  "fmt"
  "appengine"
  "appengine/datastore"
  "strings"
  "strconv"
  "encoding/json"

)




/****************************GET LIST PAGES/POST ADD PAGES*********************************/
func control_forms(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET**************************/
  if r.Method == "GET"{
  
    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Forms",
     } 


    //QUERY
    q := datastore.NewQuery("Forms").Limit(100)


    //DB GET ALL
    var db []*Forms
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
  renderControl(w, r, "/control/forms.html", data, dbData)
}
  


/********************************POST*********************************/
} else {


  //GET FORM VALS
  formVal := func(val string)string{
    return r.FormValue(val)
  }

  //newFields := map[string]string {}
  
  //fmt.Fprintln(w,r)

  TheFields := []FieldType {}
  FieldCount,_ := strconv.Atoi(formVal("field_count")) 
 
    for i := 0; i <= FieldCount; i++ {

      iCount := strconv.Itoa(i)

      TheFields = append(TheFields, FieldType {
        Name:   formVal("fieldName" + iCount),
        //Label:  formVal("fieldLabel" + iCount),
        UI:     formVal("fieldUI" + iCount),
        //Errors: formVal("fieldErrors" + iCount),
      })

      fmt.Fprintln(w,formVal("fieldName" + iCount))
      fmt.Fprintln(w,formVal("fieldName0"))
      fmt.Fprintln(w, FieldCount)
      fmt.Fprintln(w,iCount)
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


  //MAP FORM VALS
  newType := DataType {
    Title: formVal("typeName"),
    TemplateList: formVal("data_type_template_list"),
    TemplateItem: formVal("data_type_template_item"),
    Fields:  TheFields,
  }

  fmt.Fprintln(w,newType)


  //DB PUT
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "DataType", nil), &newType)
    
    //IF ERRORS
    if err != nil {
        fmt.Fprint(w,"error adding")
        return

    //NO ERRORS
    } else {

            cacheFlush("types",r)


      fmt.Fprintln(w,"added successfully")
      fmt.Fprintln(w,"key: " + key.Encode())

      /*
      //PREP JSON
      m := map[string]string{
        "message":"new item added",
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
      */
      
    //END ERRORS
    }
 

  //END POST
  }
  
//END FUNC
}


//FUNC CONTROL PAGES
func control_forms_edit(w http.ResponseWriter, r *http.Request) {
  
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

  /*******************************GET**************************/
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


  /********************************POST*********************************/
  } else {
  
  
    //GET FORM VALS
    formVal := func(val string)string{
      return r.FormValue(val)
    }
    
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

    /*
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
func control_forms_delete(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  /********************************POST*********************************/
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

