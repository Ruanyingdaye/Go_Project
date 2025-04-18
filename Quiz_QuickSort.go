package main

import "fmt"

func QuickSort(arr []int) {
	if len(arr) == 0 || len(arr) == 1 {
		return
	}
	pivot := arr[0]
	//后面先动
	left, right := 0, len(arr)-1
	for left < right {
		for left < right && arr[right] > pivot {
			right--
		}
		arr[left] = arr[right]
		left++
		if left >= right {
			break
		}
		for left < right && arr[left] < pivot {
			left++
		}
		arr[right] = arr[left]
		right--
	}
	arr[left] = pivot
	// fmt.Println(arr, left, right)
	QuickSort(arr[:left])
	QuickSort(arr[left+1:])
	return
}

func main() {
	testArray := []int{3, 6, 3, 8, 10, 2, 5, 1, 8}
	QuickSort(testArray)
	fmt.Println(testArray)
}
