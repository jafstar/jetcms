package app

import (
 "appengine/memcache"
   "net/http"
  "fmt"
 "appengine"
 "strings"
  "encoding/json"

)
func cacheFlushAll(w http.ResponseWriter, r *http.Request){

	  //ADMIN
  checkAdmin(w,r)

		 c := appengine.NewContext(r)

	errFlush := memcache.Flush(c)

	theMsg := ""

	if errFlush != nil {

		theMsg = "Error flushing cache"

	} else {

 		theMsg = "Cache Flush Successful"


    //END ELSE
	}

	   //PREP JSON
    data := map[string]string{
    	"message": theMsg,
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


//END FUNC
}

func cacheAdd(k string, r *http.Request, pages []string){

	 c := appengine.NewContext(r)

	// Create an Item
	item := &memcache.Item {
	    Key:   k,
	    //Value: []byte("Oh, give me a home"),
	    Value: []byte(strings.Join(pages,",")),
	}

	// Add the item to the memcache, if the key does not already exist
	if err := memcache.Add(c, item); err == memcache.ErrNotStored {
	    
	    c.Infof("item with key %q already exists", item.Key)

	} else if err != nil {
	    c.Errorf("error adding item: %v", err)
	}

}

func cacheAddImg(k string, r *http.Request, val string){

	 c := appengine.NewContext(r)

	// Create an Item
	item := &memcache.Item {
	    Key:   k,
	    //Value: []byte("Oh, give me a home"),
	    Value: []byte(val),
	}

	// Add the item to the memcache, if the key does not already exist
	if err := memcache.Add(c, item); err == memcache.ErrNotStored {
	    
	    c.Infof("item with key %q already exists", item.Key)

	} else if err != nil {
	    c.Errorf("error adding item: %v", err)
	}

}

func cacheAddMap(r *http.Request, pages map[string]string){

	 c := appengine.NewContext(r)

	 //var itemArr []*memcache.Item

	//FOR LOOP
    for k,v := range pages {

    	if k != "key" {
    		
    	

		// Create an Item
		item := &memcache.Item {
		    Key:   k,
		    //Value: []byte("Oh, give me a home"),
		   
		    //Value: []byte(pages["CoreHeader"]),
		    Value: []byte(v),
		}

		//itemArr[0] = item

		// Add the item to the memcache, if the key does not already exist
		if err := memcache.Add(c, item); err == memcache.ErrNotStored {
		    
		    c.Infof("item with key %q already exists", item.Key)

		} else if err != nil {
		    c.Errorf("cache error adding item: %v", err)
		}

		}

	//END FOR
	}

	

}

func cacheFlush(k string, r *http.Request){

	 c := appengine.NewContext(r)

	if err := memcache.Delete(c, k); err != nil {
	    c.Errorf("error deleting single item: %v", err)
	}

}

func cacheFlushMulti(k []string, r *http.Request){

	 c := appengine.NewContext(r)

	if err := memcache.DeleteMulti(c, k); err != nil {
	    c.Errorf("error deleting item: %v", err)
	}

}

func cacheGet(k string, r *http.Request)(*memcache.Item, error){

	 c := appengine.NewContext(r)

	// Get the item from the memcache
	item, err := memcache.Get(c, k)

	if err == memcache.ErrCacheMiss {
	    c.Infof("item not in the cache")
	} else if err != nil {
	    c.Errorf("cache error getting item: %v", err)

	} else {
	   // c.Infof("the getCache value is %q", item.Value[:])
	}

	return item,err

}

func cacheGetMulti(k []string, r *http.Request)(map[string]*memcache.Item, error){

	 c := appengine.NewContext(r)

	// Get the item from the memcache
	items, err := memcache.GetMulti(c, k)

	if err == memcache.ErrCacheMiss {
	    c.Infof("item not in the cache")
	} else if err != nil {
	    c.Errorf("cache error getting multi items: %v", err)

	} else {
	    //c.Infof("the value is %q", items)
	}

	return items,err

}



