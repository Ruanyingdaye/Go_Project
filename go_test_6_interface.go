package main

import (
	"fmt"
	"reflect"
)

type Animal interface {
	move()
	fight()
	yell(a Animal)
}

// struct/interface名字一定要大写

type Cat struct{}
type Dog struct {
	//这里面只有变量，没有方法
}

func (c Cat) move() {
	fmt.Println("This is the cat move")
}

func (c Cat) fight() {
	fmt.Println("This is the cat fight")
} // add the implement will make the interface can be assignment

func (c Cat) yell(a Animal) {
	fmt.Printf("Hell, what's the matter with you %s?\n", reflect.TypeOf(a).Name())
} // add the implement will make the interface can be assignment

func (d Dog) move() {
	fmt.Println("This is the cat move")
}

func (d Dog) fight() {
	fmt.Println("This is the dog fight")
}

func (d Dog) yell(a Animal) {
	fmt.Printf("Hell, what's the matter with you %s?\n", reflect.TypeOf(a).Name())
} // add the implement will make the interface can be assignment

func main() {
	cat := Cat{}
	dog := Dog{}
	var cat2 Animal
	cat.move()     //不用全部实现所有的func，结构体的方法也能实现，但是这个结构体可能属于不全面的，因此没有办法作为右值进行赋值
	cat2 = cat     //如果没有cat的fight函数则无法用cat2作为cat的实现进行赋值，因为不是所有的方法都实现
	cat2.yell(dog) // 只有实现了所有方法的函数，才能作为interface的变量进行赋值
}
