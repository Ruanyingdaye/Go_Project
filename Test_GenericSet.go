package main

import "fmt"

/// new Test
// 利用map+泛型实现一个Set
type Set[T comparable] map[T]struct{}

// 其中comparable
// 其中泛型只有定义的时候才需要有T的类型，后面都是直接使用这个类型了，就用T就可以了
func (s Set[T]) newSet(elements ...T) Set[T] {
	//这里可以设置多个参数，可变参数
	//用map的遍历方式对应elements的遍历
	newset := make(Set[T])
	for _, ele := range elements {
		//!!!struct{}{}, struct{}是一种类型，表示空结构体，struct{}{}是其对应的类型
		newset[ele] = struct{}{}
	}
	return newset
	//返回的是结构体本身，其大小主要取决于里面的元素，里面的内容主要包括key和value
}

func (s Set[T]) Add(t T) {
	//map操作的时候，不需要返回值，因为map底层是指针
	s[t] = struct{}{}
}
func (s Set[T]) Remove(t T) {
	delete(s, t)
}

func (s Set[T]) Contains(t T) bool {
	if _, ok := s[t]; ok {
		return true
	}
	return false
}

// func (s Set[T]) Union(other Set[T]) Set[T] {
// 	newset := make(Set[T])
// 	for k := range s {
// 		newset[k] = struct{}{}
// 	}
// 	for k := range other {

// 	}
// }

// func Union

type newInt[T comparable] struct{}

func testFunc() {
	// ni := make(newInt[string])
	// make只能用于slice channel map这三种类型及其衍生类型，其他的类型都是会报错的，只能用Func(){}的方式创建
	ni := newInt[int32]{}
	fmt.Println(ni)
}

type newType[T interface{ ~*int | ~*float32 }] []T

type NewFunc func(a, b int) int

func add(a, b int) int {
	return a + b
}
func main() {
	var newFunc NewFunc = add
	fmt.Println(newFunc(1, 2))

	testmap := make(map[int]int)
	testmap[1] = 1
	delete(testmap, 1)
	if _, ok := testmap[1]; ok {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}

}
