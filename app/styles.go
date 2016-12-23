package app


import (
  "net/http"
    //"net/url"
  "strings"
  //"fmt"
 
)

//SITEMAP
func styles(w http.ResponseWriter, r *http.Request) {


  //GET
  if r.Method == "GET"{

    path := r.URL.Path
    file := strings.Split(path,"/")
    name := strings.Split(file[2],".")

    //fmt.Fprintln(w,name[0])
  
    styles := getStyles(w,r,name[0])
  
    settings :=  getSettings(w,r)
  /*  pages := getPages(w,r)
*/
    data := map[string]string {
      "get":"true",
      "title": "Styles",
      "StyleSheet": styles["StyleSheet"],
      "cacheControl" : settings["CacheControl"],
      "cacheDays" : settings["CacheDays"],
      "cacheMonths": settings["CacheMonths"],
      "content":"<h1>Styles</h1>",
    }
/*
    var filtered[]string
*/
/*
    //FILTER FOR HOMEPAGE
    for _,val := range pages {
      // fmt.Fprintln(w,  val)
      if val != "/" {
        filtered = append(filtered,val)
      }
    }*/

      //PARSE URL
  //realURL,_ := url.Parse(r.RequestURI)
  //url := realURL.String()

  //DEBUG
  //l := len(url)
 	//fmt.Fprintln(w, "Clean URL: " + url)
 	//fmt.Fprintln(w,"Host : "+ r.Host)
 	//fmt.Fprintln(w,"Scheme: "+ r.URL.String())
 // fmt.Fprintln(w, "Real URL: " + realURL.Path)
 // fmt.Fprintln(w,"Host ": + realURL.Host)
  //fmt.Fprintln(w, l)
  
    
    renderCSS(w, "misc/styles.css",data)
    //fmt.Fprintln(w, r.RemoteAddr)

    
  } 

  //return true
  

}