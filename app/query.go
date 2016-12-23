package app

import (
  "net/http"
  "fmt"
 "appengine"
  "appengine/datastore"
  // "appengine/blobstore"
  "strings"
    "encoding/json"

)

//GET DATA URLs
func getDataURLs(w http.ResponseWriter, r *http.Request)[]string{

  urlItems :=  []string{}

  //urlItems = append(urlItems, "none")
  
  cache,err := cacheGet("dataURLs",r)

        c := appengine.NewContext(r)


  //fmt.Fprint(w,string(cache[:]))
  //fmt.Fprintf(w, "memcache: Val=%q", cache)
        c.Errorf("error getting cache data item url: %v", err)


  //IF CACHE EMPTY
 if cache == nil  {

     // p := []string{}
      q := datastore.NewQuery("DataItem").Filter("Published=",true).Limit(50)

      var gg []*DataItem
      _,err := q.GetAll(c,&gg)


      if err != nil {
        fmt.Fprint(w,"error getting types: " + err.Error())
        return nil
      }

        for key := range gg {
          //fmt.Fprintln(w, gg[key].Slug)
         urlItems = append(urlItems,strings.ToLower(gg[key].URL))
          //pages = [gg[key].Title]: gg[key].Slug,
        }

        cacheAdd("dataURLs",r,urlItems)

        //return p
          return urlItems

    //USE CACHE
    } else {
      urls := strings.Split(string(cache.Value[:]), ",")
      //c.Infof(strings.Join(urls,", "))

      return urls
    
    }


//END FUNC 
}

//GET TYPE TEMPLATE
func getTypeTemplate(w http.ResponseWriter, r *http.Request, dataType string, tplType string) string {

  //fmt.Fprintln(w,dataType)
    //CONTEXT
    c := appengine.NewContext(r)

     q := datastore.NewQuery("DataType").Filter("Title =",strings.Title(dataType)).Limit(1)

     var gg []*DataType
    _,err := q.GetAll(c,&gg)

    if err != nil {
       fmt.Fprint(w,"error getting type template: " + err.Error())
      //return ""
    }

    //dbHTML := gg[0].Template
    theHTML := ""
    //fmt.Fprintln(w,dbHTML)
    if tplType == "list" {
      theHTML = gg[0].TemplateList //`<div><h1>{{.Data.title}}</h1> <h3>Date: {{.Data.date}} <br> Location: {{.Data.location}} <br> Description: {{.Data.description}}</h3></div>`
    } else {
      theHTML = gg[0].TemplateItem
    }


    return theHTML 

//END FUNC 
}



//GET TYPES
func getTypes(w http.ResponseWriter, r *http.Request) []string {

  cache,err := cacheGet("types",r)

        c := appengine.NewContext(r)


  //fmt.Fprint(w,string(cache[:]))
  //fmt.Fprintf(w, "memcache: Val=%q", cache)
        c.Errorf("error getting cache types item: %v", err)


  //IF CACHE EMPTY
  if cache == nil  {

      p := []string{}
      q := datastore.NewQuery("DataType").Limit(100)

      var gg []*DataType
      _,err := q.GetAll(c,&gg)


      if err != nil {
        fmt.Fprint(w,"error getting types: " + err.Error())
        return nil
      }

        for key := range gg {
          //fmt.Fprintln(w, gg[key].Slug)
          p = append(p,strings.ToLower(gg[key].Title))
          //pages = [gg[key].Title]: gg[key].Slug,
        }

        cacheAdd("types",r,p)

        return p

    //USE CACHE
    } else {
      urls := strings.Split(string(cache.Value[:]), ",")
      //c.Infof(strings.Join(urls,", "))

      return urls
    
    }

//END FUNC   
}

//GET PAGES
func getPages(w http.ResponseWriter, r *http.Request) []string {

  cache,err := cacheGet("pages",r)

        c := appengine.NewContext(r)


  //fmt.Fprint(w,string(cache[:]))
  //fmt.Fprintf(w, "memcache: Val=%q", cache)
        c.Errorf("error getting cache pages item: %v", err)


  //IF CACHE EMPTY
  if cache == nil  {

      p := []string{}
      q := datastore.NewQuery("page").Limit(100)

      var gg []*Page
      _,err := q.GetAll(c,&gg)


      if err != nil {
        fmt.Fprint(w,"error getting pages: " + err.Error())
        return nil
      }

        for key := range gg {
          //fmt.Fprintln(w, gg[key].Slug)
          p = append(p,gg[key].Slug)
          //pages = [gg[key].Title]: gg[key].Slug,
        }

        cacheAdd("pages",r,p)

        return p

    //USE CACHE
    } else {
      urls := strings.Split(string(cache.Value[:]), ",")
      //c.Infof(strings.Join(urls,", "))

      return urls
    
    }

//END FUNC   
}


