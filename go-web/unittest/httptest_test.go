package unittest

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HelloHandler")
}

func SayHello(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()
	name := r.URL.Query().Get("name")
	lastname := query["lastname"]
	address := query["address"]
	// address := r.URL.Query().Get("address")

	// if name == "" {
	// 	fmt.Fprintln(w, "Hello Anonymus")
	// } else {
	fmt.Fprintf(w, "Hello %s %s, my address %s", name, lastname, strings.Join(address, ", "))
	// }
}

func TestHttp(t *testing.T) {
	host := "localhost"
	port := 5000

	url := "http://" + host + ":" + strconv.Itoa(port)

	// rRequest := httptest.NewRequest(http.MethodGet, url, nil)
	// wRecorder := httptest.NewRecorder()
	// HelloHandler(wRecorder, rRequest)

	// rRequest := httptest.NewRequest(http.MethodGet, url+"?name=ugik+dev&lastname=Dev&address=Pringsewu&address=Bangka+Belitung", nil)
	rRequest := httptest.NewRequest(http.MethodGet, url+"?name=ugik+dev&address=Pringsewu&address=Bangka+Belitung", nil)
	wRecorder := httptest.NewRecorder()
	SayHello(wRecorder, rRequest)

	response := wRecorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println("bodyString : ", bodyString)
	// fmt.Println("body : ", body)
}
