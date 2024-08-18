package unittest

import (
	"fmt"
	"net/http"
	"strconv"
	"testing"
)

func TestServer(t *testing.T) {
	host := "localhost"
	port := 5000

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "Server running on : http://"+host+":"+strconv.Itoa(port))
	}

	server := http.Server{
		Addr:    host + ":" + strconv.Itoa(port),
		Handler: handler,
	}

	fmt.Println("Server running on : http://" + host + ":" + strconv.Itoa(port))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

	fmt.Println("Server running on : http://" + host + ":" + strconv.Itoa(port))
}
