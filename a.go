package main

import "fmt"

func bubblesort(nums []int) []int {
	// まずは繰り返す回数を決定 -> 配列の要素分
	for i := 0; i < len(nums); i++ {
		// jは動かす回数 配列は0からアクセスするので1を引く、そしてどんどん順に左に詰めるので-iをする
		for j := 0; j < len(nums)-1-i; j++ {
			// 配列の0番目と0番目+1を比較する
			if nums[j] > nums[j+1] {
				// 左の方が大きければ入れ替える 例えば 10 > 8
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// 問題: 昇順で並び替えて出力
// 前提: スライスは参照型
func main() {
	nums := []int{7, 2, 9, 3, 6, 7, 8} // 要素数 -> 7
	fmt.Println(bubblesort(nums))
}
