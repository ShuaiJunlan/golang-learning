package main

import (
	"fmt"
	"time"
)

func main() {
	o := make(chan int)
	c := make(chan int)
	go func() {
		for {
			select {
			case a := <-c:
				fmt.Println(a)
			case <-time.After(5 * time.Second):
				o <- 0
				break
			}
		}
	}()
	for i := 0; i < 100; i++ {
		c <- i
	}
	<- o
}