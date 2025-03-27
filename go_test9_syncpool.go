package main

import (
	"fmt"
	"sync"
)

type MyObject struct {
	value int
}

func main() {
	// 创建一个对象池
	objectPool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new object")
			return &MyObject{}
		},
	}

	// 从池中获取对象
	obj1 := objectPool.Get().(*MyObject)
	obj1.value = 42
	fmt.Println("Object 1:", obj1)

	// 归还对象到池中
	objectPool.Put(obj1)

	// 再次从池中获取对象，应该是同一个对象
	obj2 := objectPool.Get().(*MyObject)
	fmt.Println("Object 2:", obj2)
}
