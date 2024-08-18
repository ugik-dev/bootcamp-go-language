package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LogJson(data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}

type ObjectOne struct {
	Id          int
	Name        string
	Description string
}

func FirstJson(res http.ResponseWriter, req *http.Request) {

}
