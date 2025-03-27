package main

import (
	"fmt"
)

func fastSort(array []int, start, end int) {
	//分治和递归
	if len(array) != 0 {
		i, j := start+1, end
		for i < j {
			for array[i] < array[start] && i < j {
				i++
			}
			array[start] = array[i]
			j--
			if i >= j {
				break
			}

			for array[j] > array[start] && i < j {
				j--
			}
			i++
		}
		array[i] = array[start]
		fastSort(array, start, i-1)
		fastSort(array, i+1, end)
	}
}

func main() {
	mms := []int{2, 30, 6, 7, 3, 6, 8, 2, 2, 34, 65, 5}
	fastSort(mms, 0, len(mms)-1)
	fmt.Println(mms)
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Printf("catch panic: %v", err) //recover()只有一个返回值, 用defer就可以，因为panic会导致程序退出
	// 	}
	// }()

	// var a [8]int
	// b := [8]int{1, 2, 3}
	// c := [...]int{1, 2, 3, 4}

	// d := []int{}
	// e := []int{1, 2, 3, 4}

	// aa := &b
	// ab := unsafe.Pointer(&a)
	// ac := (*int8)(ab)
	// ad := unsafe.Pointer(aa)
	// ae := (*int8)(ad)

	// fmt.Println(c, d, e, aa)
	// fmt.Println(ab, unsafe.Sizeof(ab), a)
	// fmt.Println(ac, unsafe.Sizeof(ac), *ac)      //这个8表示的是len，因为看起来只有len, 第一个元素是0
	// fmt.Println(ad, unsafe.Sizeof(ad), *aa, *ae) //此时1表示的就是第一个元素，此时我们将指针用uintPtr转换下，得到第二个元素, 因为指针指向的就是默认第一个元素
	// newP := unsafe.Pointer(uintptr(ad) + 16)
	// fmt.Println(newP, *(*int16)(newP)) //可以用这个算结构体的unsafe.Offsetof(aa)
	// fmt.Println("**********************")
	// map2 := make(map[int][]int)
	// tmp := []int{1, 2, 3}
	// map2[2] = tmp
	// fmt.Println(map2) //这是ok的

	// map3 := make(map[int]map[int]int)
	// tmp3 := make(map[int]int)
	// tmp3[3] = 3
	// map3[2] = tmp3
	// fmt.Println(map3) //这是ok的

	// map4 := make(map[int]map[int]int)
	// map4[2][3] = 1
	// fmt.Println(map4) //will panic: assignment to entry in nil map

	// map2[2][2] = 1
	// fmt.Println(map2) //可以

}
