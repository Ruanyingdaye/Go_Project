package main

import (
	"fmt"
	"reflect"
)

type Generic[T int | string | float32] struct {
	value T
}

type GenericPro[T ~int | ~string | ~float32] struct {
	value T
}

type GenericProMax[T interface{ *int | *float32 }] struct {
	value T
}

func (g Generic[T]) Log() {
	fmt.Println(g.value)
}

func (g GenericPro[float32]) Log() {
	fmt.Println(g.value, "for log")
}

func (g GenericProMax[string]) getCount() {
	fmt.Println(g.value)
}

func TestGenericFunc[T interface{ *int | *float32 }](t T) bool {
	fmt.Println(t)
	return true
}

func TestGenericFunc2[T any](t T) bool {
	if reflect.ValueOf(t).IsNil != nil {
		typea := reflect.TypeOf(t)
		fmt.Println(typea)
	}
	fmt.Println(t)
	return true
	//如果想基于不同的类型进行输出的话可以采用reflect去判断类型
}

//t应该是是一个interface这样的类型才可以，而泛型在具体编译的时候会转化为具体的类型，而具体的类型是不能用switch case的！

func main() {
	//struct定义不用make
	g := Generic[int]{
		value: 1,
	}
	g.Log()

	// test the GenericPro
	gp := GenericPro[string]{
		value: "aaaa",
	}
	gp.Log()

	TestGenericFunc(&g.value) //因为interface中传递的内容是一个指针，因此参数要&来进行表示
	TestGenericFunc2(&gp.value)
}
