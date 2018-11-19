package _for

//func TestForLoop(t *testing.T) {
//	//fmt.Println(getIntI())
//	fmt.Println(getIntJ())
//}
func getIntI() (i int32) {
	for {
		i = 10
		break
	}
	return
}
func getIntJ1() (j int32) {
	for {
		j = 20
		if j == 21 {
			continue
		}
		return
	}
}
func main() {

}
