package main // import "github.com/shuaijunlan/golang-learning/vgo-test"

import (
	"fmt"

	"github.com/shuaijunlan/golang-learning/hello"
	"rsc.io/quote"
)

func main() {
	fmt.Printf("hello")
	hello.Hello()

	fmt.Println(quote.Hello())
}
