package app

import (
  "net/http"
)


//INIT
func init() {

  //PARSER
  http.HandleFunc("/", parser)

  //http.HandleFunc("/", staticParser)

  //DYNAMIC STYLES
  http.HandleFunc("/styles/",styles)


  //DYNAMIC IMAGES
  http.HandleFunc("/images/",imagesNormal)
 http.HandleFunc("/images/cloud/",imagesCloud)
  
  http.HandleFunc("/images/fast/",imagesFast)
  http.HandleFunc("/images/desktop/",imagesResize)
  http.HandleFunc("/images/tablet/",imagesResize)
  http.HandleFunc("/images/mobile/",imagesResize)

  http.HandleFunc("/images/large/",imagesResize)
  http.HandleFunc("/images/medium/",imagesResize)
  http.HandleFunc("/images/small/",imagesResize)


  http.HandleFunc("/images/icon/",imagesResize)
  http.HandleFunc("/images/mini/",imagesResize)
  http.HandleFunc("/images/micro/",imagesResize)


  //CONTROL ROOM
  http.HandleFunc("/"+AdminSlug+"/", room)
  http.HandleFunc("/"+AdminSlug+"/stats", room)
  http.HandleFunc("/"+AdminSlug+"/room", room)


  //CONTROL SETTINGS
  http.HandleFunc("/"+AdminSlug+"/settings/", control_settings)


  //CONTROL FLUSH CACHE
  http.HandleFunc("/"+AdminSlug+"/cache/flush", cacheFlushAll)


  //CONTROL STATIC GENERATOR
  http.HandleFunc("/"+AdminSlug+"/static/generator/", staticGenerator)


  //CONTROL DESIGN
  //http.HandleFunc("/"+AdminSlug+"/design/", control_design)
  //http.HandleFunc("/"+AdminSlug+"/design/edit/", control_design_edit)
  //http.HandleFunc("/"+AdminSlug+"/design/delete/", control_design_delete)

  //CORE
  http.HandleFunc("/"+AdminSlug+"/core/", control_core)

  //STYLE
  http.HandleFunc("/"+AdminSlug+"/style/", control_style)
  http.HandleFunc("/"+AdminSlug+"/style/edit/", control_style_edit)
  http.HandleFunc("/"+AdminSlug+"/style/delete/", control_style_delete)


  //CONTROL PAGES
  http.HandleFunc("/"+AdminSlug+"/pages/edit/", control_pages_edit)
  http.HandleFunc("/"+AdminSlug+"/pages/list/", control_pages_list)
  http.HandleFunc("/"+AdminSlug+"/pages/delete/", control_pages_delete)

  //CONTROL DATA TYPES
  http.HandleFunc("/"+AdminSlug+"/data/type/", control_data_type)
  http.HandleFunc("/"+AdminSlug+"/data/edit/", control_data_edit)
  http.HandleFunc("/"+AdminSlug+"/data/delete/", control_data_delete)
  http.HandleFunc("/"+AdminSlug+"/data/type/details/", control_data_type_details)

  //CONTROL DATA ITEMS
  http.HandleFunc("/"+AdminSlug+"/data/item/", control_data_item)
  http.HandleFunc("/"+AdminSlug+"/data/item/edit/", control_data_item_edit)
  http.HandleFunc("/"+AdminSlug+"/data/item/delete/", control_data_item_delete)


  //CONTROL FILES
  http.HandleFunc("/"+AdminSlug+"/files/", control_files)
  http.HandleFunc("/"+AdminSlug+"/files/delete/", control_files_delete)
  http.HandleFunc("/"+AdminSlug+"/files/upload", control_files_upload_url)

  //CONTROL FORMS
  http.HandleFunc("/"+AdminSlug+"/forms/", control_forms)


//END INIT FUNC
}



