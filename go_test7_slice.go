package main

import "fmt"

func main() {

	//test yihuo
	flag := 0
	for i := 0; i < 10; i++ {
		flag = flag ^ 1
		fmt.Println("flag:", flag)
	}

	// go func() {
	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			fmt.Printf("catch panic: %v", err) //recover()只有一个返回值, 用defer就可以，因为panic会导致程序退出
	// 		}
	// 	}()
	// 	ch := make(chan int)
	// 	ch <- 1
	// }()

	// //test kuaipai

	// testnums := []int{7, 3, 5, 2, 1, 4, 8}
	// quickSort(0, 6, testnums)
	// fmt.Println(testnums)
	// fmt.Println("****************************************************************************")
	// a := [...]int{1, 2, 3, 4, 5}
	// testModifyNum(a) //don't modify the array value
	// fmt.Println(a)

	// b := 1
	// c := b
	// fmt.Printf("%v, %v\n", &b, &c)

	// aa := []int{1, 2, 3, 4, 5}
	// bb := []int{11, 22, 33}
	// cc := append(aa, bb...)
	// fmt.Println(aa, bb, cc)

	// dd := append(aa[0:2], bb...)
	// // aa变成了以该位置后面appendbb[0:1]的内容
	// // 如果bb数量少，不需要扩容，则aa就在原基础上就修改，这样的append会导致aa也被修改了；
	// // 但如果aa的cap不够大，就会在原来基础上进行扩容，所有的修改都是新的内存去操作，则变化不到a了
	// fmt.Println(aa, bb, dd)

	// testnums = append(testnums[3:], testnums[:3]...)
	// fmt.Println(testnums)

	// fmt.Println(max(b, b+1), min(b, b+1)) //原来golang有默认的max和min函数

	// tss := make([][]int, 5)
	// fmt.Println(tss)

}

func quickSort(start, end int, nums []int) {
	//明确输入输出参数
	//因为传的是slice，因此直接就可以影响原nums
	if len(nums) == 0 || start >= end {
		return
	}
	//check input边界条件

	//判断条件
	i := start
	j := end
	flagnum := nums[i]

	for i < j {
		for nums[j] >= flagnum && i < j {
			fmt.Println(j, "j--")
			j--
		}
		nums[i] = nums[j]

		for nums[i] <= flagnum && i < j {
			fmt.Println(i, "i++")
			i++
		}
		nums[j] = nums[i]

	}
	nums[i] = flagnum
	fmt.Println(nums, i, start, end)

	//调用递归需要判断
	quickSort(start, i-1, nums)
	quickSort(i+1, end, nums)

}

func testModifyNum(nums [5]int) {
	nums[0] = 1000
}
