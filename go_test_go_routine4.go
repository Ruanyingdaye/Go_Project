package main

import (
	"fmt"
)

// 几种停止go routine的方法

func main() {
	// ch := make(chan int, 1) // 需要给有i经定义大小的channel中塞东西才可以，不然也会报deadlock
	// ch := make(chan int)

	//***************************************************************************************1
	// go func() {
	// 	ch <- 1 //如果直接塞这个ch就会有问题，如果用go就没问题
	// 	// 因为如果make的直接是没有1的话，创建的是无缓冲区的channel，只要塞一个值，就会阻塞进程，此时要么将这个操作用go func（）要么就是读取用go func()才能使进程继续下去
	// 	// 不然的话就会一直阻塞ing
	// 	// 1**. 因此方法1就是，给一个无缓冲区channel塞东西，就会阻塞
	// }()

	// <-ch //2**. 方法2就是直接读取一个无缓冲区的内容，这个应该直到缓冲区关都不会继续的
	//***************************************************************************************2
	//. 测试
	// go func() {
	// 	// close(ch) //close如果去了就会发生deadlock
	// 	<-ch //从ch读会阻塞，如果写ch就会panic
	// 	ch <- 1
	// 	fmt.Println("总算关了ch")
	// }()

	// go func() {
	// 	close(ch) //close如果去了就会发生deadlock
	// 	fmt.Println("总算关了zr22")
	// }()
	// // 如果不控制的话自己就关了

	// time.Sleep(10)

	// for item := range ch {
	// 	fmt.Println(item)
	// 	close(ch) //close如果去了就会发生deadlock
	// }

	//*****************************************************************************************3
	// 区分chan和<-chan/chan<-是不同的，有的只是能输入或输出的chan
	// in1 := make(<-chan int, 10)
	// in2 := make(<-chan int, 10)
	// out := make(chan<- int, 20)

	// // go mixChannel(in1, in2, out)

	// go func(in1, in2 <-chan int, out chan<- int) {
	// 	for {
	// 		select {
	// 		//其中case 要么有case <-, 要么就是case v<-ch: 总之需要有一个接受才ok
	// 		case v := <-in1:
	// 			out <- v
	// 		case v := <-in2:
	// 			out <- v
	// 		case <-time.After(5 * time.Second):
	// 			fmt.Println("超时")
	// 			return
	// 		}
	// 	}
	// }(in1, in2, out)

	//********************************************************************************************4
	// 用channel设置缓冲区
	// 能否用waitGroup设置缓冲区
	// 	ch := make(chan int, 5)
	// 	//!!!! ch:= make(chan int,0)此处应该表示的是缓冲区为0，但是非缓冲通道，也是能存取数的，一个chan只能容纳一个而已！！
	// 	var wg sync.WaitGroup

	// 	fmt.Println("开始塞数据")
	// 	for i := 0; i < 10; i++ {
	// 		wg.Add(1)
	// 		go testFunc(i, ch)
	// 	}

	// 	// select这里一定要加一个for循环，不然不会结束的
	// 	go func() {
	// 		for {
	// 			select {
	// 			case v := <-ch:
	// 				time.Sleep(5 * time.Second)
	// 				fmt.Println("读数据", v)
	// 				wg.Done()
	// 			case <-time.After(15 * time.Second):
	// 				fmt.Println("超时了")
	// 				return
	// 			}
	// 		}
	// 	}()

	// 	// time.Sleep(100 * time.Second) //主进程不能这么快结束，一定要stop住，应该可以用其他的方法搞定，比如wait()就是等其他的协程结束了才一起结束的！
	// 	wg.Wait()
	// 	defer close(ch)
	// }

	// func testFunc(i int, ch chan int) {
	// 	ch <- i
	// 	fmt.Println(time.Now(), i)
	// }

	//************************************************************************************5
	// ch := make(chan int, 100)
	// var wg sync.WaitGroup
	// for i := 0; i < 5; i++ {
	// 	wg.Add(1)
	// 	go func(i int) {
	// 		ch <- i
	// 		defer wg.Done()
	// 	}(i)
	// }

	// go func() {
	// 	wg.Wait()
	// 	defer close(ch)
	// }() //添加了这个就可以避免死锁，死锁的原因是一直会从ch读内容，但是ch已经在等close，关闭不了就会死锁

	// for v := range ch {
	// 	fmt.Println(v)
	// } // will cause deadlock
	//************************************************************************************6
	ch := make(chan int, 100)
	// ch <- 1
	a := <-ch
	fmt.Println(a)

}
