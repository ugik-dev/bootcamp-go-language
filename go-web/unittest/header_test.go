package unittest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func ResponWrite(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("X-Powered-By", "Ugik Developer")
	fmt.Fprint(res, "OK")
}

func TestReqHeader(t *testing.T) {
	rRequest := httptest.NewRequest(http.MethodPost, BaseUrl+"/get-header", nil)
	rRequest.Header.Add("Content-Type", "application/json")
	wRecorder := httptest.NewRecorder()

	// menambah custom HEADER KERESPONSE
	wRecorder.Header().Add("X-Powered-By", "Ugik Fullstack Developer")
	// RequestHeader(wRecorder, rRequest)
	ResponWrite(wRecorder, rRequest)
	fmt.Println("rRequest.Header = ", rRequest.Header.Get("X-Powered-By"))
	GetResponseBody(wRecorder)
}
