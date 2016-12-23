package app

import (
  "fmt"
  "strings"
  //"net/url"
  "net/http"
  "html/template"
  "appengine"
  "appengine/user"
  "strconv"
  //"appengine/blobstore"
  "time"
)

//RENDER FORM ERROR
func renderFormError(w http.ResponseWriter, r *http.Request,url string,errors map[string]string,settings map[string]string){
 
      //GET PAGE DATA
      data := getPageData(w,r,url)

      errors["title"] = "Errors in Form"
      errors["js"] =   data["js"]                      
      errors["css"] = data["css"]                       
      errors["keywords"] = data["keywords"]
      errors["description"] = data["description"]
      errors["class"] = data["class"]                   
      errors["analytics"] = settings["SiteAnalytics"]
      errors["content"] = data["content"]
 

    if(data["template"] == "default" || data["template"] == "" ){
      render(w, r,"core/page.html", errors,nil)
    } else {
     render(w,r, "core/" + data["template"], errors,nil)
    }

  
//END FUNC
}


//FUNC RENDER CSS
func renderCSS(w http.ResponseWriter, file string, data map[string]string){


 // t := time.Now().Format("2006")
  style := template.HTML(``+data["StyleSheet"]+``)

  html := map[string]template.HTML{}

  //data["Year"] = t
  html["style"] = style

  newView := cssView {
    CSS:html,
    //Data: data,
    //Pages:pages,
  }


  //PARSE
  htmlFiles,err := template.ParseFiles(
   "templates/" + file,
  )

  //ERROR
  if err != nil {
    fmt.Fprint(w, err.Error)
  }


  //TEMPLATES  
  tpl := template.Must(htmlFiles,err)
  

  cMonths,_ := strconv.Atoi(data["cacheMonths"])
  cDays,_ := strconv.Atoi(data["cacheDays"])
  
  //SET HEADERS
  w.Header().Set("Content-Type", "text/css;charset=utf-8")
  w.Header().Set("Cache-Control", data["cacheControl"])//"max-age:3600, public")
  w.Header().Set("Expires", time.Now().AddDate(0, cMonths, cDays).Format(http.TimeFormat))


  //EXECUTE
  tpl.Execute(w, newView)
                
//END FUNC
}

//FUNC RENDER XML
func renderXML(w http.ResponseWriter, file string, data map[string]string,pages []string){


  t := time.Now().Format("2006")
  content := template.HTML(``+data["content"]+``)

  html := map[string]template.HTML{}

  data["Year"] = t
  html["content"] = content

  newView := xmlView {
    HTML:html,
    Data: data,
    Pages:pages,
  }


  //PARSE
  htmlFiles,err := template.ParseFiles(
   "templates/" + file,
  )

  //ERROR
  if err != nil {
    fmt.Fprint(w, err.Error)
  }


  //TEMPLATES  
  header := template.Must(htmlFiles,err)
  
  cMonths,_ := strconv.Atoi(data["cacheMonths"])
  cDays,_ := strconv.Atoi(data["cacheDays"])
  
  //SET HEADERS
  w.Header().Set("Content-Type", "application/xml;charset=utf-8")
  w.Header().Set("Cache-Control", data["cacheControl"])//"max-age:3600, public")
  w.Header().Set("Expires", time.Now().AddDate(0, cMonths, cDays).Format(http.TimeFormat))

  //EXECUTE
  header.Execute(w, newView)
                
//END FUNC
}

func justSayIt(args ...interface{}) string {
    return HelloWorld
}



//TEMPLATE FUNC - FORMAT HTML
func FormatHTML(stuff string) template.HTML {
    h := template.HTML(``+stuff+``)
    return h
}




//TEMPLATE FUNC - FORMAT DATE
func FormatDate(t string) string {

  newTime,_ := time.Parse("2006-01-02",t)

  //return strconv.Itoa(newTime.Year())//newTime.Format("Jan 2, 2006")
  return newTime.Format(FormatDateDisplay)
}

