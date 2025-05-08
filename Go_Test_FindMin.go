package main

import "fmt"

func getMin(start, end int, arr []int) (int, int, int) {
	// return start, end, MinIndex(default -1)

	mid := (start + end) / 2
	if start == mid || start == mid-1 {
		return start, end, mid
	}

	fmt.Println(arr[start], arr[mid], arr[end])

	if arr[start] >= arr[mid] {
		return getMin(start, mid, arr)
	} else if arr[end] <= arr[mid] {
		return getMin(mid, end, arr)
	}

	return start, end, -1

}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	start := 0
	end := len(arr) - 1
	_, _, Min := getMin(start, end, arr)
	fmt.Println(Min)
}
