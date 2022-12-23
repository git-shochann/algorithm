package main

import "fmt"

// 並び替えが発生するとき最小単位で考える 一気に考えようとしない

func BubbleSort(nums []int) []int {

	lenNumbers := len(nums) // 5

	for i := 0; i < lenNumbers; i++ {
		// 前提: 配列は0, 1, 2, 3, 4と数える
		for j := 0; j < lenNumbers-1-i; j++ { // -1ずつずらしていけばOK 配列は0,1,2,3,4とアクセスすることに注意
			// 0番目 0+1番目で左と右を比較
			if nums[j] > nums[j+1] { // 手前の方が大きければ入れ替える
				nums[j], nums[j+1] = nums[j+1], nums[j] // 0の部分を0+1にする のと 0+1の部分を0にする こうすることで数字を一気に入れ替える
			}
		}

	}

	return nums
}

func main() {
	nums := []int{2, 6, 1, 9, 7} // 要素数は5
	fmt.Println(BubbleSort(nums))
}
