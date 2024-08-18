package main

import (
	"html/template"
	"net/http"
)

// go run .\template.go .\main.go

func SimpleHTML(res http.ResponseWriter, req *http.Request) {
	textTemplate := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(textTemplate)
	// if err != nil {
	// 	panic(err)
	// }

	t := template.Must(template.New("SIMPLE").Parse(textTemplate))
	t.ExecuteTemplate(res, "SIMPLE", "Template Engine")
}

func SimpleHTMLFile(res http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))
	t.ExecuteTemplate(res, "simple.gohtml", "Template File Engine")
}

// rekomendasi
func TemplateEmbed(res http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(res, "simple.gohtml", "Template File Engine")
}

func TemplateDataMap(res http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	data := map[string]interface{}{
		"Title": "Templating Parsing Data",
		"Name":  "Fullstack Developer",
		"Info": map[string]interface{}{
			"Address": "Jalan dari Map",
			"Phone":   "08MAP",
		},
	}
	t.ExecuteTemplate(res, "parsing_data.gohtml", data)
}

type InfoUser struct {
	Address string
	Phone   string
}

type Page struct {
	Title string
	Name  string
	Info  InfoUser
}

type Task struct {
	Index       int
	Name        string
	Description string
}

func TemplateDataStruct(res http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	data := Page{
		Title: "Title Struct",
		Name:  "Fullstack Developer - Struct Data",
		Info: InfoUser{
			Address: "JLRaya",
			Phone:   "081279748967",
		},
	}
	t.ExecuteTemplate(res, "parsing_data.gohtml", data)
}

func TemplateDataRange(res http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	tasks := []Task{
		{1, "Frontend", "React JS, Vue JS"},
		{2, "Backend", "Golang, Express JS"},
		{3, "Data Analys", "MYSQL, MongoDB"},
	}
	data := map[string]interface{}{
		"Title": "Templating Action Range",
		"Name":  "Ugik Fullstack Developer",
		"Tasks": tasks,
	}

	t.ExecuteTemplate(res, "action_range.gohtml", data)
}

func TemplateAction(res http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	/**
	eq equal =
	ne not equal !=
	lt lest than <
	le lest equal <=
	gt greater than >
	ge greater equal >=


	 --GLobal Function
	 https://github.com/golang/go/blob/master/src/text/template/funcs.go
		FuncMap{
		"and":      and,
		"call":     emptyCall,
		"html":     HTMLEscaper,
		"index":    index,
		"slice":    slice,
		"js":       JSEscaper,
		"len":      length,
		"not":      not,
		"or":       or,
		"print":    fmt.Sprint,
		"printf":   fmt.Sprintf,
		"println":  fmt.Sprintln,
		"urlquery": URLQueryEscaper,

		// Comparisons
		"eq": eq, // ==
		"ge": ge, // >=
		"gt": gt, // >
		"le": le, // <=
		"lt": lt, // <
		"ne": ne, // !=
	*/
	data := map[string]interface{}{
		"Title": "Templating Parsing Data",
		"Name":  "Fullstack Developer",
		"Phone": "",
		"Age":   25,
	}

	t.ExecuteTemplate(res, "action.gohtml", data)
}

type MyPage struct {
	Title string
	Name  string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", my name is" + myPage.Name
}

func MyFunctionTest(value string) string {
	return "hello " + value
}
func TemplateFunction(res http.ResponseWriter, req *http.Request) {
	// t := template.New("FUNCTION")
	// t = t.Funcs(map[string]interface{}{
	// 	"upper": func(value string) string {
	// 		return strings.ToUpper(value)
	// 	},
	// })

	// registrasi function custom
	t := template.New("template").Funcs(template.FuncMap{
		// "upper": func(value string) string {
		// 	return strings.ToUpper(value)
		// },
		"MyFunctionTest": MyFunctionTest,
		"secound":        MyFunctionTest,
	})
	// GUNAKAN INI baris atas karna untuk identifikasi function upper terlebih dahulu

	t = template.Must(t.ParseFS(templates, "templates/*.gohtml"))
	// t = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	data := MyPage{
		Title: "Templating Function",
		Name:  "Ugik Developer",
	}
	t.ExecuteTemplate(res, "function.gohtml", data)
}
