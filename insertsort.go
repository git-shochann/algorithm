package main

import "fmt"

func insertsort(nums []int) []int {
	// 整列していない部分から挿入する値を取り出すループ i
	for i := 1; i < len(nums); i++ {
		temp := nums[i] // 整列していない最初の数値

		// 整列済みと比較するループ j
		// どんどん整列済みデータを右から左に見ていくため、indexがj--と減り、indexが0まで見ていく
		// j := i - 1 は整列済みデータの一番右を示す
		var j int
		for j = i - 1; j >= 0; i-- {
			// 例えば7>2であれば
			// 例えば2>8であれば入れ替えない？
			if nums[j] > temp {
				// ７を右にずらす
				nums[j+1] = nums[j]
			} else {
				break
			}
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
