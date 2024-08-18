package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

// gunakan import "html/template" untuk escape bukan text/template
//
//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(res http.ResponseWriter, req *http.Request) {
	data := map[string]interface{}{
		"Title": "Templating Caching",
		"Name":  "Fullstack Developer",
		"Info": map[string]interface{}{
			"Address": "Jalan dari Map",
			"Phone":   "08MAP",
		},
	}
	myTemplates.ExecuteTemplate(res, "caching.gohtml", data)
}

func TemplateXxs(res http.ResponseWriter, req *http.Request) {
	// selain template.HTML ada template.CSS template.JS dll.
	data := map[string]interface{}{
		"Title":        "Templating Caching",
		"Name":         "Fullstack Developer",
		"InjectSript":  "<script> alert('Hacked') </script>",
		"NoAutoEscape": template.HTML("<p style='color: blue'>Ini konten dengan html</p>"),
		"Info": map[string]interface{}{
			"Address": "Jalan dari Map",
			"Phone":   "<b>08MAP <script> alert('Hacked') </script> <b>",
		},
	}
	myTemplates.ExecuteTemplate(res, "caching.gohtml", data)
}

func RedirectTo(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Pesan Redirect")
}

func TemplatrRedirect(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/template/redirect-to", http.StatusTemporaryRedirect)
}
func RedirectUgikSite(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "https://ugikdev.site/", http.StatusPermanentRedirect)
}
