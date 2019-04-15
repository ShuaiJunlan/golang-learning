package main

import "fmt"

func main() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(i)
	fmt.Println(i, f, u)
	test1()
}

func test1() {
	//var i int = 42
	//var f float64 = i
	//var u uint = i
	//fmt.Println(i, f, u)
	i := 't'
	fmt.Println(i)
	fmt.Printf("%T \n", i)
}
