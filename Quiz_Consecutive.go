package main

import (
	"fmt"
)

func findSubarrays(nums []int, target int) interface{} {
	const maxResult = 1000000000
	var result [][]int
	count := 0
	prefixSum := make(map[int][]int)
	prefixSum[0] = []int{-1} // 初始前缀和为0出现在索引-1处
	currentSum := 0

	for i := 0; i < len(nums); i++ {
		currentSum += nums[i]

		// 查找所有使prefixSum[j] = currentSum - target的j
		if indices, exists := prefixSum[currentSum-target]; exists {
			for _, j := range indices {
				if count >= maxResult {
					return -1
				}
				count++
				// 仅当数组较小时存储具体结果
				if len(nums) <= 10000 {
					result = append(result, []int{j + 1, i})
				}
			}
		}

		// 记录当前前缀和的位置
		prefixSum[currentSum] = append(prefixSum[currentSum], i)
	}

	// 处理大数据量情况
	if len(nums) > 10000 {
		if count > maxResult {
			return -1
		}
		return count
	}

	return result
}

func main() {
	// 测试用例1: [2,-2,3,0,4,-7], target=0
	A1 := []int{2, -2, 3, 0, 4, -7}
	target1 := 0
	fmt.Println("Test case 1:")
	printResult(findSubarrays(A1, target1), A1)

	// 测试用例2: 100,000个零
	A2 := make([]int, 100000) // 100,000个零
	target2 := 0
	fmt.Println("\nTest case 2 (100,000 zeros):")
	printResult(findSubarrays(A2, target2), A2)

	// 测试用例3: 超过限制的情况
	A3 := make([]int, 1000000001) // 1,000,000,001个零
	target3 := 0
	fmt.Println("\nTest case 3 (1,000,000,001 zeros):")
	printResult(findSubarrays(A3, target3), A3)
}

func printResult(res interface{}, nums []int) {
	switch v := res.(type) {
	case int:
		fmt.Println(v)
	case [][]int:
		fmt.Printf("Found %d subarrays:\n", len(v))
		for _, pair := range v {
			fmt.Printf("(%d, %d) -> %v\n", pair[0], pair[1], nums[pair[0]:pair[1]+1])
		}
	default:
		fmt.Printf("Count: %v\n", v)
	}
}
