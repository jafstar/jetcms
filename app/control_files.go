package app


import (
  "net/http"
  //"net/url"
  "fmt"
  "appengine"
  "appengine/datastore"
  "appengine/blobstore"
  //"encoding/csv"
  "appengine/image"

  "strings"
  //"strconv"
  "encoding/json"
  "time"
)

type UploadURLOptions struct {
    MaxUploadBytes        int64 // optional
    MaxUploadBytesPerBlob int64 // optional

    // StorageBucket specifies the Google Cloud Storage bucket in which
    // to store the blob.
    // This is required if you use Cloud Storage instead of Blobstore.
    // Your application must have permission to write to the bucket.
    // You may optionally specify a bucket name and path in the format
    // "bucket_name/path", in which case the included path will be the
    // prefix of the uploaded object's name.
    StorageBucket string
}

//var queryType = "data"

/****************************GET UPLOAD URL*********************************/
func control_files_upload_url(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET URL**************************/
  if r.Method == "GET"{

/*
    var uploadOptions *blobstore.UploadURLOptions;

    uploadOptions {
      StorageBucket: "staging.jetcmsapp.appspot.com",
    }*/
  
     uploadURL, errURL := blobstore.UploadURL(c, "/"+AdminSlug+"/files/", nil)

     if errURL != nil {
      fmt.Fprintln(w,"error with making upload URL on GET:" + errURL.Error())
      return
     }


    //SETUP JSON DATA
     data := map[string]string{
       "get":"true",
       "uploadURL":uploadURL.String(), 
     }

          
        //MARSHAL JSON
        j,errJSON := json.Marshal(data)

        if errJSON != nil {
          fmt.Fprintln(w,"error with JSON")
        }

        //SET CONTENT-TYPE
        w.Header().Set("Content-Type", "application/json")


        //DISPLAY JSON
        fmt.Fprint(w,string(j))

  //END IF GET
  }

//END FUNC
}

