package main

import (
	"fmt"
	"reflect"
)

type testI1 interface {
}

type testI2 interface {
	Speak()
}

type real2 struct {
	a int
}

func (r *real2) Speak() {
	fmt.Println(reflect.ValueOf(r), reflect.TypeOf(r))
}

func main() {
	var t1 testI1
	if t1 == nil {
		fmt.Println("Is nil")
	}

	var t2 testI2
	// t2.Speak() //nil pointer will panic
	if t2 == nil {
		fmt.Println("Is nil too")
	}

	fmt.Println(reflect.TypeOf(t1))
	t1 = 2
	tv1 := reflect.ValueOf(t1)
	fmt.Println(tv1.Type())
	// var tt2 testI2
	// tt2.Speak()

	// var tt3 testI2 //没有方法的interface就是万能容器，有method，就是需要先实现方法的，就不是万能容器了
	// tt3 = 100

	a := "msdfsdf"
	c := 0
	for range a {
		c++
	}
	fmt.Println(c)
}
