package project

import (
	"fmt"
	"sync"
	"time"
)

//实现生产者消费者，生产者生产x个东西放到channel中，然后消费者分为若干协程去消费这些内容
//实现一个自适应的队列，从而提升或者下降消费者协程的数量，以及对应的消费能力

type Base interface {
	Producer()
	Consumer()
}

// struct继承interface，只需要实现对应的方法即可
type TestConsumeA struct {
	Pipe           chan int64 // for consume the data produced
	PipeNum        int64
	ProduceDataNum int64
}

func (t TestConsumeA) Producer(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done() //需要打开一个协程通知生产者生产完毕，因此也是需要的
	for item := range t.ProduceDataNum {
		t.Pipe <- item
	}
}

func (t TestConsumeA) Consumer(m map[int64]struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	wg.Done()
	var mu sync.Mutex
	for _ = range t.PipeNum {
		//如果没有new的东西是可以不用:=进行range的，但是一般为了log，也会添加一个！
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			msg := <-t.Pipe
			m[msg] = struct{}{}
			mu.Unlock()
		}()
	}
	wg.Wait()
}

func (t TestConsumeA) StopChannel() {
	timeout := time.After(10 * time.Second)
	go func() {
		select {
		case msg := <-t.Pipe:
			fmt.Println(msg)
		case <-timeout:
			fmt.Println("finish consume")
			close(t.Pipe)
		}

	}()

	time.Sleep(5 * time.Second)
}

// var test_func = TestFunc1()

func TestFunc1() {
	t := &TestConsumeA{
		Pipe:           make(chan int64),
		PipeNum:        10,
		ProduceDataNum: 1000,
	}
	var wg sync.WaitGroup
	go t.Producer(&wg)
}

func main() {
	TestFunc1()
}
