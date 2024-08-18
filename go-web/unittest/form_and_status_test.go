package unittest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	firstName := req.PostForm.Get("first_name")
	lastName := req.PostForm.Get("last_name")
	if firstName == "" {
		// Menambahkan Status code custom, defaulnya 200
		res.WriteHeader(404)
		fmt.Fprintf(res, "First name not found!")
	} else {
		res.WriteHeader(200)
		fmt.Fprintf(res, "Hello %s %s", firstName, lastName)
	}
}

func TestPostForm(t *testing.T) {

	reqBody := strings.NewReader("first_name&last_name=Developer")
	rRequest := httptest.NewRequest(http.MethodPost, BaseUrl+"", reqBody)
	rRequest.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	wRecorder := httptest.NewRecorder()
	FormPost(wRecorder, rRequest)
	// fmt.Println("rRequest.Header = ", rRequest.Header.Get("X-Powered-By"))
	GetResponseBody(wRecorder)
}
