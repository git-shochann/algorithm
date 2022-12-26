package main

import "fmt"

func selectionsort(nums []int) []int {
	// ソートする回数 = 要素の数-1
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i                        // まずは最初(0番目)のインデックス番号
		for j := i + 1; j < len(nums); j++ { // 右と比較するため+1で比較する
			if nums[j] < nums[minIndex] { // nums[1] < nums[0]であれば、入れ替える
				// 入れ替え作業
				nums[j], nums[minIndex] = nums[minIndex], nums[j]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{6, 5, 2, 9, 1, 4}
	fmt.Println(selectionsort(nums))
}
