package main

import (
	"fmt"
)

// 平均値 全ての値を足して、データの個数で割った値
// Goだと小数点を求めるのにfloat64型で計算する必要がありそう

func avg(nums []int) float64 {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return float64(sum) / float64(len(nums))
}

func main() {
	fmt.Println(avg([]int{1, 3, 10, 2, 8}))
}
