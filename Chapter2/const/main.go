// 定数
package main

import "fmt"

// Example 2-3
const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"
	fmt.Println(x)
	fmt.Println(y)
	// 以下2行はエラーになる。
	x = x + 1
	y = "bye"

	fmt.Println(x)
	fmt.Println(y)
}
