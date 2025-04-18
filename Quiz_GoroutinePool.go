package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine pool
// producer/consumer pool

const procNum = 10

type PCBase interface {
	Producer()
	Consumer()
}

type PCInt struct {
	ch chan int
}

func ProducerFunc(i int, ch chan int) {
	ch <- i
	fmt.Println("Input:", i)
}

func ConsumerFunc(i int, ch chan int) {
	rs := <-ch
	fmt.Println("Output:", rs)
}

func (pc *PCInt) Producer(wg *sync.WaitGroup) {
	for i := range procNum {
		wg.Add(1)
		go ProducerFunc(i, pc.ch)
		wg.Done()
	}
	wg.Wait()
}

func (pc *PCInt) Consumer(wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	for i := range procNum {
		wg.Add(1)
		go ConsumerFunc(i, pc.ch)
		wg.Done()
	}
	wg.Wait()
}

type Constructor func() *PCInt

/*
1. 定义Constructor作为统一的构造函数
2. 定义每一个具体的构造函数
3. 可以用map[string]Constructor来进行统一管理
*/

//slice map func是不能作为map的key的

func main() {

	IntConFunc := func() *PCInt {
		return &PCInt{
			ch: make(chan int),
		}
	}
	var ccFunc Constructor = IntConFunc
	cc := ccFunc()
	var wg sync.WaitGroup
	cc.Producer(&wg)
	cc.Consumer(&wg)

	//这样肯定是错误的，需要再最外层开启go func()的协程，读写才能同步！！！里面是不太行的
	// go func go Add go wait最好都在最外层！！
}
