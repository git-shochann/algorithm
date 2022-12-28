package main

import "fmt"

// aとbの値を入れ変える
func main() {
	a := 10
	b := 20
	a, b = b, a
	fmt.Println(a, b)

	c := 30
	d := 40
	fmt.Println(c, d)

	t := c
	c = d
	d = t
	fmt.Println(c, d)
}
