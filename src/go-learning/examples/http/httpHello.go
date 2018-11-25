package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(200 * time.Millisecond)
		http.ServeFile(writer, request, "front-end/static/" + request.URL.String())

	})
	http.HandleFunc("/img", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == "GET" {
			t, _ := template.ParseFiles("front-end/http2.html")
			writer.Header().Set("Content-Type", "text/html")
			_ = t.Execute(writer, nil)
		}
	})
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", nil)
	}
}
