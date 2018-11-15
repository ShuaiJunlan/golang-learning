package main

import (
	"fmt"
	"sync"
)

type User struct {
	lock sync.Mutex
	name string
	age  int
}

func main() {
	var u User
	fmt.Println(u.age)

	var u1 *User
	//fmt.Println(u1.age) //panic: runtime error: invalid memory address or nil pointer dereference

	u1 = new(User)
	fmt.Println(u1.age)
	u.lock.Lock()

	u1.name = "shuai"
	u1.age = 1212
	fmt.Println(u1)

	u2 := new(User)
	u2.lock.Lock()
	u2.name = "shuai"
	u2.age = 123
	fmt.Println(u2)
}
