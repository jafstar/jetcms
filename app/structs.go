package app

import (
	//"reflect"

  "appengine"
 // "appengine/datastore"
  "html/template"
  "time"
)

/*
type Map map[string]interface{}

func (m Map) Load(c <-chan datastore.Property) error {
	for p := range c {
		if p.Multiple {
			value := reflect.ValueOf(m[p.Name])
			if value.Kind() != reflect.Slice {
				m[p.Name] = []interface{}{p.Value}
			} else {
				m[p.Name] = append(m[p.Name].([]interface{}), p.Value)
			}
		} else {
			m[p.Name] = p.Value	
		}
	}
	return nil
}

func (m Map) Save(c chan<- datastore.Property) error {
	defer close(c)
	for k, v := range m {
		c <- datastore.Property {
			Name: k,
			Value: v,
		}
	}
	return nil
}*/

//IMAGE STRUCTS
type ServingURLOptions struct {
    Secure bool // whether the URL should use HTTPS

    // Size must be between zero and 1600.
    // If Size is non-zero, a resized version of the image is served,
    // and Size is the served image's longest dimension. The aspect ratio is preserved.
    // If Crop is true the image is cropped from the center instead of being resized.
    Size int
    Crop bool
}


//BLOB STRUCTS
type BlobUploadSession struct {
    BlobKey      appengine.BlobKey
    ContentType  string    `datastore:"content_type"`
    CreationTime float64     `datastore:"creation"`
    Filename     string    `datastore:"filename"`
    Hash		 string	   `datastore:"md5_hash"`
    Size         int64     `datastore:"size"`
    User		 string	   `datastore:"user"`
    SuccessPath  string    `datastore:"success_path"`
    State		 string	   `datastore:"state"`
    MaxBytes	 string	   `datastore:"max_bytes_total"`
    MaxBytesPB   string    `datastore:"max_bytes_per_blob"`

}

type BlobInfo struct {
	BlobKey      appengine.BlobKey
    ContentType  string    `datastore:"content_type"`
    CreationTime time.Time `datastore:"creation"`
    Filename     string    `datastore:"filename"`
    Size         int64     `datastore:"size"`
    MD5          string    `datastore:"md5_hash"`
    UploadID	string		`datastore:"upload_id"`

    // ObjectName is the Google Cloud Storage name for this blob.
    ObjectName string `datastore:"gs_object_name"`
}


//BASIC STRUCTS
type Page struct {
  Slug string			`json:"slug"`
  Title string			`json:"title"`
  Description string	`json:"description"`
  Keywords string		`json:"keywords"`
  Content string 		`datastore:",noindex" json:"content"`
  JS string				`json:"js"`
  CSS string			`json:"css"`
  Template string		`json:"template"`
  Single string			`json:"single"`
  //List []map[string]string
}

type Content struct {
	List []map[string]string
	Data map[string]string
	Globals map[string]string
}


type Core struct {
	Core_Header string 		`datastore:",noindex" json:"core_header"`
  	Core_Footer string 		`datastore:",noindex" json:"core_footer"`
  	Core_Version string		`json:"core_version"`
  	Last_Update time.Time
}

type Style struct {
	Style_Name string 		`json:"style_name"`
  	Style_Sheet string 		`datastore:",noindex" json:"style_sheet"`
  	Last_Update time.Time
}


type Settings struct {
	Site_Title string	`json:"site_title"`
	Site_Closed string	`json:"site_closed"`
	Site_Analytics string `json:"site_analytics"`

	Home_Title string	`json:"home_title"`
	Home_Desc string	`json:"home_desc"`
	Home_KeyW string	`json:"home_keyw"`

	Cache_Control string `json:"cache_control"`
	Cache_Days string 		 `json:"cache_days"`
	Cache_Months string 	 `json:"cache_months"`
	
	Soon_Title string	`json:"soon_title"`
	Soon_Desc string	`json:"soon_desc"`
	Soon_KeyW string	`json:"soon_keyw"`
	Soon_URL string		`json:"soon_url"`

	Admin_Email string	`json:"admin_email"`
	From_Email string	`json:"from_email"`
	To_Email string 	`json:"to_email"`
	Last_Update time.Time
}


//VIEW STRUCTS
type View struct {
	HTML map[string]template.HTML
    Data map[string]string
    List []map[string]map[string]string
    //Format func(string)string
   // Menus Menu
    //Site map[string]string
}

type xmlView struct {
	HTML map[string]template.HTML
    Data map[string]string
    Pages []string
}

type cssView struct {
	CSS map[string]template.HTML
}

//FORMS
type Forms struct {
	//Key string			`json:"key"`
	Title string			`json:"title"`
	TemplateList  string 
	TemplateItem  string 
	Fields []FieldType      `json:"fields"`

}

type FormFieldType struct {
	Name string
	//Label string
	UI string
	//Value string
	//Errors string
}


//DATA STRUCTS
type DataType struct {
	//Key string			`json:"key"`
	Title string			`json:"title"`
	URL string				`json:"typeurl"`
	Description string
	Keywords string
	TemplateList  string 
	TemplateItem  string 
	Fields []FieldType      `json:"fields"`

}

type FieldType struct {
	Name string
	//Label string
	UI string
	//Value string
	//Errors string
}

type DataItem struct {
	Content []byte //[]FieldItem //map[string]string
	Type string
	Tags []string
	Title string
	URL  string
	TypeURL string
	Description string
	Keywords string
	DateOrder	time.Time //`datastore:"date_order"`
	Published bool
	//CreationTime time.Time //`datastore:"creation"`

}

type DataTemplate struct {
	TPL map[string]template.HTML
}


/*
type FieldItem struct {
	//Value map[string]string
	Value string
	//Name string
	//Value string
}

type FiedName struct {
	Name string
}
*/

/*

type Menu struct {
  Name string
  Items []MenuItem
}

type MenuItem struct {
	Value string
	URL string
  	Sub []SubItem1
}

type SubItem1 struct {
	Value string
	URL string
	Sub []SubItem2
}

type SubItem2 struct {
	Value string
	URL string
}*/

/*
type Menu struct {
  Name string
  Items []MenuItem
}

type MenuItem struct {
	Value string
	URL string
  Sub []SubItem
}

type SubItem struct {
	Value string
	URL string

}
*/

/*
var menus = Menu {
  	Name: "Main",
  	Items:  []MenuItem {
  	  
	  	0:MenuItem{
	  		Value: "Home",
	  		URL: "/",
	  	},
	  	1:MenuItem{
	  		Value:"About",
	  		URL: "/about",
	  		Sub: []SubItem {
	  			0:{
	  				Value:"History",
	  				URL:"/about/history",
	  			},
	  			1:{
	  				Value:"Tokyo",
	  				URL:"/tokyo",
	  			},
	  		},
	  	},
	  	2:MenuItem{
	  		Value: "Contact",
	  		URL:"/contact",
	  	},
  	},

}
*/


