package main

import "fmt"

// 最大値
// 最小値

func sum(nums []int) int {
	var max = nums[0]
	for i, _ := range nums {
		// 2 > 3
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func main() {
	fmt.Println(sum([]int{2, 4, 5, 7, 1}))
}
