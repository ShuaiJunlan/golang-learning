package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is weekday")

	}
	fmt.Println(rand.Intn(1000))

	test1()

	test2()
}

func test1() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
		fallthrough //?
	case "windows":
		fmt.Println("windows")
		fallthrough //?
	case "openbsd":
		fmt.Println("openbsd")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func test2() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")

	}
}
