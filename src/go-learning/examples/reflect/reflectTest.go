package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Age  int
}

type Hello interface {
	Print(prefix string)
}

//http://www.flysnow.org/2017/06/13/go-in-action-go-reflect.html
func main() {
	u := User{"shuaijunlan", 23}
	t := reflect.TypeOf(u)
	fmt.Println(t)

	fmt.Println(t.Kind())

	v := reflect.ValueOf(u)
	fmt.Println(v)

	t1 := v.Type()
	fmt.Println(t1)

	u1 := v.Interface().(User)
	fmt.Println(u1)

	mPrint := v.MethodByName("Print")
	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))


	u.Print("hello")
}

func (u User)Print(prefix string) {
	fmt.Printf("%s:Name is %s, Age is %d \n", prefix, u.Name, u.Age)
}