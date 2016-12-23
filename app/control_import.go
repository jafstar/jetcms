package app
/*

import (
  "net/http"
  "fmt"
  "appengine"
  "appengine/datastore"
  "appengine/blobstore"
  "encoding/csv"

  //"strings"
  //"strconv"
  //"encoding/json"

)


//var queryType = "data"


/****************************GET LIST PAGES/POST ADD PAGES********************************
func control_import(w http.ResponseWriter, r *http.Request) {
  
  //ADMIN
  checkAdmin(w,r)

  //CONTEXT
  c := appengine.NewContext(r)

  /*******************************GET*************************
  if r.Method == "GET"{
  
  	 uploadURL, err := blobstore.UploadURL(c, "/"+AdminSlug+"/import/", nil)

  	 if err != nil {
  	 	fmt.Fprintln(w,"error with upload:" + err.Error())
  	 	return
  	 }

    //DATA
     data := map[string]string{
       "get":"true",
       "title": "Import",
       "uploadURL":uploadURL.String(), 
     } 


    //QUERY
    q := datastore.NewQuery("__BlobUploadSession__").Limit(100)


    //DB GET ALL
    var db []*BlobUploadSession 
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
        map[string]string {
         "title": db[i].Filename,
         "key":k,

       },
      )
    } 


  renderControl(w, r, "/control/import.html", data, dbData)
  


/********************************POST********************************
} else {


blobs, _, err := blobstore.ParseUpload(r)

if err != nil {
	fmt.Fprintln(w,"error with upload:"+err.Error())
	//serveError(c, w, err)
    return
}

file := blobs["file"]

if len(file) == 0 {
    fmt.Fprintln(w,"file is nil")  
    //http.Redirect(w, r, "/", http.StatusFound)
    return
}

//fmt.Fprintln(w,"whaddup")
//http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)

//blobstore.Send(w, appengine.BlobKey(string(file[0].BlobKey)))

b := blobstore.NewReader(c,file[0].BlobKey)

//record, _ := reader.Reader() 
//appengine.BlobKey(string(file[0].BlobKey))

reader := csv.NewReader(b)  

		
		//FIRST RECORD
        record, err := reader.Read()

  		if err != nil {
            fmt.Println("Read Err:", err.Error())
            return
        }
 
 

        //ALL OTHER RECORDS
        records, err := reader.ReadAll()

        if err != nil {
        	fmt.Fprintln(w,"Read All Err: "+err.Error())
        	return
        }


        //ITERATE
        for k,_ := range records {
        	for j,_ := range record {
        	  fmt.Fprintln(w,record[j] + " : "+records[k][j])
        	}
        	fmt.Fprintln(w,"")
        }
    
    	//EOF
    	fmt.Fprintln(w,"nothing left to see...")


//END POST
}

//END FUNC
}
*/