//GET PAGE DATA
func getPageData(w http.ResponseWriter, r *http.Request,t string) map[string]string{

    c := appengine.NewContext(r)

    url := t


     q := datastore.NewQuery("page").Filter("Slug =",url).Limit(10)

          //fmt.Fprintln(w, q)

     var gg []*Page
    _,err := q.GetAll(c,&gg)

    if err != nil {
       fmt.Fprint(w,"error getting page data: " + err.Error())
      return nil
    }

    if gg != nil {

    data := map[string]string{
      "get": "true",
      //"id":gg[0].Key,
      "title": gg[0].Title,
      "description": gg[0].Description,
      "keywords": gg[0].Keywords,
      "content":gg[0].Content,
      "js": gg[0].JS, //cdn["ms"]+"/jquery.validate/1.9/jquery.validate.min.js",
      "css":gg[0].CSS,
      "template":gg[0].Template,
      "single":gg[0].Single,

      "class": strings.Trim(gg[0].Slug, "/"),//"page",

    }


    return data
  } else {
    return nil
  }

//END FUNC
}

//GET DATA ITEMS
func getSingleDataItem(w http.ResponseWriter, r *http.Request,dURL string) map[string]map[string]string {

    c := appengine.NewContext(r)

    //url := dType


     q := datastore.NewQuery("DataItem").Filter("URL =",dURL).Filter("Published =",true).Limit(1)

          //fmt.Fprintln(w, q)

     var gg []*DataItem
    _,err := q.GetAll(c,&gg)

    if err != nil {
       fmt.Fprint(w,"error getting type data: " + err.Error())
      return nil
    }


      //fmt.Fprintln(w,jsonContent)

    //IF NOT NIL
    if gg != nil {

        data := map[string]map[string]string{}

        //FOR LOOP
        //for key := range gg {


          //UNMARSHALL
          var jsonContent = map[string]string{}
          json.Unmarshal(gg[0].Content, &jsonContent)
        

          //DATA
          data = map[string]map[string]string{
          //data = append(data,map[string]map[string]string {
            "info":map[string]string{
                "title":gg[0].Title, //strings.Title(dType),
                "url":gg[0].URL,
                //"dateOrder":gg[0].DateOrder,
                //"template":gg[0].Template,
                "type":gg[0].Type,
              },
            "content": jsonContent,

          //END DATA
         // })
          }
        //END FOR
       // }

      //fmt.Fprintln(w,jsonContent)    

    return data

  //ELSE IS NIL
  } else {

    return nil

  //END ELSE
  }

//END FUNC
}

//GET DATA ITEMS
func getListDataItems(w http.ResponseWriter, r *http.Request,dType string) []map[string]map[string]string {

    c := appengine.NewContext(r)

    //url := dType


     q := datastore.NewQuery("DataItem").Filter("Type =",strings.ToUpper(dType)).Filter("Published =", true).Order("-DateOrder").Limit(100)

          //fmt.Fprintln(w, q)

     var gg []*DataItem
    _,err := q.GetAll(c,&gg)

    if err != nil {
       fmt.Fprint(w,"error getting type data: " + err.Error())
      return nil
    }


      //fmt.Fprintln(w,jsonContent)

    //IF NOT NIL
    if gg != nil {

        data := []map[string]map[string]string{}

        //FOR LOOP
        for key := range gg {


          //UNMARSHALL
          var jsonContent = map[string]string{}
          json.Unmarshal(gg[key].Content, &jsonContent)
        

          //DATA
          //data := map[string]map[string]string{
          data = append(data,map[string]map[string]string {
            "info":map[string]string{
                "title":gg[key].Title, //strings.Title(dType),
                "url":gg[key].URL,
                //"template":gg[0].Template,
                "type":gg[key].Type,
              },
            "content": jsonContent,

          //END DATA
          })

        //END FOR
        }

    return data

  //ELSE IS NIL
  } else {

    return nil

  //END ELSE
  }

//END FUNC
}

