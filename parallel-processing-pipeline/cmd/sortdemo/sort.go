package main

import (
	"fmt"
	"sort"
)

func main() {
	//creates a slice of int
	a := []int{3,4,65,7,9,9,12,43,5,6687}
	sort.Ints(a)
	for i, v := range a{
		fmt.Println(i, v)
	}

	for _, v := range a{
		fmt.Println(v)
	}
}
