package app

import(
  "net/http"
  "appengine"
  "appengine/user"
)

func getCurrentUser(w http.ResponseWriter,r *http.Request)string{
    c := appengine.NewContext(r)
    u := user.Current(c)

      if u == nil {
 

    return "Admin"
  
  } else {

      return u.String()
    
  }


}


//CHECK USER
func checkUser(w http.ResponseWriter,r *http.Request){
  c := appengine.NewContext(r)
    u := user.Current(c)
    if u == nil {
        url, err := user.LoginURL(c, r.URL.String())
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Location", url)
        w.WriteHeader(http.StatusFound)
        return
    }
}


//CHECK ADMIN
func checkAdmin(w http.ResponseWriter,r *http.Request){
  c := appengine.NewContext(r)
 
  u := user.IsAdmin(c)
 
  if u == false {
 
    url, err := user.LoginURL(c, r.URL.String())
    
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Location", url)
    w.WriteHeader(http.StatusFound)
    return
  
  }
}