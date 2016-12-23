package app

import (
  "net/http"
 // "net/url"
 // "sync" 
  //"fmt"
  //"appengine"
  //"appengine/datastore"
   // "appengine/user"

  //"muncheez/app/data"
)

func displayHome(w http.ResponseWriter, r *http.Request, settings map[string]string) bool{

             homepage := getPageData(w,r,"/")

             if homepage == nil {

    	           data := map[string]string{
                  "title": HOME_TITLE,
                  "description": HOME_DESC,
                  "keywords": HOME_KEYW,
                  //"title":settings["HomeTitle"],
                  //"description":settings["HomeDesc"],
                  //"keywords":settings["HomeKeyW"],
                  "analytics":settings["SiteAnalytics"],
                  "cacheControl" : settings["CacheControl"],
                  "cacheDays" : settings["CacheDays"],
                  "cacheMonths": settings["CacheMonths"],
                  "content":"<p>Default Homepage</p>",
                  "class":"home",
                  "css":"/css/home.css",
                }

                //render(w,r,"core/home.html",data)
                render(w,r,"misc/default.html",data,nil)

            } else {


                data := map[string]string{
                  //"title": TITLE_HOME,
                  //"description": DESC_HOME,
                  //"keywords": KEYW_HOME,
                  //"title":settings["HomeTitle"],
                  //"description":settings["HomeDesc"],
                  //"keywords":settings["HomeKeyW"],
                  "title":settings["SiteTitle"],
                  "description":homepage["description"],
                  "keywords":homepage["keywords"],

                  "analytics":settings["SiteAnalytics"],
                  "cacheControl" : settings["CacheControl"],
                  "cacheDays" : settings["CacheDays"],
                  "cacheMonths": settings["CacheMonths"],
                  "class":"home",
                  //"css":"/css/home.css",
                  "css":homepage["css"],
                  "content":homepage["content"],
                }

                render(w,r,"core/page.html",data,nil)

                //renderPage(w,r,"/",settings)


            }
            return true

}

func displayPage(w http.ResponseWriter, r *http.Request, pages []string, url string, settings map[string]string) bool {
	 //FOR PAGE
	 disp := false

          for _,val := range pages {
            //if val == "/"+slug {
            if val == url {
              //fmt.Fprintln(w,  val)
              renderPage(w,r,val,settings)
              disp = true
              //return
            } 

          //END FOR  
          }

          return disp

}

func displaySoon(w http.ResponseWriter, settings map[string]string) bool {
      data := map[string]string{
        "title": SOON_TITLE,
        "description": SOON_DESC,
        "keywords": SOON_KEYW,
        //"title":settings["SoonTitle"],
        //"description":settings["SoonDesc"],
        //"keywords":settings["SoonKeyW"],
        "analytics":settings["SiteAnalytics"],
        "cacheControl" : settings["CacheControl"],
        "cacheDays" : settings["CacheDays"],
        "cacheMonths": settings["CacheMonths"],
      }

      renderSingle(w, "/misc/soon.html", data)
      	return true
}

func display404(w http.ResponseWriter, settings map[string]string){
	  	data := map[string]string{
    	"title": TITLE_404,
    	"description": DESC_404,
    	"keywords": KEYW_404,
      "analytics":settings["SiteAnalytics"],
      "cacheControl" : settings["CacheControl"],
      "cacheDays" : settings["CacheDays"],
      "cacheMonths": settings["CacheMonths"],
  	}

      //fmt.Fprintln(w, "404 Not Found")
      renderSingle(w,"/misc/404.html", data)
}