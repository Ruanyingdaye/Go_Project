package main

import (
	"fmt"
)

// struct中没有func，或者只能有匿名函数，否则就是要实现了，这个应该作为方法有方法的接收器
// interface中只能定义方法，不能定义变量，因为interface{}的大小就是1
// 比如这样type Printer interface {
// Print()
// }

type Test interface {
	Init() interface{}
	Exec(func(interface{}) interface{}) interface{}
	Validate(interface{}) bool
}

// type NewTest1 struct{
// Log()
// }
// func (n *NewTest1) Init() interface{}      { return nil }
// func (n *NewTest1) Exec(f func(), i interface{}) {}

// *************************************start for test1****************************************
// Reverse Link Node
type NewTest1 struct {
}

func (n *NewTest1) Init() *LinkNode {
	//
	dummyHead := &LinkNode{}
	curr := dummyHead
	cnt := 5
	for cnt > 0 {
		tmp := &LinkNode{}
		tmp.val = cnt
		curr.next = tmp
		curr = curr.next
		// fmt.Println(tmp.val)
		cnt--
	}
	return dummyHead.next
}

func (n *NewTest1) Exec(f func(*LinkNode) *LinkNode) *LinkNode {
	// exec
	return f(n.Init())
}

func (n *NewTest1) Validate(node *LinkNode) bool {
	//define the target
	source := []int{1, 2, 3, 4, 5}
	//////////////////////////////////////////////
	target := []int{}
	for node != nil {
		target = append(target, node.val)
		node = node.next
	}
	if len(target) == len(source) {
		for i := 0; i < 5; i++ {
			if source[i] != target[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

// add test1 struct
type LinkNode struct {
	val  int
	next *LinkNode
}

func reverseLinkNode(node *LinkNode) *LinkNode {
	if node == nil {
		return nil
	}
	dummyHead := &LinkNode{}
	dummyHead.next = nil
	for node != nil {
		tmp := node.next
		node.next = dummyHead.next
		dummyHead.next = node
		node = tmp
	}
	return dummyHead.next
}

//**************************************end for test1****************************************

// **************************************start for test2**************************************
// QuickSort
type NewTest2 struct {
	source []int
	target []int
}

func (n *NewTest2) Init() {
	n.source = []int{3, 2, 1, 4, 6, 5}
}

func (n *NewTest2) Exec(f func([]int, int, int)) {
	// slice排序底层就是对array排序，传递的是引用，因此可以直接修改
	// exec
	f(n.source, 0, len(n.source)-1)
	n.target = append(n.target, n.source...)
}

func (n *NewTest2) Validate(arr []int) bool {
	//匿名函数或者具名函数都能达到同样的效果，不过struct中是没有承载匿名函数的变量的，只有函数体本身
	//define the source
	source := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(source, arr)
	//////////////////////////////////////////////
	if len(arr) == len(source) {
		for i := 0; i < 6; i++ {
			if source[i] != arr[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func quickSort(s []int, i, j int) {
	//quickSort shouldn't have return value
	if j-i < 1 {
		return
	}
	flag := s[i]
	start, end := i, j
	for i < j {
		for s[j] >= flag && i < j {
			j--
		}
		s[i], s[j] = s[j], s[i]
		if i >= j {
			break
		}
		i++
		for s[i] <= flag && i < j {
			i++
		}
		s[i], s[j] = s[j], s[i]
		if i >= j {
			break
		}
		j--
	}

	s[i] = flag
	quickSort(s, start, i-1)
	quickSort(s, i+1, end)

}

// **************************************end for test2****************************************
func main() {
	//test1
	// test1 := &NewTest1{}
	// test1Target := test1.Exec(reverseLinkNode)
	// fmt.Println(test1.Validate(test1Target))

	// test2
	test2 := &NewTest2{}
	test2.Init()
	test2.Exec(quickSort)
	fmt.Println(test2.Validate(test2.target))

}