//GET SETTINGS
func getSettings(w http.ResponseWriter, r *http.Request) map[string]string {

  //CONTEXT
  c := appengine.NewContext(r)

  //MAP LIST
  list := []string {
     "SiteTitle",
     "SiteClosed",
     "SiteAnalytics",

     "HomeTitle",
     "HomeDesc",
     "HomeKeyW",

     "SoonTitle",
     "SoonDesc",
     "SoonKeyW",
     "SoonURL",

     "CacheControl",
     "CacheDays",
     "CacheMonths",

     "AdminEmail",
     "FromEmail",
     "ToEmail",
  }

  //GET CACHE
  cache,err := cacheGetMulti(list,r)
  c.Errorf("there is a query error getting items: %v", err)
  c.Errorf("items: %v", len(cache))


  //IF CACHE EMPTY
  if len(cache) != len(list) {

      //QUERY
    q := datastore.NewQuery("settings").Order("-Last_Update").Limit(1)


    //DB GET ALL
    var db []*Settings 
    keys,err := q.GetAll(c,&db)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting settings: " + err.Error())
      return nil
    }


    //VAR
    dbData := map[string]string{}

    //fmt.Fprintln(w,keys)

    
    //FOR DB ITEMS
    for i := range db {
      
      //KEYS ENCODE
      k := keys[i].Encode()

      dbData = map[string]string {
         //"title": db[i].Filename,
         "key":k,
         "SiteTitle":db[i].Site_Title,
         "SiteClosed":db[i].Site_Closed,
         "SiteAnalytics":db[i].Site_Analytics,

         "HomeTitle":db[i].Home_Title,
         "HomeDesc":db[i].Home_Desc,
         "HomeKeyW":db[i].Home_KeyW,

         "SoonTitle":db[i].Soon_Title,
         "SoonDesc":db[i].Soon_Desc,
         "SoonKeyW":db[i].Soon_KeyW,
         "SoonURL":db[i].Soon_URL,

         "CacheControl":db[i].Cache_Control,
         "CacheDays":db[i].Cache_Days,
         "CacheMonths":db[i].Cache_Months,

         "AdminEmail":db[i].Admin_Email,
         "FromEmail":db[i].From_Email,
         "ToEmail":db[i].To_Email,

       }
      
    }

    //ADD CACHE
    cacheAddMap(r,dbData)


    return dbData

    //USE CACHE
    } else  {
      
      //VAR
      dbData := map[string]string{}

          
      dbData = map[string]string {
      
         "SiteTitle":string(cache["SiteTitle"].Value[:]),
         "SiteClosed":string(cache["SiteClosed"].Value[:]),
         "SiteAnalytics":string(cache["SiteAnalytics"].Value[:]),

         "HomeTitle":string(cache["HomeTitle"].Value[:]),
         "HomeDesc":string(cache["HomeDesc"].Value[:]),
         "HomeKeyW":string(cache["HomeKeyW"].Value[:]),

         "SoonTitle":string(cache["SoonTitle"].Value[:]),
         "SoonDesc":string(cache["SoonDesc"].Value[:]),
         "SoonKeyW":string(cache["SoonKeyW"].Value[:]),
          "SoonURL":string(cache["SoonURL"].Value[:]),

          "CacheControl":string(cache["CacheControl"].Value[:]),
          "CacheDays":string(cache["CacheDays"].Value[:]),
          "CacheMonths":string(cache["CacheMonths"].Value[:]),

         "AdminEmail":string(cache["AdminEmail"].Value[:]),
         "FromEmail":string(cache["FromEmail"].Value[:]),
         "ToEmail":string(cache["ToEmail"].Value[:]),

         }

      return dbData

    //END ELSE
    }

//END FUNC
}

//GET SETTINGS
func getCore(w http.ResponseWriter, r *http.Request) map[string]string {

    //CONTEXT
  c := appengine.NewContext(r)

  //MAP LIST
  list := []string {
    "CoreHeader",
    "CoreFooter",
  }

  //GET CACHE
  cache,err := cacheGetMulti(list,r)

  c.Errorf("there is a query error getting items: %v", err)
  c.Errorf("items: %v", len(cache))

  //fmt.Fprint(w,"the error is: %v",err)

  /*STRANGE BUG BEHAVIOR* 
    could not use "if cache == nil" because cache was empty but not nil...
    and err was not nil on the getCache but is here...
  */

  //IF CACHE EMPTY
  if len(cache) != len(list)  {

      //QUERY
      q := datastore.NewQuery("core").Order("-Last_Update").Limit(1)


      //DB GET ALL
      var db []*Core 
      keys,err := q.GetAll(c,&db)

      //DB ERR
      if err != nil {
        fmt.Fprint(w,"error getting core: " + err.Error())
        return nil
      }


      //VAR
      dbData := map[string]string{}

      //fmt.Fprintln(w,keys)

      
      //FOR DB ITEMS
      for i := range db {
        
        //KEYS ENCODE
        k := keys[i].Encode()

        dbData = map[string]string {
           //"title": db[i].Filename,
           "key":k,
           "CoreHeader":db[i].Core_Header,
           "CoreFooter":db[i].Core_Footer,
           "CoreVersion":db[i].Core_Version,

         }
        
      }

      cacheAddMap(r,dbData)

      return dbData

    //USE CACHE
    } else {

      //VAR
      dbData := map[string]string{}


          
        dbData = map[string]string {

           "CoreHeader":string(cache["CoreHeader"].Value[:]),
           "CoreFooter":string(cache["CoreFooter"].Value[:]),

         }

      return dbData
    }

}

