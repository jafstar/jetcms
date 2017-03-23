package app


import (
  "net/http"
  "fmt"
  "appengine"
  "appengine/datastore"
  "strings"
  "encoding/json"
  //"os"
)




/****************************GET LIST PAGES/POST ADD PAGES*********************************/
func control_pages_list(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET ADD**************************/
  if r.Method == "GET"{
  
    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Pages",
     } 


    //QUERY
    q := datastore.NewQuery("page").Limit(100)


    //DB GET ALL
    var db []*Page
    keys,err := q.GetAll(c,&db)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting items")
      return
    }

    //VAR
    var dbData []map[string]string


    //CHECK IF DATA EXISTS
    if db != nil {

    //FOR DB ITEMS
    for i := range db {
      
      k := keys[i].Encode()

      dbData = append(dbData,
        map[string]string{
        "key":k,   
        "type":k,
         "title": db[i].Title,
         "pagetitle": db[i].PageTitle,
         "content":db[i].Content,
         "description": db[i].Description,
         "keywords": db[i].Keywords,
         "slug":db[i].Slug,

       },
      )
    }

    //ELSE INSERT DEFAULT
    } else {


        //DEFAULT VALS
        dbEntry := Page {
              Slug: "/",//formVal("slug"),
              Title: "Homepage",//formVal("title"),
              PageTitle:"Welcome to the homepage!",
              Description: "Default Jet CMS homepage",//formVal("description"),
              Keywords: "",//formVal("keywords"),
              Content: DEFAULT_PAGE_CONTENT,//"Welcome to the default Jet CMS homepage.",//formVal("content"),
              JS: "",//formVal("js"),
              CSS: "",//formVal("css"),
              Template: "",//formVal("template"),
              Single: "",//formVal("single"),
        }
        

        //DB PUT
        key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "page", nil), &dbEntry)
          
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
               "title":"Homepage",
               "adminSlug":AdminSlug,
                
             },
            )
          //END ELSE ERRORS
          }



    //END CHECK DATA
    }
    
    renderControl(w, r, "/control/pages.html", data, dbData)


/********************************POST ADD*********************************/
} else {


  //GET FORM VALS
  formVal := func(val string)string{
    return r.FormValue(val)
  }
  
  //MAP FORM VALS
  newPage := Page {
    Slug: formVal("slug"),
    Title: formVal("title"),
    PageTitle: formVal("pagetitle"),
    Description: formVal("description"),
    Keywords: formVal("keywords"),
    Content: formVal("content"),
    JS: formVal("js"),
    CSS: formVal("css"),
    Template:formVal("template"),
    Single:formVal("single"),
  }
  

  //DB PUT
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "page", nil), &newPage)
    
    //IF ERRORS
    if err != nil {
        fmt.Fprint(w,"error adding")
        return

    //NO ERRORS
    } else {

      cacheFlush("pages",r)


      //PREP JSON
      m := map[string]string{
        "message":"new item added",
        "key":key.Encode(),
        "title":newPage.Title,
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
func control_pages_edit(w http.ResponseWriter, r *http.Request) {
  
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
    q := datastore.NewQuery("page").Filter("__key__ =", key).Limit(1)

    //VAR
    var bb []*Page

    //DB GET ALL
    _,er := q.GetAll(c,&bb)
    
    //DB ERR
    if er != nil {    
      fmt.Fprintln(w, "error getting page")
      return
    }


    //PREP JSON
    m := Page {
      bb[0].Slug,
      bb[0].Title,
      bb[0].PageTitle,
      bb[0].Description,
      bb[0].Keywords,
      bb[0].Content,
      bb[0].JS,
      bb[0].CSS,
      bb[0].Template,
      bb[0].Single,
      //bb[0].List,
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
    newPage := &Page {
      Slug: formVal("slug"),
      Title: formVal("title"),
      PageTitle: formVal("pagetitle"),
      Description: formVal("description"),
      Keywords: formVal("keywords"),
      Content: formVal("content"),
      JS: formVal("js"),
      CSS: formVal("css"),
      Template:formVal("template"),
      Single:formVal("single"),
    }
 
    //DB PUT
    _,errPut := datastore.Put(c, key,newPage)
    
    //DB ERR
    if errPut != nil {
        fmt.Fprintln(w,"error Put")
        return
    } else {
        //FLUSH CACHE
        cacheFlush("pages",r)
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

  //END POST
  }
  
  
//END FUNC
}



//FUNC CONTROL PAGES
func control_pages_delete(w http.ResponseWriter, r *http.Request) {
  
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
      cacheFlush("pages",r)
      fmt.Fprintln(w,"delete successful")
      return
    }


  //END POST
  } else {
    fmt.Fprintln(w, "say whaaa")
  }
  
  
//END FUNC
}

