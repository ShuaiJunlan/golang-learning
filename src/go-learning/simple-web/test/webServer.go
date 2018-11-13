package test

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
func sayHelloName(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	fmt.Println(request.Form)
	fmt.Println("path", request.URL.Path)
	fmt.Println("scheme", request.URL.Scheme)
	fmt.Println(request.Form["url_long"])
	for k, v := range request.Form{
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	time.Sleep(time.Millisecond * 50)
	fmt.Fprintf(writer, "hello astaxie!")
}
