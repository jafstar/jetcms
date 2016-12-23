package app


import (
  "net/http"
    //"net/url"

  //"fmt"
 
)

//SITEMAP
func sitemap(w http.ResponseWriter, r *http.Request)bool{


  //GET
  if r.Method == "GET"{
  
    settings :=  getSettings(w,r)
    pages := getPages(w,r)

    data := map[string]string {
      "get":"true",
      "title": "Sitemap",
      "analytics": settings["SiteAnalytics"],
      "cacheControl" : settings["CacheControl"],
      "cacheDays" : settings["CacheDays"],
      "cacheMonths": settings["CacheMonths"],
      "content":"<h1>Sitemap</h1>",
      "host": r.Host,
    }

    var filtered[]string

    //FILTER FOR HOMEPAGE
    for _,val := range pages {
      // fmt.Fprintln(w,  val)
      if val != "/" {
        filtered = append(filtered,val)
      }
    }

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
  
    
    renderXML(w, "misc/sitemap.xml", data,filtered)
    //fmt.Fprintln(w, r.RemoteAddr)

    
  } 

  return true
  

}