package unittest

import (
	"fmt"
	"net/http"
	"strconv"
)

func MySetCookie(res http.ResponseWriter, req *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Username"
	cookie.Value = "Username_" + req.URL.Query().Get("first_name")
	cookie.Path = "/"

	cookieToken := new(http.Cookie)
	cookieToken.Name = "X-Token"
	cookieToken.Value = "Token_" + req.URL.Query().Get("first_name")
	cookieToken.Path = "/"

	http.SetCookie(res, cookie)
	http.SetCookie(res, cookieToken)
	fmt.Fprint(res, "Success Create Cookies")
}

func GetCookie(res http.ResponseWriter, req *http.Request) {
	cookieUser, err := req.Cookie("X-Username")

	if err != nil {
		fmt.Fprint(res, err.Error())
	} else {
		valUsername := cookieUser.Value
		fmt.Fprintf(res, "Helo %s", valUsername)

	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", MySetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)
	StartServer(mux)
}

func StartServer(mux http.Handler) {
	fmt.Println("Server running on : " + BaseUrl)
	server := http.Server{
		Addr:    Host + ":" + strconv.Itoa(Port),
		Handler: mux,
	}
	server.ListenAndServe()
}
