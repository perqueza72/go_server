package constants

const (
	HOME_PAGE = iota
	HELLO_PAGE
	FORM_PAGE
)

const (
	GET  = "GET"
	POST = "POST"
	PUT  = "PUT"
	DEL  = "DEL"
)

var PAGES = map[int]string{
	HOME_PAGE:  "/",
	HELLO_PAGE: "/hello",
	FORM_PAGE:  "/form",
}

var URLS = map[string][]string{
	"/":      {GET},
	"/hello": {GET},
	"/form":  {GET, POST},
}

const PORT string = "8080"
