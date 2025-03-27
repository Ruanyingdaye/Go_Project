package main

import "fmt"

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	pivot := arr[0]
	var left, right []int
	for _, num := range arr[1:] {
		if num < pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}
	quickSort(left)
	quickSort(right)
	arr[0] = left[0]
	arr[len(arr)-1] = right[len(right)-1]
	quickSort(arr[1 : len(arr)-1])
}

// test how to use slice append
func appendTest(t []int) {
	// append the int
	t = append(t, 1)
	// append the slice
	t = append(t, []int{2, 3}...)
}

//因此append的方式最好有返回值，不然只要扩容了就没有办法将修改传上去了

func modifySliceTest(t []int) {
	t[0] = 1
}

func main() {
	// quickSort([]int{3, 2, 1, 4, 6, 5})
	stest := []int{0, 0, 1, 1, 2}
	appendTest(stest)
	fmt.Println(stest)
	//
	modifySliceTest(stest)
	fmt.Println(stest)
	// we see, modify can work, but append can't
	fmt.Println("new start")

}
