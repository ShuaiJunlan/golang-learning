package main

import "fmt"

func main() {
	ages := make(map[string]int) //mapping from strings to ints
	ages["shuai"] = 32
	ages["junlan"] = 232
	fmt.Println(ages)
	ages1 := map[string]int{
		"shuai": 32,
		"junlan": 232,
	}
	delete(ages1, "shuai")
	fmt.Println(ages1)

	for name, age := range ages{
		fmt.Printf("%s %d \n", name, age)
	}


}