//FUNC RENDER
func render(w http.ResponseWriter, r *http.Request,file string, data map[string]string, list []map[string]map[string]string){





  core := getCore(w,r)
  header := template.HTML(``+core["CoreHeader"]+``)
  footer := template.HTML(``+core["CoreFooter"]+``)

 htmlTemp := template.New("html content")
 tempHead := template.New("header")
 tempFoot := template.New("footer")

 //MUST COME BEFORE PARSE
  htmlTemp.Funcs(template.FuncMap{
    "justSayIt": justSayIt,
    "FormatDate":FormatDate,
    "FormatHTML":FormatHTML,
  })


  //PARSE
 htmlContent, _ := htmlTemp.Parse(data["content"])
 htmlHeader, _  := tempHead.Parse(core["CoreHeader"])
 htmlFooter, _  := tempFoot.Parse(core["CoreFooter"])
 




  //const layout = "2006"
  t := time.Now().Format("2006")
  ie := template.HTML(`<!--[if lt IE 10]><link rel='stylesheet' type='text/css' href='/css/ie.css' /><![endif]-->`)
  
  content := template.HTML(``+data["content"]+``)
  
  //tempContent,_ := template.New("html content").Parse(data["content"])
  //content := template.HTML(``+ htmlContent +``)
  //data["content"] = template.HTML(data["content"])

  html := map[string]template.HTML{}



  data["Year"] = t
  html["IE"] = ie 

  html["content"] = content
  html["Header"] = header
  html["Footer"] = footer
  //data["class"] = 

  newView := View {
    //Menus: menus,
    HTML:html,
    Data: data,
    List: list,
    //Site: site,
  }

/*

  //PARSE
  htmlFiles1,err1 := template.ParseFiles(
   //"templates/" + file,
   //"templates/header.html",
   //"templates/footer.html",
    //"templates/newHeader.html",
    "templates/plainHeader.html",
  )

  htmlFiles2,err2 := template.ParseFiles(
   //"templates/" + file,
   //"templates/header.html",
   //"templates/footer.html",
    //   "templates/newFooter.html",

   "templates/plainFooter.html",

  )



    
  //ERROR
  if err1 != nil || err2 != nil {
    fmt.Fprint(w, err1.Error)
  }

*/

  //MUST PASS CHECKS  
  //tpl1 := template.Must(htmlFiles1,nil)
  tpl1 := template.Must(htmlHeader,nil)
  tpl0 := template.Must(htmlContent,nil)
  tpl2 := template.Must(htmlFooter,nil)
  //tpl2 := template.Must(htmlFiles2,nil)


  
  /*
  body := template.Must(template.ParseFiles(
    "templates/" + file,
  ))
  
  
  footer := template.Must(template.ParseFiles(
    "templates/footer.html",
  ))*/

//fmt.Fprintln(w,"HUH: " + data["cacheControl"])

  cMonths,_ := strconv.Atoi(data["cacheMonths"])
  cDays,_ := strconv.Atoi(data["cacheDays"])
  
  //SET HEADERS
  w.Header().Set("Content-Type", "text/html;charset=utf-8")
  w.Header().Set("Cache-Control", data["cacheControl"])//"max-age:3600, public")
  w.Header().Set("Expires", time.Now().AddDate(0, cMonths, cDays).Format(http.TimeFormat))
  

  if(data["single"] == "1"){
    tpl0.Execute(w, newView)


  } else {
    //fmt.Fprintln(w,data["single"])
      //EXECUTE
    tpl1.Execute(w, newView)
    tpl0.Execute(w, newView)
    tpl2.Execute(w, newView)
  }




  //body.Execute(w, newView)
  //footer.Execute(w, newView)
  //template.ExecuteTemplate(w,"templates/header.html",newView)
                

  //NOTES
  //http://stackoverflow.com/questions/24837883/golang-templates-minus-function
  //http://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.4.html
}




