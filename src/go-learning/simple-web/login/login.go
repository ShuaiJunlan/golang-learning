package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie")
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		currentTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(currentTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("front-end/login.html")
		w.Header().Set("Content-Type", "text/html")
		log.Println(t.Execute(w, token))
	} else {
		r.ParseForm() //?? So?

		if len(r.Form["username"][0]) == 0 {
			fmt.Fprintf(w, "username can't be null!")
			return
		}

		if len(r.Form["password"][0]) == 0 {
			fmt.Fprintf(w, "password isn't correct")
			return
		}

		fmt.Println("username", r.Form["username"])
		fmt.Println("password", r.Form["password"])

		fmt.Fprintf(w, "login successfully")
	}
}

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
