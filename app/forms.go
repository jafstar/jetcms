package app

import (
  "net/http"
  "fmt"
 
)

//CONTACT
func forms(w http.ResponseWriter, r *http.Request)bool{
  

  //GET
  if r.Method == "GET"{

    /*
  
    settings :=  getSettings(w,r)


  
    data := map[string]string {
      "get":"true",
      "title": "Contact",
      "description": "Contact Us",
      "keywords": "contact, contact form",
      "js": "/js/libs/validate.js", //cdn["ms"]+"/jquery.validate/1.9/jquery.validate.min.js",
      "css":"/css/form.css",
      "class":"contact",
      "analytics": settings["SiteAnalytics"],
    }
  
    
    render(w,r, "core/contact.html", data)
    //fmt.Fprintln(w, r.RemoteAddr)
    */
    //fmt.Fprintln(w,"GETTING...")

    fmt.Fprintln(w, r.URL)


  //POST  
  } else {

    //fmt.Fprintln(w,"POSTING...")

  
 
    //VARS MAP
    errorMap := map[string]string{}
    data  := map[string]string{}


    //FORM MAP
    formMap := map[string]string{
      "Name": r.FormValue("name"),
      "Email": r.FormValue("email"),
      "Phone": r.FormValue("phone"),
      "Message": r.FormValue("message"),    
    }

    //GET SETTINGS
    settings := getSettings(w,r)


    //CHECK FORM
    errors := checkForm(formMap, errorMap)
    
    //EMAIL VARS
    //subject := site["name"] + " - Contact Form"
    subject := settings["SiteTitle"] + " -  Contact Form"
    //subject := "Contact Form"
    message := ""
    
    //FOR
    for key := range formMap {
      message = message + "\n" +  formMap[key]
    }
    
    /*
    //DEBUG
    fmt.Fprintln(w,errors)
    fmt.Fprintln(w, settings)
    fmt.Fprintln(w, errorMap)
    fmt.Fprintln(w, data)
    fmt.Fprintln(w, formMap)
    */
    //return true

    
    
    //IF NO ERRORS
    if len(errors) == 0 {

      //SEND EMAILS
      sendEmail(w,r,subject,message)

      data["thanks"] = "true"
      data["title"] = "Thank you"
      data["content"] = "Thank you, the form has been submitted."
      data["message"] = "Thank you, the form has been submitted."
      data["name"] = formMap["Name"]
      data["css"] = "styles/desktop.css"
      
      
      render(w,r, "misc/thanks.html", data,nil)
    
    //ELSE ERRORS
    } else {

        //fmt.Fprintln(w,"THERE ARE FORM ERRORS")

          
        /*
      errors["errors"] = "true"
      errors["title"] = "Errors in Form"
      //data["token"] = token(r)
      errors["js"] = "/js/libs/validate.js" //cdn["ms"]+"/jquery.validate/1.9/jquery.validate.min.js"
      errors["css"] = "/styles/desktop.css"
      errors["class"] = "contact"
      errors["content"] = "<br>"
      */
     // render(w,r,  /* r.URL.String()"core/contact.html"*/ "core/page.html", errors)
      renderFormError(w,r,"/contact",errors,settings)

    
    //END ERRORS
    } 

  
  
  //END POST
  }
  
  return true
  
//END FUNC
}



//CONTACT FORM
func checkForm(form map[string]string, errors map[string]string)map[string]string {
  
  
  //NOT EMPTY
  for key, _ := range form {
    if form[key] == "" {
      errors[key+"Empty"] = key + " is empty"
    } 
  }
  
  //NOT EMAIL
  if !isEmail(form["Email"]) && errors["EmailEmpty"] == ""  {
    errors["NotEmail"] = "Invalid email address"
  }
  
  //NOT PHONE
  if !isPhone(form["Phone"]) && errors["PhoneEmpty"] == "" {
    errors["NotPhone"] = "Invalid phone number"
  }
  
  //ERRORS
  if len(errors) != 0 {
  
    //REMAP VALUES
    for key, val := range form {
      errors["val"+key] = val
    }
  }
 
  //RETURN
  return errors

}