//RENDER TYPE
func renderType(w http.ResponseWriter, r *http.Request,dType string,settings map[string]string,tplType string, dURL string){

  //VARS
  //dataItems := map[string]map[string]string{}

    //QUERIES
    if tplType == "list" {

          typeTemplate := getTypeTemplate(w,r, dType,tplType)
          dataItems := getListDataItems(w,r,dType)   

          //CONTENT DATA
          data := map[string]string{
            //"content":dataItems["content"]["description"],
            "content":typeTemplate,//theHTML, //typeTemplate,
            "class": dType, //dataItems[0]["info"]["url"],//strings.Trim(gg[0].Slug, "/"),//"page",
            "analytics":settings["SiteAnalytics"],
            "cacheControl" : settings["CacheControl"],
            "cacheDays" : settings["CacheDays"],
            "cacheMonths": settings["CacheMonths"],
            "title":strings.Title(dType),
            //"list": dataItems,
          }   

          render(w, r,"core/type.html", data, dataItems)

    //SINGLE ITEM
    } else {

            typeTemplate := getTypeTemplate(w,r, dType,tplType)
            dataItems := getSingleDataItem(w,r,dURL)
            

                //CONTENT DATA
          data := map[string]string{
            //"content":dataItems["content"]["description"],
            "content":typeTemplate,//theHTML, //typeTemplate,
            "class": dType, //dataItems[0]["info"]["url"],//strings.Trim(gg[0].Slug, "/"),//"page",
            "analytics":settings["SiteAnalytics"],
            "cacheControl" : settings["CacheControl"],
            "cacheDays" : settings["CacheDays"],
            "cacheMonths": settings["CacheMonths"],
            "title":strings.Title(dType),
            //"list": dataItems,
          }   



              //SINGLE ITEM - INDIVIDUAL FIELDS
          
            for key,val := range dataItems["content"] {
                 //fmt.Fprintln(w,key + " : " + val["Title"])
               data[key] = val
            }

          render(w, r,"core/type.html", data, nil)

           //DEBUG
          //fmt.Fprintln(w,data)
    }



   
    //for key,val := range dataItems["content"] {
         // fmt.Fprintln(w,key + " : "+val)
    //}

    //TEST DATA
   //theHTML := `<div><h1>{{.Data.title}}</h1> <h3>Date: {{.Data.date}} <br> Location: {{.Data.location}} <br> Description: {{.Data.description}}</h3></div>`
    //theHTML := `<div><h1>{{.Data.title}}</h1> <h3>Date: {{.Data.date}} <br> Location: {{.Data.location}} <br> Description: {{.Data.description}}</h3></div>`
 
    // theHTML := dataItems["info"]["template"]
  

  /*
    //TRIAL
    tplData := template.New("data_item")
    parsedData, _  := tplData.Parse(theHTML)
    
    tplX := template.Must(parsedData,nil)

    dataView := View {
      //Menus: menus,
     // HTML:html,
      Data: dataItems["content"],
      //Site: site,
    }
      //ie := template.HTML(`<!--[if lt IE 10]><link rel='stylesheet' type='text/css' href='/css/ie.css' /><![endif]-->`)

    tplX.Execute(w, dataView)
  */



    

        //data["analytics"]= settings["SiteAnalytics"]


    //SINGLE ITEM - INDIVIDUAL FIELDS
    /*
      for key,val := range dataItems[1]["content"] {
           //fmt.Fprintln(w,key + " : "+val)
          data[key] = val
      }
      */
  

/*
    //MULTIPLE ITEMS - BATCHED FIELDS
    for k1,_ := range dataItems{

      for key,val := range dataItems[k1]["content"] {
            fmt.Fprintln(w,key + " : "+val)
          //data[key] = val
      }

    }*/
      

   // fmt.Fprintln(w,dataItems)
    //fmt.Fprintln(w,data)

/*
    //RENDER
    if(data["template"] == "default" || data["template"] == "" ){
      render(w, r,"core/type.html", data, dataItems)
    } else {
      render(w, r,"core/type.html", data, dataItems)

      //render(w,r, "core/" + data["template"], data)
    }
  */

//END FUNC
}



//RENDER PAGE
func renderPage(w http.ResponseWriter, r *http.Request,t string,settings map[string]string){
 

 /*
  key, err := datastore.Put(c, datastore.NewIncompleteKey(c, "page", nil), &about)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
*/
     //var doc Page
    //if err = datastore.Get(c, key, &doc); err != nil {
      //  http.Error(w, err.Error(), http.StatusInternalServerError)
      //  return
    //}
    data := getPageData(w,r,t)

    data["analytics"] = settings["SiteAnalytics"]
    data["cacheControl"] = settings["CacheControl"]
    data["cacheDays"] = settings["CacheDays"]
    data["cacheMonths"] = settings["CacheMonths"]
 

    if(data["template"] == "default" || data["template"] == "" ){
      render(w, r,"core/page.html", data,nil)
    } else {
      render(w,r, "core/" + data["template"], data,nil)
    }
    
    //fmt.Fprintln(w, r.RemoteAddr)
    //fmt.Fprintln(w, ":)")


  
  
//END FUNC
}