/****************************FILES ADD*********************************/
func control_files(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET ADD**************************/
  if r.Method == "GET"{
  

  /*	 uploadURL, errURL := blobstore.UploadURL(c, "/"+AdminSlug+"/files/", nil)

  	 if errURL != nil {
  	 	fmt.Fprintln(w,"error with making upload URL on GET:" + errURL.Error())
  	 	return
  	 }*/

    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Files",
       //"uploadURL":uploadURL.String(), 
     } 
 

    //QUERY
    q := datastore.NewQuery("__BlobInfo__").Limit(100)


    //DB GET ALL
    var db []*BlobInfo 
    keys,err := q.GetAll(c,&db)

    //DB ERR
    if err != nil {
      fmt.Fprint(w,"error getting items: " + err.Error())
      return
    }

    //VAR
    var dbData []map[string]string


    //fmt.Fprintln(w,keys)

    
    //FOR DB ITEMS
    for i := range db {
      
      //KEYS ENCODE
      k := keys[i].Encode()

      dbData = append(dbData,
        map[string]string{
         "title": db[i].Filename,
         "key":k,

         /*Fields: map[string]string{
            "Name":"text",
            "Email":"text",
            "Phone":"text",
            "Message":"textarea", 
          },*/


       },
      )
    } 


  renderControl(w, r, "/control/files.html", data, dbData)
  


/********************************POST ADD*********************************/
} else {

      //fmt.Fprintln(w,"posted ajax:")
      
//PARSE
blobs,_,err := blobstore.ParseUpload(r)


if err != nil {
	
  fmt.Fprintln(w,"error with upload:"+err.Error())
	//serveError(c, w, err)
  return


} else {  

          var files map[string][]*blobstore.BlobInfo //map[string][]*BlobInfo
  
          theMsg := "nothing changed..."
          fileNames := ""
          objNames := ""
          bKey := ""
          MD5 := ""
          files = blobs //blobs["files[]"]
          imgKey := ""

      if len(files) == 0 {
          
          //SET CONTENT-TYPE
         // w.Header().Set("Content-Type", "application/json")

          //fmt.Fprintln(w,"message':'file is nil") 
          theMsg = "file is nil"
          //DISPLAY JSON
          
          //fmt.Fprint(w,theMsg)
          //http.Redirect(w, r, "/admin/files/", http.StatusFound)
          //return
      } else {

          //fmt.Fprintln(w,files["images[]"][0])




          for i := range files["images[]"]{
            fileNames = fileNames + "|" + files["images[]"][i].Filename
           // objNames = objNames + "|" + url.QueryEscape(files["images[]"][i].ObjectName)
            objNames = objNames + "|" + files["images[]"][i].ObjectName
            bKey = string(files["images[]"][i].BlobKey)
            MD5 = files["images[]"][i].MD5
            //imgURL,_ := image.ServingURL(c,files["images[]"][i].BlobKey,nil)
            //bKey = bKey + "|" +   imgURL.String()
              

            //fileNames = append(fileNames,string(files[i].Filename))
          }

                  //fmt.Fprintln(w, fileNames)
      //fmt.Fprintln(w,string(files[0].BlobKey))
      //fmt.Fprintln(w,string(files[0].Filename))

          theMsg = "upload success" //+ files["files"][0].Filename



      //END IF LEN IS 0
      }

     /* 
    //RENEW UPLOAD SESSION URL
    uploadURL, errURL := blobstore.UploadURL(c, "/"+AdminSlug+"/files/", nil)

     if errURL != nil {
      fmt.Fprintln(w,"error with creating upload url on success:" + errURL.Error())
      return
     }*/
    

       //stringKey := getImageKey(w,r,files["images[]"][0].Filename)

          //imgKey := string(appengine.BlobKey(bKey))

      //  imgKey := stringKey.Get(bKey)

     //TIMEOUT
      time.Sleep(time.Second * 2)


      //q := datastore.NewQuery("__BlobInfo__").Filter("Entity_Name =", bKey)
      q := datastore.NewQuery("__BlobInfo__").Filter("md5_hash =", MD5)
      //q := datastore.NewQuery("__BlobInfo__").Order("-creation").Limit(1)
      //fmt.Fprintln(w,string(stringKey))

      //DB GET ALL
      var db []*BlobInfo 
      keys,err2 := q.GetAll(c,&db)

      //DB ERR
      if err2 != nil {
        fmt.Fprint(w,"error getting items entity key: " + err2.Error())
        return
      }


    //FOR DB ITEMS
    for i := range db {
      
      //KEYS ENCODE
      imgKey = imgKey + keys[i].Encode()

    } 


        //SETUP JSON
        m := map[string]string{
          "message": theMsg,
          "cloud":objNames,
          "files": fileNames,
          "adminSlug":AdminSlug,
          //"uploadURL": uploadURL.String(),
          "keys":bKey,
          "MD5": MD5,
          "imgKey": imgKey,
         // "files":string(files[0].Filename),
          //"title": string(file[0].Filename), //newItem.Title,
        }

        //MARSHAL JSON
        j,errJSON2 := json.Marshal(m)
        if errJSON2 != nil {
          fmt.Fprintln(w,"error with JSON")
        }

        //SET CONTENT-TYPE
        w.Header().Set("Content-Type", "application/json")


        //DISPLAY JSON
        fmt.Fprint(w,string(j))

//END ERR NIL
}





//fmt.Fprintln(w,"whaddup")
//http.Redirect(w, r, "/images/"+string(file[0].BlobKey), http.StatusFound)

//http.Redirect(w, r, "/admin/files/", http.StatusFound)

//http.Redirect(w, r, "/images/"+string(file[0].Filename), http.StatusFound)
//blobstore.Send(w, appengine.BlobKey(string(file[0].BlobKey)))
//b := blobstore.NewReader(c,file[0].BlobKey)
//record, _ := reader.Reader() 
//appengine.BlobKey(string(file[0].BlobKey))



//END POST
}

//END FUNC
}


/****************************FILES DELETE*********************************/
func control_files_delete(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)



  /********************************POST DELETE*********************************/
  if r.Method == "POST"{
  
    //CONTEXT
    c := appengine.NewContext(r)

    formVal := func(val string)string{
      return r.FormValue(val)
    }

    //GET KEY
    itemID := strings.Split(r.RequestURI,"/") 

    
    //KEY DECODE
    key,err := datastore.DecodeKey(itemID[4])
    
    //KEY ERR
    if err != nil {
      fmt.Fprintln(w, "error decoding")
      return
    }

    //BLOB KEY
    theBlobKey := appengine.BlobKey(key.StringID())
   
   //DELETE SERVING URL
   errDelete0 := image.DeleteServingURL(c, theBlobKey)

    if errDelete0 != nil {
      fmt.Fprintln(w,"error deleting")
      return
    }


    //DB DELETE
    //errDelete := datastore.Delete(c, key)
    //errDelete2 := blobstore.Delete(c, appengine.BlobKey(itemID[4]))

   //DELETE FROM BLOBSTORE
    errDelete2 := blobstore.Delete(c, theBlobKey)
    
    //DB ERR
    //if errDelete != nil && 
    if errDelete2 != nil {
      fmt.Fprintln(w,"error deleting")// + key.StringID() + itemID[4])
      return
    
    //NO ERR
    } else {


      //FLUSH CACHE
      cacheFlush(string("image_" + formVal("imageName")),r)
      //cacheFlush(string("StyleSheet_" + formVal("styleName")),r) 

      fmt.Fprintln(w,"delete successful ") //+ key.StringID()  + itemID[4] + " images")
      return
    }


  //END POST
  } else {
    fmt.Fprintln(w, "say whaaa can't delete file...")
  }


//END FUNC
}
