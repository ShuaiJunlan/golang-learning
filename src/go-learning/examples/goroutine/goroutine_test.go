package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func printString(t string) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println(t)
}
func TestGoroutine(t *testing.T) {
	go printString("hello goroutine!")
	fmt.Println("hello main")
}

func TestGoroutineWithGosched(t *testing.T) {
	go printString("hello goroutine!")
	runtime.Gosched()
	fmt.Println("hello main")
}

func TestGoroutine1(t *testing.T) {
	go println("go")
	runtime.Gosched()
}

func TestGoroutine2(t *testing.T) {
	name := "shuai"
	go func() {
		fmt.Printf("hello %s! \n", name)
	}()
	name = "junlan"
	time.Sleep(1000 * time.Millisecond)
}

/*
output:
Hello, lan !
Hello, lan !
Hello, lan !
why? how to solve it
*/
func TestGoroutine3(t *testing.T) {
	names := []string{"shuai", "jun", "lan"}
	for _, name := range names {
		go func() {
			fmt.Printf("Hello, %s !\n", name)
		}()
	}
	time.Sleep(1000 * time.Millisecond)
}
func TestGoroutine3Solving(t *testing.T)  {
	names := []string{"shuai", "jun", "lan"}
	for _, name := range names{
		go func(s string) {
			fmt.Printf("Hello, %s ! \n", s)
		}(name)
	}
	time.Sleep(1000 * time.Millisecond)
}