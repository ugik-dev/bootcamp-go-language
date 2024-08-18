package unittest

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

func TestServerMux(t *testing.T) {
	host := "localhost"
	port := 5000

	// router
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Main Page")
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	})
	mux.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Dashboard")
	})
	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "images")
	})
	mux.HandleFunc("/images/thumbnail", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "thumbnail")
	})

	server := http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: mux,
	}

	fmt.Println("Server running on : http://" + host + ":" + strconv.Itoa(port))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on : http://" + host + ":" + strconv.Itoa(port))
}

// func TestRequest(mux *http.ServeMux) {

// }