//GET SETTINGS
func getStyles(w http.ResponseWriter, r *http.Request,file string) map[string]string {

  //CONTEXT
  c := appengine.NewContext(r)

  //CACHE LIST
  list := []string {
   "StyleName_"+file,
   "StyleSheet_"+file,
  }

  //GET CACHE
  cache,err := cacheGetMulti(list,r)

  c.Errorf("there is a query error getting items: %v", err)
  c.Errorf("items: %v", len(cache))


  //IF CACHE EMPTY
  if len(cache) != len(list) {

      //QUERY
      q := datastore.NewQuery("style").Filter("Style_Name =",file).Order("-Last_Update").Limit(1)


      //DB GET ALL
      var db []*Style 
      keys,err := q.GetAll(c,&db)

      //DB ERR
      if err != nil {
        fmt.Fprint(w,"error getting styles: " + err.Error())
        return nil
      }


      //VARS
      dbData := map[string]string{}
      cacheData := map[string]string{}

      
      //FOR DB ITEMS
      for i := range db {
        
        //KEYS ENCODE
        k := keys[i].Encode()

        dbData = map[string]string {
           //"title": db[i].Filename,
           "key":k,
           "StyleName" : db[i].Style_Name,
           "StyleSheet" : db[i].Style_Sheet,

         }

        cacheData = map[string]string {
          "StyleName_" + db[i].Style_Name : db[i].Style_Name,
          "StyleSheet_" + db[i].Style_Name : db[i].Style_Sheet,
        }
        
      }

 

      //ADD CACHE
      cacheAddMap(r,cacheData)


      return dbData

    //USE CACHE
    } else {

      //VAR
      dbData := map[string]string{}

      //MAP
      dbData = map[string]string {
        "StyleName": string(cache["StyleName_" + file].Value[:]),
        "StyleSheet": string(cache["StyleSheet_" + file].Value[:]),
      }

      return dbData

    //END ELSE
    }

}



//GET IMAGE
func getImageKey(w http.ResponseWriter, r *http.Request, filename string) string {



  c := appengine.NewContext(r)

  cache,err := cacheGet("image_" + filename,r)

  //fmt.Fprint(w,string(cache[:]))
  //fmt.Fprintf(w, "memcache: Val=%q", cache)
  c.Errorf("error getting cache image item: %v", err)


  //IF CACHE EMPTY
  if cache == nil  {



     // p := []string{}
      q := datastore.NewQuery("__BlobInfo__").Filter("filename =",filename).Limit(1)

      var gg []*BlobInfo
      keys,err := q.GetAll(c,&gg)


      if err != nil {
        fmt.Fprint(w,"error getting image: " + err.Error())
        //return nil
      }

      k := keys[0].StringID()

      //DEBUG
      //fmt.Fprintln(w,keys[0])
      //fmt.Fprintln(w,keys[0].StringID())
      //fmt.Fprintln(w,keys[0].Kind())
      //fmt.Fprintln(w, gg[0].Size)
      //fmt.Fprintln(w, filename)
      
    
       // for key := range gg {
          //fmt.Fprintln(w, string(gg[key].Filename))
        //  p = append(p,string(gg[key].BlobKey))
          //pages = [gg[key].Title]: gg[key].Slug,
        //}
      


        cacheAddImg("image_" + filename,r,k)

       return k


    //USE CACHE
    } else {

      urls := string(cache.Value[:])
      //c.Infof(strings.Join(urls,", "))

      return urls
    
    }

    //NOTES
    //https://cloud.google.com/appengine/docs/go/datastore/reference#Key.Encode
    //https://code.google.com/p/appengine-go/source/browse/appengine/blobstore/blobstore.go

//END FUNC   
}
