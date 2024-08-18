package main

import (
	"fmt"
	"net/http"
)

type LogMidleware struct {
	Handler http.Handler
}

func (midleware *LogMidleware) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Before LogMidleware")
	midleware.Handler.ServeHTTP(res, req)
	fmt.Println("After LogMidleware")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Before errorHandler")
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("show error message")
			res.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(res, "Terjadi kesalahan!!")
		}
	}()
	errorHandler.Handler.ServeHTTP(res, req)

	fmt.Println("After errorHandler")
}
