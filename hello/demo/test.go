/*
 * @Author: Shuai Junlan
 * @LastEditors: Shuai Junlan
 * @Date: 2019-04-13 20:38:07
 * @LastEditTime: 2019-04-15 15:54:53
 */
package main // import "github.com/shuaijunlan/golang-learning/hello"

import (
	"fmt"
	"hello"

	"rsc.io/quote"
)

func main() {
	fmt.Printf("hello")
	hello.Hello()

	fmt.Println(quote.Hello())
}
