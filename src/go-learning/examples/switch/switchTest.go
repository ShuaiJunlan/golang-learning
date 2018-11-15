package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is weekday")

	}
}
