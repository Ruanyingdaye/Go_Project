package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

// ***********************************
//
//	func main() {
//		for i := 0; i < 10; i++ {
//			go func() {
//				fmt.Println(i)
//			}()
//		}
//	}
//
// 这种情况下是不会有具体的output的！因为没有wait，主程序直接关闭了。。
// ************************************
func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	//trace用于查看内存泄露

	//可以加锁, 这个不是锁是sync
	var wg sync.WaitGroup
	ch := make(chan interface{})
	go func() {
		time.Sleep(10 * (time.Second)) //当将这个ch关了之后，反而可以<-ch可以得到内容了，基于go routine的顺序控制
		close(ch)
	}()

	// close(ch)
	// <-  ch //当这个ch没有东西的时候，所有携程都会卡在这里

	// wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Add(1)
			fmt.Println("a1", i)
			if err := recover(); err != any(nil) {
				fmt.Println("panic")
			}
			wg.Done() //和Add()是相对应的 //这个是乱序的
		}(i) //注意go func中的函数会是一个闭包，因此如果想拿到传过来的值，最好就go func(i)这样，不然可能i就是随机值了
		// 同时run的多个go routine()一定会乱序的
	}

	wg.Wait()

}
