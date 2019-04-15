package main

import (
	"fmt"
	"golang.org/x/net/http2"
	"html/template"
	"log"
	"net/http"
	"time"
)
// TODO: https://colobu.com/2018/09/06/Go-http2-%E5%92%8C-h2c/
//  https://yq.aliyun.com/articles/640871
const idleTimeout = 5 * time.Minute
const activeTimeout = 5 * time.Minute
func main() {
	var srv http.Server
	srv.Addr = ":8972"
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.URL)

		//_, _ = writer.Write([]byte("hello http2"))
		time.Sleep(200 * time.Millisecond)
		http.ServeFile(writer, request, "front-end/static/" + request.URL.String())
	})

	http.HandleFunc("/img", writeImage)

	_ = http2.ConfigureServer(&srv, &http2.Server{})
	go func() {
		log.Fatal(srv.ListenAndServeTLS("examples/http2/server.crt", "examples/http2/server.key"))
	}()
	select {

	}
}

func writeImage(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		t, _ := template.ParseFiles("front-end/http2.html")
		writer.Header().Set("Content-Type", "text/html")
		_ = t.Execute(writer, nil)
	}
	fmt.Println(request.URL)

	//if request.URL.s {
	//
	//}

	//http.ServeFile(writer, request, "front-end/static/images/f-4.jpg")
}