//RENDER
func renderSingle(w http.ResponseWriter, file string, data map[string]string){

  //const layout = "2006"
  t := time.Now().Format("2006")



  data["Year"] = t


  newView := View {
    //Menus: menus,
    Data: data,
    //Site: site,
  }



  //TEMPLATES  
  /*
  header := template.Must(template.ParseFiles(
    "templates/header.html",
  ))
  
  body := template.Must(template.ParseFiles(
    "templates/" + file,
  ))
  */

  
  soon := template.Must(template.ParseFiles(
    "templates/"+file,
  ))
  
  cMonths,_ := strconv.Atoi(data["cacheMonths"])
  cDays,_ := strconv.Atoi(data["cacheDays"])
  
  //SET HEADERS
  w.Header().Set("Content-Type", "text/html;charset=utf-8")
  w.Header().Set("Cache-Control", data["cacheControl"])//"max-age:3600, public")
  w.Header().Set("Expires", time.Now().AddDate(0, cMonths, cDays).Format(http.TimeFormat))
  
  //EXECUTE
  //header.Execute(w, newView)
  //body.Execute(w, newView)
  soon.Execute(w, newView)



  /*
  //ERROR
  err := 
  if err != nil {
    fmt.Fprint(w, err)
  }
  */

}



//RENDER CONTROL
func renderControl(w http.ResponseWriter, r *http.Request, file string, data map[string]string, list []map[string]string){

  //ADMIN?
  //fmt.Fprint(w, r.Header["X-Appengine-Inbound-User-Is-Admin"] )
  // fmt.Fprint(w, r.TransferEncoding )
  //fmt.Fprint(w, r.Header["X-Requested-With"] )

  //CONTEXT
  c := appengine.NewContext(r)

  //const layout = "2006"
  t := time.Now().Format("2006")



  var options = map[string]string{
    "AdminSlug":AdminSlug,
    "Year": t,
  }


  newView := Content {
    Data: data,
    Globals: options,
    List: list,
  }
  
  //TEMPLATES  
  header := template.Must(template.ParseFiles(
    "templates/control/control_header.html",
  ))
  
  body := template.Must(template.ParseFiles(
    "templates/" + file,
  ))
  
  scripts := template.Must(template.ParseFiles(
    "templates/control/control_scripts.html",
  ))
  
  
  footer := template.Must(template.ParseFiles(
    "templates/control/control_footer.html",
  ))
  
  //SET HEADERS
  w.Header().Set("Content-Type", "text/html")
  

  //GET AJAX
  ajax := r.Header.Get("X-Requested-With")



  //CHECK AJAX
  if ajax  == "XMLHttpRequest" {

    //EXECUTE  
    body.Execute(w, newView)
    scripts.Execute(w, data)

  } else {
  
    url,_ := user.LogoutURL(c, "/")
    data["logout"] = url
    data["AdminSlug"] = AdminSlug
    data["Year"] = t
    //newView["Data"]["AdminSlug:AdminSlug"]

    //EXECUTE
    header.Execute(w, data)
    body.Execute(w, newView)
    footer.Execute(w, data)
  }




  /*
  //ERROR
  err := 
  if err != nil {
    fmt.Fprint(w, err)
  }
  */
  
  //fmt.Fprint(w, r.Host   )


}




/*

//RENDER CONTROL
func renderControl(w http.ResponseWriter, r *http.Request, file string, data map[string]string){

  //ADMIN?
  //fmt.Fprint(w, r.Header["X-Appengine-Inbound-User-Is-Admin"] )
  // fmt.Fprint(w, r.TransferEncoding )
  //fmt.Fprint(w, r.Header["X-Requested-With"] )


  //CONTEXT
  c := appengine.NewContext(r)
  
  //TEMPLATES  
  header := template.Must(template.ParseFiles(
    "templates/control/control_header.html",
  ))
  
  body := template.Must(template.ParseFiles(
    "templates/" + file,
  ))
  
  scripts := template.Must(template.ParseFiles(
    "templates/control/control_scripts.html",
  ))
  
  
  footer := template.Must(template.ParseFiles(
    "templates/control/control_footer.html",
  ))
  
  //SET HEADERS
  w.Header().Set("Content-Type", "text/html")
  

  //GET AJAX
  ajax := r.Header.Get("X-Requested-With")



  //CHECK AJAX
  if ajax  == "XMLHttpRequest" {

    //EXECUTE  
    body.Execute(w, data)
    scripts.Execute(w, data)

  } else {
  
    url,_ := user.LogoutURL(c, "/")
    data["logout"] = url
  
    //EXECUTE
    header.Execute(w, data)
    body.Execute(w, data)
    footer.Execute(w, data)
  }


  /*
  //ERROR
  err := 
  if err != nil {
    fmt.Fprint(w, err)
  }
  */
  
  //fmt.Fprint(w, r.Host   )


//}**/