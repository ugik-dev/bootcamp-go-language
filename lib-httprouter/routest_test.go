package main

import (
	"io"
	"lib-httprouter/utils"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", Index)

	request := httptest.NewRequest("GET", utils.BaseUrl+"/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	assert.Equal(t, "Welcome!", bodyString)
}

func TestGetProduct(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:product_id", ShowProduct)

	request := httptest.NewRequest("GET", utils.BaseUrl+"/products/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)
	assert.Equal(t, "Product ID = 2", bodyString)
}
