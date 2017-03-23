package app

import (
  "fmt"
  "strings"
  "net/url"
  "net/http"
  //"html/template"
  "appengine"
  //"appengine/user"
  "appengine/blobstore"
  "appengine/image"
  "time"

  "log"
  "io"
  "io/ioutil"
  //"bytes"
  //"image"
    "appengine/urlfetch"

   //     "cloud.google.com/go/storage"
//        "golang.org/x/net/context"

          //  "google.golang.org/appengine/file"
      //  "google.golang.org/appengine/log"

)

//FUNC IMAGES WITH BLOBKEY
func imagesFast(w http.ResponseWriter, r *http.Request) {

  //PARSE URL
  realURL,_ := url.Parse(r.RequestURI)
  url := realURL.String()
  urlArr := strings.Split(url,"/")

  //RENDER IMAGE
  blobstore.Send(w, appengine.BlobKey(urlArr[3]))


//END FUNC
}

//FUNC IMAGES WITH FILENAME
func imagesNormal(w http.ResponseWriter, r *http.Request) {

  //PARSE URL
  realURL,_ := url.Parse(r.RequestURI)
  url := realURL.String()
  urlArr := strings.Split(url,"/")


  //IF THERE IS A FILENAME
  if strings.ContainsAny(urlArr[2],"."){



	  //FILENAME
	  fileName := urlArr[2];


	  //GET IMAGE KEY
	  stringKey := getImageKey(w,r,fileName)

	  //CONVERT STRING TO BLOBKEY
	  imgKey := appengine.BlobKey(stringKey)


	  //DEBUG
	  //l := len(url)
	 	//fmt.Fprintln(w, "Clean URL: " + url)
	  //fmt.Fprintln(w, "Real  URL: " + realURL.Path)
	  //fmt.Fprintln(w, "FileName: " + fileName)

	  w.Header().Set("Cache-Control", "max-age:3600, public")
	  w.Header().Set("Expires", time.Now().AddDate(0, 0, 1).Format(http.TimeFormat))

	  //RENDER IMAGE
	  blobstore.Send(w,imgKey)

  //DOES NOT CONTAIN FILENAME
	} else {

		settings :=  getSettings(w,r)
  	display404(w,settings)
  	return

	//END ELSE	
	}


//END FUNC
}


//FUNC IMAGES STORED IN CLOUD BUCKET
func imagesCloud(w http.ResponseWriter, r *http.Request) {

c := appengine.NewContext(r)
/*
     //BUCKET NAME
        bucketName, err := file.DefaultBucketName(ctx); 
        
        if err != nil {
            //log.Errorf(ctx, "failed to get default GCS bucket name: %v", err)
            fmt.Fprintln(w,"error with making upload URL on GET:" + err.Error())
            return
        }*/

  //PARSE URL
  realURL,_ := url.Parse(r.RequestURI)
  url := realURL.String()
  urlArr := strings.Split(url,"/")


  //IF THERE IS A FILENAME
  if strings.ContainsAny(urlArr[3],"."){



	  //FILENAME
	  fileName := urlArr[3];


	  //GET IMAGE KEY
	  stringKey := getImageKey(w,r,fileName)

	  //CONVERT STRING TO BLOBKEY
	  imgKey,_ := blobstore.BlobKeyForFile(c,"/gs/test_bucket/"+fileName)


	  //DEBUG
	  //l := len(url)
	 fmt.Fprintln(w, "Cloud URL: " + stringKey)
	 	 fmt.Fprintln(w, "Blok Key: " + imgKey)

	  //fmt.Fprintln(w, "Real  URL: " + realURL.Path)
	  //fmt.Fprintln(w, "FileName: " + fileName)

	  w.Header().Set("Cache-Control", "max-age:3600, public")
	  w.Header().Set("Expires", time.Now().AddDate(0, 0, 1).Format(http.TimeFormat))

	  //RENDER IMAGE
	 // blobstore.Send(w,imgKey)

  //DOES NOT CONTAIN FILENAME
	} else {

		settings :=  getSettings(w,r)
  	display404(w,settings)
  	return

	//END ELSE	
	}


//END FUNC
}


