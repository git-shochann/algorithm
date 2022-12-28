package main

import "fmt"

// 合計値 配列に入っている値を全て足す 1つ1つ順番に足していく

func sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum([]int{2, 4, 5, 7, 1}))
}
