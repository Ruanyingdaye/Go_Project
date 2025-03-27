package main

import (
	"fmt"
)

//  tell us the closure, just cap the free variable

func main() {
	ss := 100
	testC := testCount()
	testC(ss)
	testC(ss)
	testC(ss)

	// testD := testCount()
	// testD(200)
	// testD(200)
	// testD(200)

	tm := make(map[int]int)
	tm[1] = 1
	tm[2] = 2
	tm[3] = 3
	tm[4] = 4
	tm[5] = 5
	for _, item := range tm {
		fmt.Println(item)
	}
	rs := f3()
	fmt.Println("*****************")
	fmt.Println(rs)

	fmt.Println("*****************")
	rs2 := f2()
	fmt.Println(rs2)

	fmt.Println("*****************")

	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	fmt.Printf("%p, %p\n", &a, &b)
	fmt.Println(a, b)
	b = a
	c := b
	fmt.Printf("%p, %p, %p\n", &a, &b, &c) //两个指针是不一样的，只是指向同一个地址，因为赋值相等的是两个引用而不是地址本身啊
	// 因为map是引用，输出的也是引用地址，如果copy的话，引用也不会一样
	fmt.Println(a, b, c)
	c = append(c, 1)
	fmt.Println(a, b, c)

	fmt.Printf("%p, %p, %p\n", &a, &b, &c)
	bs := &a
	fmt.Printf("%p, %p, %p\n", &a, bs, *bs, &c)
	//指针指向的内容
	// ########################################################
	fmt.Println("\n\n*****************")
	ts := make([]int, 10)
	ts = append(ts, 1, 2, 3, 4, 5)
	// cc := append(ts, 1, 2, 3, 4, 5, 6)
	// fmt.Println(ts, cc)
	ts2 := make([]int, 20)
	ts2 = append(ts2, ts...)
	fmt.Println(ts2)
	ts3 := make([]int, 40)
	copy(ts3, ts2)
	ts4 := make([]int, 30)
	copy(ts4, ts2)
	fmt.Println("ts4", ts4)
	//copy不会自动扩容，手动copy能够更好的支持手动扩容, 用法是将后面的复制给前面的
	//因为copy不会自动扩容，因此copy不过来的就自动截断了！！！！！
	//因此一定要保证空间足够充足
	fmt.Printf("ts2 %p, %v, ts3 %p, %v", &ts2, ts2, &ts3, ts3)
	fmt.Println(len(ts2), len(ts3))
}

func f3() (r int) {
	defer func() {
		r = r + 5
	}()
	return r
}

func f2() (r int) {
	r = 2
	defer func(r int) { // defer在执行的时候，传入的参数r=0
		r = r + 5
		fmt.Println(r)
	}(r)
	// 传入的参数时0， 先确定defer func的参数，然后set return的函数，然后时执行defer，最后返回return的值
	return 1
}

func testCount() func(a int) {
	b := 100
	return func(a int) {
		// a++
		// fmt.Println(a)
		b++
		fmt.Println(b)
	}
}