//FUNC IMAGES WITH RESIZE
func imagesResize(w http.ResponseWriter, r *http.Request) {

c := appengine.NewContext(r)

  //PARSE URL
  realURL,_ := url.Parse(r.RequestURI)
  url := realURL.String()
  urlArr := strings.Split(url,"/")

  //VARS
  fileName := ""
  var urlOptions *image.ServingURLOptions

  //IF RESIZING
  if(len(urlArr) > 3){

  	//GRAB FILENAME
  	fileName = urlArr[3];

  	//DESKTOPS
  	if urlArr[2] == "desktop"{
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 1220,
		  	Crop: false,
		  	}

		//TABLETS
  	} else if urlArr[2] == "tablet"{
			
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 720,
		  	Crop: false,
	  		}

	  //MOBILES
  	} else if urlArr[2] == "mobile"{
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 320,
		  	Crop: false,
	  		}
	//LARGE
  	} else if urlArr[2] == "large"{
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 900,
		  	Crop: false,
	  		}


	//MEDIUM
  	} else if urlArr[2] == "medium"{
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 550,
		  	Crop: false,
	  		}

	//SMALL
  	} else if urlArr[2] == "small"{
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 240,
		  	Crop: false,
	  		}  		

	//ICON
  	} else if urlArr[2] == "icon"{
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 125,
		  	Crop: false,
	  		}
	//MINI
  	} else if urlArr[2] == "mini"{
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 64,
		  	Crop: false,
	  		}

	//MICRO
  	} else if urlArr[2] == "micro"{
			urlOptions = &image.ServingURLOptions{
		  	Secure: false,
		  	Size: 32,
		  	Crop: false,
	  		}

	  //NONE
  	} else {
  		urlOptions = nil
  		//BREAKS
  	}

  //NO RESIZE
  } else {
  	fileName = urlArr[2];
  }


  //GET IMAGE KEY
  stringKey := getImageKey(w,r,fileName)

  //CONVERT STRING TO BLOBKEY
  imgKey := appengine.BlobKey(stringKey)

  //GET SERVING URL FOR RESIZE/CROP
  srvURL,_ := image.ServingURL(c,imgKey,urlOptions)


  	//DEBUG
  	//l := len(url)
 	//fmt.Fprintln(w, "Clean URL: " + url)
  	//fmt.Fprintln(w, "Real  URL: " + realURL.Path)
  	//fmt.Fprintln(w, "FileName: " + fileName)
  	//fmt.Fprintln(w,srvURL)

  cacheName := "image_" + urlArr[2] + "_" + fileName
  cacheType := "mime_"  + urlArr[2] + "_" + fileName 


  //MAP LIST
  list := []string {
    cacheName,
    cacheType,
  }

  //GET CACHE
  cache,err := cacheGetMulti(list,r)

  c.Errorf("there is a query error getting items: %v", err)
  c.Errorf("items: %v", len(cache))

  //fmt.Fprint(w,string(cache[:]))


  //IF CACHE EMPTY
  if len(cache) != len(list)  {
	  

	  //SETUP URL FETCH
	  client := urlfetch.Client(c)
	    
	    //GET URL
	    resp, err := client.Get(srvURL.String())

	    //ERR
	    if err != nil { 
	    	//DISPLAY WITH REDIRECT
	  		http.Redirect(w,r,srvURL.String(),http.StatusFound)
	  		log.Fatal(err.Error()) //http.Error(w, err.Error(), http.StatusInternalServerError)
	        return
	    }

	  	//VAR
	  	contentType := resp.Header["Content-Type"][0]

		//READCLOSER -> BYTES
		img, err := ioutil.ReadAll(resp.Body)

		//CLOSE
		resp.Body.Close()
		
		//ERR
		if err != nil {
			log.Fatal(err)
		}

		//DEBUG
	  	//fmt.Fprintln(w, contentType)
	 	//fmt.Fprintln(w,srvURL.String())
	  	//fmt.Fprint(w,resp.Body)
	  	//fmt.Fprintln(w,img)

		//CACHE
		cacheAddImg(cacheName, r, string(img))
		cacheAddImg(cacheType, r, contentType)


		//SET HEADERS
	  	w.Header().Set("Content-Type", contentType)
		w.Header().Set("Cache-Control", "max-age:3600, public")
  		w.Header().Set("Expires", time.Now().AddDate(0, 0, 1).Format(http.TimeFormat))

	  	//DISPLAY FROM BYTES
		io.WriteString(w,string(img))

	//ELSE USE CACHE
	} else {

		//VARS		 
		 cachedImg  := string(cache[cacheName].Value[:])
		 cachedType := string(cache[cacheType].Value[:])

		 //DEBUG
		 //fmt.Fprintln(w,"using cache")
		 //fmt.Fprintln(w, string(cache[cacheType].Value[:]))
		 //fmt.Fprint(w,[]byte(cachedImg))

		//SET HEADERS
	  	w.Header().Set("Content-Type", cachedType)
	  	w.Header().Set("Cache-Control", "max-age:3600, public")
  		w.Header().Set("Expires", time.Now().AddDate(0, 0, 1).Format(http.TimeFormat))

	  	//DISPLAY FROM BYTES
		io.WriteString(w,cachedImg)


	//END IF CACHE 
	}



	//BYTE TO IMAGE NOTES
  	//https://www.socketloop.com/tutorials/golang-convert-byte-to-image

  
  	/***
  	//NOTES ON OPTIONS
  	//http://stackoverflow.com/questions/25148567/list-of-all-the-app-engine-images-service-get-serving-url-uri-options
	//http://[image-url]=s200-r90-cc-c0xFF00FF00-fSoften=1,20,0:

  	SIZE / CROP

	s640 — generates image 640 pixels on largest dimension
	s0 — original size image
	w100 — generates image 100 pixels wide
	h100 — generates image 100 pixels tall
	p — smart square crop, attempts cropping to faces
	pp — alternate smart square crop, does not cut off faces (?)
	cc — generates a circularly cropped image
	ci — square crop to smallest of: width, height, or specified =s parameter

	ROTATION

	fv — flip vertically
	fh — flip horizontally
	r90 — rotate 90 degrees (or 180 or 270)

	IMAGE FORMAT

	rj — forces the resulting image to be JPG
	rp — forces the resulting image to be PNG
	rw — forces the resulting image to be WebP
	rg — forces the resulting image to be GIF
	Forcing PNG, WebP and GIF outputs can work in combination with circular crops for a transparant background. Forcing JPG can be combined with border color to fill in backgrounds in transparent images.

	ANIMATED GIFs

	rh — generates an MP4 from the input image
	k — kill animation (generates static image)

	MISC.

	b10 — add a 10px border to image
	c0xAARRGGBB — set border color, eg. =c0xffff0000 for red
	d — adds header to cause browser download
	e7 — set cache-control max-age header on response to 7 days
	l100 — sets JPEG quality to 100% (1-100)

	Filters

	fSoften=1,100,0: - where 100 can go from 0 to 100 to blur the image
	fVignette=1,100,1.4,0,000000 where 100 controls the size of the gradient and 000000 is RRGGBB of the color of the border shadow
	
	Caveats
	
	Some options (like =l for JPEG quality) do not seem to generate new images. If you change another option (size, etc.) and change the l value, the quality change should be visible. Some options also don't work well together. This is all undocumented by Google, probably with good reason.
	Moreover, it's probably not a good idea to depend on any of these options existing forever. Google could remove most of them without notice at any time.

	****/

  	//NOTES ON PYTHON PIL & JPG SUPPORT
  	//http://stackoverflow.com/questions/14654781/installing-pil-on-osx-mountain-lion-for-google-app-engine
  	//http://www.janeriksolem.net/2012/09/installing-pil-on-os-x-mountain-lion.html


//END FUNC
}