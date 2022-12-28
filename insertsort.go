package main

import "fmt"

func insertsort(nums []int) []int {
	// 整列していない部分から挿入する値を取り出すループ i
	for i := 1; i < len(nums); i++ {
		temp := nums[i] // 整列していない動かしたい要素

		// 整列済みの全ての要素と比較するループ j
		// どんどん整列済みデータを右から左に見ていくため、indexがj--と減り、indexが0まで繰り返す
		// j := i - 1 は整列済みデータの一番右を示す
		j := i - 1
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j--
		}
		// ずらす処理が終わった位置に、挿入する値を入れる
		// なぜj+1？
		nums[j+1] = temp
	}
	return nums
}

func main() {
	nums := []int{7, 2, 8, 4, 0}
	fmt.Println(insertsort(nums))
}
