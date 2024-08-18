package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func PanicHandler(res http.ResponseWriter, req *http.Request, err interface{}) {
	fmt.Fprint(res, "Panic : ", err)
}

func NotFound(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Ups!!, Halaman yang anda cari tidak ditemukan")
}

func WrongMethod(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Ups!!, anda mau ngapain?")
}

func GetResponseBody(wRecorder *httptest.ResponseRecorder) {
	response := wRecorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("StatusCode : ", response.StatusCode)
	fmt.Println("Status : ", response.Status)
	fmt.Println("Body : ", bodyString)
}

type LogMidleware struct {
	http.Handler
}

func (midleware *LogMidleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Before LogMidleware")
	fmt.Println("---- Execute : ", req.URL.Path)
	midleware.Handler.ServeHTTP(res, req)
	fmt.Println("After LogMidleware")
}

type ErrorHandler struct {
	Handler http.Handler
}
