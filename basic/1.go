package main

// P165 確認問題1

import "fmt"

// x < y になればOK

var x = 5
var y = 3

// 5 / 3 -> 2 であまりは1なのでrは2, qは1になる

var q = 0
var r = x

func main() {

	for {
		if r < y {
			return
		} else {
			r = r - y
			q = q + 1
			fmt.Println(r) // 2
			fmt.Println(q) // 1
		}
	}
}
