package main

import "fmt"

func insertsort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		// 0番目の要素と1番目の要素を比べる -> 0番目の方が大きければ、1番目と入れ替える
		for i >= 0 && nums[i-1] > temp {
			nums[i-1] = nums[i]
			i++
		}
	}
	return nums
}

func main() {
	nums := []int{7, 2, 9, 3, 0, 7, 8}
	fmt.Println(insertsort(nums))
}
