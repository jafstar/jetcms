package app


//ADMIN OPTIONS
var AdminSlug string = "admin"

//DATE FORMATS
var ItemDateDisplay = "01-02-2006 3:04pm" //control_data_item.go
var FormatDateDisplay = "Jan 2, 2006"

//PAGE CACHE
//var CACHE_TIME string = "max-age:3600, public"
var CACHE_CONTROL string = "max-age:0, public"
var CACHE_DAYS int = 0
var CACHE_MONTHS int = 0

//404
var TITLE_404 string = "404 Page Not Found"
var DESC_404 string = "Page Not Found"
var KEYW_404 string = "404"

//RANDOM
var HelloWorld string = "Hello World!"
var HolaWorld string = "Hola World!"

//DEFAULT COMING SOON
var SOON_TITLE string = "Coming Soon"
var SOON_DESC string = "This website is under maintenance, please check back soon."
var SOON_KEYW string = "coming soon"

//DEFAULT HOMEPAGE
var HOME_TITLE string = ""
var HOME_DESC string = ""
var HOME_KEYW string = ""




//DATA TEMPLATES
var TPL_LIST = "<h1>{{.Data.title}}</h1> <b>Date: {{.Data.date}}</b><p>{{.Data.story}}</p>"
var TPL_ITEM = " <h1>News</h1> <br> {{range $k,$v := .List}} <h2>{{$v.content.title}}</h2><p>{{$v.content.date}}</p>  <p>{{formatDate $v.content.date,\"Jan 2, 2006\"}}</p>  {{end}}"

//DEFAULTS
var DEFAULT_CSS string = `body {
	background: #000;
	font-family:arial;
	color:#ccc;
	text-align: center;
}
a {
  font-family: 'numans';
  color: #fff;
  text-decoration: none;
  font-size: 60px;
}

#logo {
  width: 300px;
  height: 100px;
  position: absolute;
  left: 50%;
  top: 50%;
  margin-top: -100px;
margin-left: -150px;
}

#logo p {
 color: #fff;
  font-family: arial;
}`

var DEFAULT_CSS_HEADER  string = `<!DOCTYPE html>
<head>
  <!--TITLE-->
  <title>{{.Data.title}}</title>

  <!--META-->
  <meta charset="utf-8" />
  <meta name='description' content='{{.Data.description}}' />
  <meta name='keywords' content='{{.Data.keywords}}' />
  <meta name="viewport" content="width=device-width,initial-scale=1, maximum-scale=1, user-scalable=no" />
  
  <!--LINK-->
  <link href='https://fonts.googleapis.com/css?family=Archivo+Black|Numans' rel='stylesheet' type='text/css'>
  <link rel='stylesheet' type='text/css' href='/styles/default.css' />
  <link rel='icon' href='/favicon.ico' type='image/x-icon' />

</head>
<body class="{{.Data.class}}">`

var DEFAULT_CSS_FOOTER string = `
</body>
</html>
`
var DEFAULT_TYPE_LIST string = `<h1>News</h1>
<br />
{{range $k,$v := .List}}
<h2>{{FormatDate $v.content.date}} - <a href="/news/{{$v.info.url}}">{{$v.content.title}}</a></h2>
{{end}}`

var DEFAULT_TYPE_ITEM string = `<h1>{{.Data.title}}</h1>
<b>{{FormatDate .Data.date}}</b>
<div>{{FormatHTML .Data.story}}</div>`


var DEFAULT_PAGE_CONTENT = `<div id='logo'>
<h1>
  <a href='/'>
    <span id="logo_1">JET</span>
    <img src="/img/core/logo.png" />
  </a>
</h1>
  
<p>Welcome to the default homepage.</p>

</div>`