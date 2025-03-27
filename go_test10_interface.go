package main

import "fmt"

type baseStruct interface{}

type stack struct {
	ss []interface{} //interface需要这样define
}

func testAppend(t []int) {
	t = append(t, 100)
}
func testNums(t [3]int) {
	t[0] = 100
}
func testSlice(t []int) {
	t[0] = 100
}

type testFunc struct {
	A int
	M map[int]int
}

func modifyStructFunc(t testFunc) {
	t.A = 10
	// t.M[10] = 10
}

func (t testFunc) modifyStructMethod() {
	t.A = 10
	// t.M[10] = 10
}

func modifyStructFunc2(t *testFunc) {
	t.A = 10
	t.M[10] = 10
}

func (t *testFunc) modifyStructMethod2() {
	t.A = 10
	t.M[10] = 10 //这里的map一定要初始化，给一个nil的map直接赋值肯定会报错的！！！
}

func main() {
	fmt.Println("this is the stack initialization")

	mm := map[int]int{}
	mm2 := map[int]int{1: 2}
	mm3 := []int{}
	fmt.Println(mm, mm2, mm3)
	t := []int{1, 2, 3, 4, 5}
	testAppend(t)
	fmt.Println(t)
	//没错还是上次的问题，因为函数没有返回值，原来的那个slice进行了扩容，
	//新的地方的slice发生了变化，而原来的没有，我们调用的还是原来的slice
	//因此没发生变化，新地址的append后的内容没有返回
	//如果有append一定要记得返回，不然是不会修改slice的！！！！
	t2 := [3]int{1, 2, 3}
	testNums(t2)
	fmt.Println(t2)

	t3 := []int{1, 2, 3}
	testSlice(t3)
	fmt.Println(t3)

	t4 := testFunc{} //此时如果是结构体而不是结构体引用，那么传递可能就会有问题，但一般我们都是&xxx{}
	modifyStructFunc(t4)
	fmt.Println(t4.A, t4.M) //可以看到struct如果不加指针，返回的结果也不能进来
	t4.modifyStructMethod()
	fmt.Println(t4.A, t4.M) //可以看到struct作为接收器如果不加指针，返回的结果也不能进来

	t5 := &testFunc{} //此时如果是结构体而不是结构体引用，那么传递可能就会有问题，但一般我们都是&xxx{}
	t5.M = make(map[int]int)
	modifyStructFunc2(t5)
	fmt.Println(t5.A, t5.M) //可以看到struct如果不加指针，返回的结果也不能进来
	t5.modifyStructMethod2()
	fmt.Println(t5.A, t5.M) //可以看到struct作为接收器如果不加指针，返回的结果也不能进来

}
