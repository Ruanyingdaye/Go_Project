package main

import (
	"fmt"
	"reflect"
	"sync"
)

type Config struct {
	data int
}

var configOnce sync.Once
var config Config //config一定只能全局set，因为once.Do只能initial,而不能直接返回一个内容，设置一个全局的，只要有一个协程初始化了之后其他的协程都可以拿到内容了，这才是once.Do的作用！

func initialConfig() {
	fmt.Println("only once")
	config.data = 200
}

func getconfig() int {
	// modify once use reflect?
	v := reflect.ValueOf(&configOnce).Elem() //Elem是调用元素的有效方式，后面可以用fieldByName来获取并修改
	// valueof后面一定要跟一个结构体指针，而不是结构体
	done := v.FieldByName("done")
	done = done.Elem().FieldByName("done")
	done.SetBool(false)
	configOnce.Do(initialConfig)
	return config.data
}

// 此处我们知道once就执行了一次值，如果我们修改某个变量的值用反射是否能执行多次呢？
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			c := getconfig()
			fmt.Println(c)
			wg.Done()
		}()
	}
	wg.Wait()
}

//////存在疑问，怎么修改？？？？？
