package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("front-end"))
	mux.HandleFunc("/", index)

	mux.Handle("/static/", http.StripPrefix("/static/", files))

	server := &http.Server{
		Addr:":8080",
		Handler:mux,
	}
	_ = server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {

}
