package main

import "fmt"

func main() {
	// 配列の使い方
	var x = [3]int{10, 20, 30}
	x[0] = 10
	fmt.Println(x[2])
	// 範囲外を読んでいるのでエラーになる。
	fmt.Println(x[3])
}
