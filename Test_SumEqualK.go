package main

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func subarraySum(nums []int, k int) int {
	resCount := 0
	tmpSum := 0
	sumMap := make(map[int]int)
	//sumMap表示key是前缀和，v表示出现次数
	sumMap[0] = 1
	for _, item := range nums {
		tmpSum += item
		if _, ok := sumMap[tmpSum]; ok {
			sumMap[tmpSum]++
			if cnt, ok := sumMap[tmpSum-k]; ok {
				resCount += cnt
			}
		} else {
			sumMap[tmpSum] = 1
		}
		fmt.Println(tmpSum, k, sumMap[tmpSum], sumMap[tmpSum-k], resCount)
	}
	return resCount
}

func testSubarraySumCase(t *testing.T) {
	nums := []int{1, 1, 1}
	res := subarraySum(nums, 2)
	//check subarrsum result
	assert.Equal(t, 2, res)
}

func main() {
	nums := []int{1, 1, 1}
	res := subarraySum(nums, 2)
	//check subarrsum result
	fmt.Println(res)
}
