package main

import (
	"fmt"
	"math"
)

//实现一个channel，需要能够传递int
//buf多路复用和

// type MyChannel struct {
// 	recvq     []int
// 	sendq     []int
// 	recvindex int
// 	sendindex int
// 	buff      unsafe.Pointer
// }

// 环形指针
// unsafe.Pointer
// uintPtr
// 为什么要用uint？？

// 实现一个hashmap
type myhash struct {
	arr     []*myNode
	hashInt int
}

type myNode struct {
	hash      int
	key       int
	value     int
	pre, next *myNode
}

func (h *myhash) get(val int) int {
	// 没找到就返回最大值
	index := val % h.hashInt
	node := h.arr[index-1]
	for node != nil {
		if node.key == val {
			return node.value
		}
		node = node.next
	}
	return math.MaxInt
}

func (h *myhash) set(key, val int) {
	index := val % h.hashInt
	if index >= len(h.arr) {
		newNode := &myNode{}
		newNode.key = key
		newNode.value = val
		newNode.pre = nil
		newNode.next = nil
		h.arr = append(h.arr, newNode)
	} else {
		node := h.arr[index-1]
		for node != nil {
			if node.key == key {
				node.value = val
			}
			node = node.next
		}
		//這裏設置flag，如果到最後都沒有就需要add一個新的内容append到最後，尾插法節點
	}
}

func (h *myhash) delete(val int) {
	index := val % h.hashInt
	if index >= len(h.arr) {
		fmt.Printf("no hash for the %d", val)
		return
	} else {
		node := h.arr[index-1]
		nodeNext := node.next
		if nodeNext == nil {
			fmt.Printf("no %d for delete", val)
			return
		} else {
			for nodeNext != nil {
				if nodeNext.key == val {
					nodeNext.next.pre = node
					node.next = nodeNext.next
				}
				node = node.next
				nodeNext = node.next
			}
			fmt.Printf("delete all value equal %d", val)
		}
	}
}

func main() {
	ch1 := make(chan<- int, 10) //表示缓冲区
	ch2 := make(chan string, 20)
	go func() {
		fmt.Printf("hello")
		ch1 <- 123
		ch2 <- "123"
	}()
}
