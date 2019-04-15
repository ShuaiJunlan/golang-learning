package main

import "fmt"

//TODO: https://blog.golang.org/defer-panic-and-recover
func main() {
	test1()
	test2()
	test3()
}
func test1() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func test2() {
	i := 10
	defer fmt.Println("world", i)
	i++
	fmt.Println("hello", i)
}

func test3() {
	fmt.Println("counting")
	for i := 0; i < 1000000; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
