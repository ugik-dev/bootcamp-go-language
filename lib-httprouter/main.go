package main

import (
	"embed"
	"fmt"
	"io/fs"
	"lib-httprouter/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//go:embed resources/*
var resources embed.FS

func main() {

	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/public/*filepath", http.FS(directory))
	router.PanicHandler = utils.PanicHandler
	router.NotFound = http.HandlerFunc(utils.NotFound)

	router.MethodNotAllowed = http.HandlerFunc(utils.WrongMethod)
	router.POST("/wrong-method", Index)

	router.GET("/", Index)
	router.GET("/products", Index)
	router.GET("/products/:product_id/detail/:sub_id", ShowProduct)
	router.GET("/src/*filepath", CatchAllParams)
	router.GET("/panic", Panic)
	router.GET("/Dashboard", func(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
		fmt.Fprint(res, "Dashboard Page")
	})

	midleware := &utils.LogMidleware{Handler: router}
	fmt.Println("Server running on : " + utils.BaseUrl)
	server := http.Server{
		Addr: utils.Host + ":" + strconv.Itoa(utils.Port),
		// Handler: router,
		Handler: midleware,
	}
	server.ListenAndServe()
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!")
}

func ShowProduct(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productId := params.ByName("product_id")
	sub_id := params.ByName("sub_id")
	fmt.Fprint(w, "Product ID = "+productId+" Sub ID = "+sub_id)
}

func CatchAllParams(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	filepath := params.ByName("filepath")
	fmt.Fprint(w, "src Path = "+filepath)
}
func Panic(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	panic("Test Page Panic")
}
