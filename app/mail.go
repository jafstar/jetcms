package app

import (
  "fmt"
  "net/http"
  "appengine"
  "appengine/mail"
)


//SEND EMAIL
func  sendEmail(w http.ResponseWriter, r *http.Request, subject string, message string){


  var settings = getSettings(w,r)


  c := appengine.NewContext(r)



  msg := &mail.Message{
          //Sender: FROM_EMAIL,
          //To: []string{TO_EMAIL},
          Sender: settings["FromEmail"],
          To: []string{settings["ToEmail"]},
          Subject: subject, 
          Body: message,
        }
        
        if err := mail.Send(c, msg); err != nil {
        
          fmt.Fprint(w, err)
        
        } else {
        
         // fmt.Fprint(w, "email sent")
        
        }
        
      
}