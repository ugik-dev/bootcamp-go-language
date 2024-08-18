package main

import (
	"embed"
	"fmt"
	"goweb/unittest"
	"net/http"
	"strconv"
)

func MySetCookie(res http.ResponseWriter, req *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Username"
	cookie.Value = "Username_" + req.URL.Query().Get("first_name")
	cookie.Path = "/"

	cookieToken := new(http.Cookie)
	cookieToken.Name = "X-Token"
	cookieToken.Value = "Token_" + req.URL.Query().Get("first_name")
	cookieToken.Path = "/"

	http.SetCookie(res, cookie)
	http.SetCookie(res, cookieToken)
	fmt.Fprint(res, "Success Create Cookies")
}

func GetCookie(res http.ResponseWriter, req *http.Request) {
	cookieUser, err := req.Cookie("X-Username")

	if err != nil {
		fmt.Fprint(res, err.Error())
	} else {
		valUsername := cookieUser.Value
		fmt.Fprintf(res, "Helo %s", valUsername)

	}
}

//go:embed resources/*
var resources embed.FS

//go:embed resources/ok.html
var resourcesOK string

//go:embed resources/404-page.html
var resources404 string

func main() {
	mux := http.NewServeMux()
	// tanpa embed dan resource ini akan ikut terbuild
	// directory := http.Dir("./resources")
	// fileServer := http.FileServer(directory)

	// mengunakan embed FS
	// http://localhost:5000/static/resources/index.js
	// fileServer := http.FileServer(http.FS(resources))

	// mengunakan embed FS
	// http://localhost:5000/static/resources/index.js

	// directory, _ := fs.Sub(resources, "resources")
	// fileServer := http.FileServer(http.FS(directory))
	fileServer := http.FileServer(http.Dir("./resources"))

	// template.go
	// mux.HandleFunc("/template/text/", SimpleHTML)
	// mux.HandleFunc("/template/file/", SimpleHTMLFile)
	// mux.HandleFunc("/template/embed/", TemplateEmbed)
	// mux.HandleFunc("/template/data/map", TemplateDataMap)
	// mux.HandleFunc("/template/data/struct", TemplateDataStruct)
	// mux.HandleFunc("/template/action/", TemplateAction)
	// mux.HandleFunc("/template/action/range", TemplateDataRange)
	// mux.HandleFunc("/template/function", TemplateFunction)

	// templating_caching.go
	mux.HandleFunc("/template/caching", TemplateCaching)
	mux.HandleFunc("/template/xxs", TemplateXxs)
	mux.HandleFunc("/template/redirect-from", TemplatrRedirect)
	mux.HandleFunc("/template/redirect-to", RedirectTo)
	mux.HandleFunc("/template/redirect-ugik", RedirectUgikSite)

	mux.HandleFunc("/form-upload", FormUpload)
	mux.HandleFunc("/process-upload", ProcessUpload)
	mux.HandleFunc("/download/", ForceDownload)

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/serve-file/", ServeFile)

	mux.HandleFunc("/set-cookie", MySetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	mux.HandleFunc("/test-midleware", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("panic test")
		panic("Something Wrong")
	})

	LogMidleware := &LogMidleware{
		Handler: mux,
	}

	// midleware lapisan ke2
	ErrorHandler := &ErrorHandler{
		Handler: LogMidleware,
	}

	StartServer(mux, ErrorHandler)
}

func ServeFile(res http.ResponseWriter, req *http.Request) {
	// metode FS
	if req.URL.Query().Get("name") != "" {
		http.ServeFile(res, req, "./resources/ok.html")
	} else {
		http.ServeFile(res, req, "./resources/404-page.html")
	}

	// Metode embed
	if req.URL.Query().Get("name") != "" {
		fmt.Fprint(res, resourcesOK)
	} else {
		fmt.Fprint(res, resources404)
	}
}

func StartServer(mux http.Handler, midleware *ErrorHandler) {
	fmt.Println("Server running on : " + unittest.BaseUrl)
	server := http.Server{
		Addr: unittest.Host + ":" + strconv.Itoa(unittest.Port),
		// Handler: mux,
		Handler: midleware,
	}
	server.ListenAndServe()
}
