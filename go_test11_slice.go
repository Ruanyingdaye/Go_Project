package main

import (
	"context"
	"fmt"
)

func aaaa(d context.Context) {
	ch := make(chan int, 100)
	for {
		select {
		case <-ch:
		default:
		}
	}
}

func main() {
	// dp := make([][]int, 10)
	// for i, _ := range dp {
	// 	dp[i] = make([]int, 3)
	// }

	// s := "afsldfjsdklfj"
	// for _, item := range s {
	// fmt.Println(reflect.TypeOf(item), item)
	// fmt.Println(item) //rune类型，打印出来就是int32
	// }
	// c := make(chan int, 5)
	// c <- 3
	// c <- 2
	// close(c)

	// if tmp, ok := <-c; ok {
	// 	fmt.Println(tmp, c)
	// }

	// if tmp, ok := <-c; ok {
	// 	fmt.Println(tmp, c)
	// }

	// if tmp, ok := <-c; ok {
	// 	fmt.Println(tmp, c)
	// } //后面不ok了，就不走这个逻辑了，但是也不会阻塞，值其实就是默认的0，关闭并不影响读取已经在channel中的数

	// testch := make(chan int)
	// ctx, cancel := context.WithCancel(context.Background()) //注意此处返回的应该是一个函数
	// defer cancel()

	// go func(ctx context.Context) {
	// 	testch <- 1
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println("done")
	// 			return
	// 		default:
	// 			fmt.Println("abc")
	// 		}
	// 	}
	// }(ctx)
	// v, ok := <-testch
	// fmt.Println(v, ok)
	// cancel() //这里调用对应的应该是select中的ctx.Done，外面的cancel对应于ctx.Done对应的操作
	// close(testch)

	// v, ok = <-testch
	// fmt.Println(v, ok) //此处当关闭之后 ok=false，值为0

	//test字符串
	s := "abceefg"
	//输出s[i]的类型
	// 遍历s，遍历的每一个值都是rune类型，转化为byte，请写一个例子

	for i := 0; i < len(s); i++ {
		fmt.Printf("%c", s[i])
	}

}
