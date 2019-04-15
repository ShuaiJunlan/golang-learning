package _for

import "fmt"

func main() {
	fmt.Println(getIntJ())
}
func getIntJ() (j int32) {
	for true {
		j = 20
		if j == 21 {
			continue
		}
		return
	}
}
