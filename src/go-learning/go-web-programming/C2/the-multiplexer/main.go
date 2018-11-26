package main

import (
	"go-learning/go-web-programming/handler"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("front-end"))

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("logout", logout)
	mux.HandleFunc("singnup", signup)
	mux.HandleFunc("singnup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)
	mux.HandleFunc("/thread/post", postThread)
	mux.HandleFunc("/thread/read", readThread)


	server := &http.Server{
		Addr:":8080",
		Handler:mux,
	}
	_ = server.ListenAndServe()
}

func readThread(writer http.ResponseWriter, request *http.Request) {

}

func postThread(writer http.ResponseWriter, request *http.Request) {

}

func createThread(writer http.ResponseWriter, request *http.Request) {

}

func newThread(writer http.ResponseWriter, request *http.Request) {

}

func authenticate(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	cookie := http.Cookie{
		Name:"hello",
		Value:"hhh",
		HttpOnly:true,
	}
	http.SetCookie(writer, &cookie)

}

func signupAccount(writer http.ResponseWriter, request *http.Request) {

}

func signup(writer http.ResponseWriter, request *http.Request) {

}

func logout(writer http.ResponseWriter, request *http.Request) {

}

func login(writer http.ResponseWriter, request *http.Request) {

}


