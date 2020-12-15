package main

import (
	"fmt"
	"net/http"
	"sesi11/controllers"
)

type NewMux struct{}

func (*NewMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		switch r.Method {
		case "GET":
			controllers.ViewLogin(w, r)
		case "POST":
			controllers.Login(w, r)
		}
	case "/user":
		controllers.Result(w, r)
	default:
		http.NotFound(w, r)
		return
	}
}

func main() {
	mux := &NewMux{}
	fmt.Println("Listen http on port :9000")
	http.ListenAndServe(":9000", mux)
}
