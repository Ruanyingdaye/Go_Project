package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//数据可以定义结构体数组，函数数组， 接口数组等
	interfaceS := []interface{}{123, "234", []int{1, 2, 3}} //接口可以容纳一切类型
	for _, item := range interfaceS {
		// 一个参数表示的是i
		// if v, ok := item.(int); ok {
		// 	fmt.Printf("%d", v)
		// }
		switch item.(type) {
		case int:
			fmt.Println(item)
		case string:
			fmt.Println(item)
		case []int:
			for _, v := range item.([]int) {
				fmt.Printf("slice: %d\n", v)
			}
		default:
		}
	}
	s := struct{}{}               //定义空结构体要两层
	fmt.Println(unsafe.Sizeof(s)) //这个是对象的固定占用大小 golang中为0， C++为1

}
