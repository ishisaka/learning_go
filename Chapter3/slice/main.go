package main

import "fmt"

func main() {
	// スライスでは配列のサイズを指定しない。配列とよく似ている。
	var x = []int{10, 20, 30}
	x[0] = 10
	fmt.Println(x[2])
	// 以下は範囲外なので実行時エラーになる
	// fmt.Println(x[3])

	// 以下Trueになる
	var y []int
	fmt.Println(y == nil)

	// len
	fmt.Println(len(x), len(y))

	// append
	x = append(x, 10)
	fmt.Println(x, len(x))
	y = append(y, 100, 200)
	// sliceの合成は以下のように行う
	x = append(x, y...)
	fmt.Println(x, len(x))
}
