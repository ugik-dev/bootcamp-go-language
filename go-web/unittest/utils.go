package unittest

import (
	"fmt"
	"io"
	"net/http/httptest"
	"strconv"
)

var Host string = "localhost"
var Port int = 5000
var BaseUrl string = "http://" + Host + ":" + strconv.Itoa(Port)

func GetResponseBody(wRecorder *httptest.ResponseRecorder) {
	response := wRecorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("StatusCode : ", response.StatusCode)
	fmt.Println("Status : ", response.Status)
	fmt.Println("Body : ", bodyString)
	// fmt.Println("resopon : ", response.Header.Get("x-powered-by"))

	// if bodyString != "application/json" {
	// 	// t.Errorf("Expected application/json, got %s", bodyString)
	// }
}
