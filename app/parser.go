package app

import (
  "net/http"
  "net/url"
  //"sync" 
  //"fmt"
  "strings"
  "appengine"
  //"appengine/datastore"
  "appengine/user"
)

//GLOBALS
//var startOnce sync.Once 



//HOME
func parser(w http.ResponseWriter, r *http.Request) {

  //VARS
  display := false

  //PARSE URL
  realURL,_ := url.Parse(r.RequestURI)
  url := realURL.String()
  splitURL := strings.Split(strings.Trim(url,"/"),"/")

  //GET
  if r.Method == "GET"{

            //QUERIES
            pages := getPages(w,r)
            types := getTypes(w,r)
            dataURLs := getDataURLs(w,r)
            settings :=  getSettings(w,r)

     
            /*
            //DEBUG
            //l := len(url)
           	fmt.Fprintln(w, "Clean URL: " + url)
            fmt.Fprintln(w,splitURL)
            fmt.Fprintln(w,len(splitURL))
            //fmt.Fprintln(w, "Real URL: " + realURL.Path)
            //fmt.Fprintln(w, l)
            fmt.Fprintln(w,pages)
            fmt.Fprintln(w,types)
            */

          //IF SITE CLOSED
          if settings["SiteClosed"] == "offline" {

            //NEW CONTEXT
            c := appengine.NewContext(r)

            //GET ADMIN
            u := user.IsAdmin(c)

              //NOT ADMIN
              if u == false {
            

                  display = displayPage(w,r,pages,settings["SoonURL"],settings)

                  if display == false{
                     display = displaySoon(w,settings)
                  }


              //IS ADMIN
              } else {


                  //IF HOMEPAGE
                  if url == "/" || url == "" {

                    display = displayHome(w,r,settings)
                  //CONTACT
                  //} else if url == "/submitForm" {
                    //  display = forms(w,r)
                  //SITEMAP
                  } else if url == "/sitemap.xml" {
                        display = sitemap(w,r)

                  //OTHER PAGES
                  } else {

                    

                    display = displayPage(w,r,pages,url,settings)


                        //IF FALSE
                        if display == false {

                          //LOOP 1
                          for _,val := range types {

                            //CHECK DATA TYPE URLS
                            if val == splitURL[0] {
                              //SINGLE
                              if(len(splitURL) > 1){
                                //fmt.Fprintln(w,"url: " + splitURL[1] )
                                //LOOP 2
                                for _,val2 := range dataURLs {

                                  //CHECK DATA ITEM URLS
                                  if val2 == splitURL[1]{
                                    renderType(w,r,val,settings,"single",splitURL[1])
                                    display = true

                                  } 

                                //END LOOP 2
                                }

                              //LIST
                              } else {
                                renderType(w,r,val,settings,"list","")
                                display = true

                              }

                            //END IF DATA TYPE
                            }

                          //END LOOP 1  
                          }

                        //END IF FALSE
                        }

                  //END ELSE OTHER PAGES
                  }

              //END ELSE ADMIN
              }
              	

          //SITE OPEN
          } else {

              

          		//IF HOMEPAGE
          		if url == "/" || url == "" {

                  display = displayHome(w,r,settings)
          		
              //CONTACT
              //} else if url == "/submitForm" {
                //  display = forms(w,r)
              //SITEMAP
              } else if url == "/sitemap.xml" {
                    display = sitemap(w,r)
          		//OTHER PAGES
          		} else {

                    //fmt.Fprintln(w,dataURLs)

                    
                        display = displayPage(w,r,pages,url,settings)

                        //IF FALSE
                        if display == false {

                          //LOOP 1
                          for _,val := range types {

                            //CHECK DATA TYPE URLS
                            if val == splitURL[0] {
                              //SINGLE
                              if(len(splitURL) > 1){
                                //fmt.Fprintln(w,"url: " + splitURL[1] )
                                //LOOP 2
                                for _,val2 := range dataURLs {

                                  //CHECK DATA ITEM URLS
                                  if val2 == splitURL[1]{
                                    renderType(w,r,val,settings,"single",splitURL[1])
                                    display = true

                                  } 

                                //END LOOP 2
                                }

                              //LIST
                              } else {
                                renderType(w,r,val,settings,"list","")
                                display = true

                              }

                            //END IF DATA TYPE
                            }

                          //END LOOP 1  
                          }

                        //END IF FALSE
                        }
                    

          		//END ELSE OTHER PAGES
          		}


            //END IF SITE CLOSED		
            }

            //CHECK 404
            if display == false {
              display404(w,settings)
            }

//ELSE POST
} else {


 // fmt.Fprintln(w,"POST REQ")
  //if url == r.URL.String() {
      display = forms(w,r)
 
  //} else {
  //  fmt.Fprintln(w,"did not equal...")
  //}


//END POST
}

//END FUNC
}
