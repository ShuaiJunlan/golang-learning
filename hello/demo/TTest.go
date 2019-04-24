/*
 * @Author: Shuai Junlan
 * @LastEditors: Shuai Junlan
 * @Date: 2019-04-13 20:38:07
 * @LastEditTime: 2019-04-24 09:13:55
 */
package demo // import "github.com/shuaijunlan/golang-learning/hello"

import (
	"fmt"

	"github.com/shuaijunlan/golang-learning/hello/hello1"
	"github.com/shuaijunlan/golang-learning/hello/print"
	"rsc.io/quote"
)

func MainFunction() {
	fmt.Printf("hello")
	hello1.Hello()
	print.Print()

	fmt.Println(quote.Hello())
}
