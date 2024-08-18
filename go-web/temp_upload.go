package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func FormUpload(res http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"Title": "Form Upload",
		"Name":  "Fullstack Developer",
	}
	myTemplates.ExecuteTemplate(res, "form_upload.gohtml", data)
}

func ProcessUpload(res http.ResponseWriter, req *http.Request) {
	file, fileHeader, err := req.FormFile("attachment")
	if err != nil {
		panic(err)
	}
	fmt.Println(fileHeader.Filename)
	fileDestination, err := os.Create("./resources/upload/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := req.PostFormValue("name")
	data := map[string]interface{}{
		"Name":     name,
		"FileLink": "/static/upload/" + fileHeader.Filename,
	}
	myTemplates.ExecuteTemplate(res, "form_result.gohtml", data)

}

func ForceDownload(res http.ResponseWriter, req *http.Request) {

	filePath := req.URL.Query().Get("file")
	if filePath == "" {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(res, "Bad Request")
		return
	}
	res.Header().Add("Content-Disposition", "attachment; filename=\""+filePath+"\"")
	http.ServeFile(res, req, "./resources/upload/"+filePath)

}
