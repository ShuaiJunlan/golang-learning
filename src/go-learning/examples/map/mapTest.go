package main

import "fmt"

func main() {
	ages := make(map[string]int) //mapping from strings to ints
	ages["shuai"] = 32
	ages["junlan"] = 232
	fmt.Println(ages)
	ages1 := map[string]int{
		"shuai":  32,
		"junlan": 232,
	}
	delete(ages1, "shuai")
	fmt.Println(ages1)

	for name, age := range ages {
		fmt.Printf("%s %d \n", name, age)
	}

}

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) bool {
	edges := graph[from]
	if edges != nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
