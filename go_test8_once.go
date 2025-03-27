package main

import (
	"fmt"
	"sync"
)

func main() {
	//***********************************1 check once
	// singleton是单例对象，然后用sync.once来保证只执行一次，就是自己定义的一个结构体，用于实现singleton的
	var once sync.Once
	count := 0
	testOnceFunc := func() {
		count++
		fmt.Println(count)
	}
	once.Do(testOnceFunc)
	once.Do(testOnceFunc)
	once.Do(testOnceFunc) // we can call more than 1 time to call the func, but just run only once

	t1 := []int{1, 2, 3, 4, 5}
	t2 := t1
	t1[0] = 100
	fmt.Println(t1, t2) //引用类型的数据默认是浅拷贝，例如slice和map。slice的复制和传值都是浅拷贝，一改都修改了

	t2 = append(t1, 100)
	fmt.Println(t1, t2)

	tmpt2 := []int{3, 4, 5, 5, 6, 7}
	t2 = append(t2, tmpt2...)
	fmt.Println(t1, t2)

	//但是我们可以发现一旦用了append函数， cap不够用了就必然会开辟一块新的内存空间，和原来的空间不共享了，因此append了t2后，t1就不变化了
	//用t3进行下对比, 因为append会返回一个新的struct，虽然可能还是同样的底层数组，但明显不是一个struct了，因此append后肯定与原来的slice是不同的了~

	t3 := make([]int, 5, 10)
	t3[0] = 1
	t3[1] = 2
	t4 := t3
	fmt.Println(t3, t4)
	fmt.Printf("%p, %p\n", &t3, &t4)

	t3 = append(t3, 1)
	fmt.Println(t3, t4)
	fmt.Printf("%p, %p\n", &t3, &t4)

	t3 = append(t3, tmpt2...)
	fmt.Printf("%v, %p\n", t3, &t3)

	//地址本身是不会变的，因为存的是一个结构体，里面有底层数组的指针，这个指针可能会发生变化的,可以用reflect.ValueOf来实现

}
