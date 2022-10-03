//go: build ignore

package main

import "fmt"

func main() {
	fmt.Println("===== 3.1 配列 ====")
	{
		var x [3]int
		fmt.Println(x)
	}
	{
		var x = [3]int{10, 20, 30}
		fmt.Println(x)
	}
	{
		var x = [12]int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	fmt.Println("---- 配列の比較 ----")
	{
		var x = [...]int{1, 2, 3}
		var y = [3]int{1, 2, 3}
		fmt.Println(x == y) //true が出力される。
	}

	fmt.Println("---- 多次元配列のシミュレート ----")
	{
		var x [2][3]int
		fmt.Println(x)
	}

	fmt.Println("---- インデックス（添え字）の指定 ----")
	{
		var x [3]int
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])

		var y [2][3]int
		y[0][2] = 12
		y[1][1] = 3
		fmt.Println("y:", y)
	}

	fmt.Println("---- len ----")
	{
		var x [3]int
		fmt.Println("len(x):", len(x))

		var y [2][3]int
		fmt.Println("len(y):", len(y))
	}

	fmt.Println("==== 3.2 スライス ====")
	{
		var x = []int{10, 20, 30}
		fmt.Println(x)
	}

	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println(x)
	}

	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])
	}

	{
		var x []int
		fmt.Println(x)
	}

	fmt.Println("---- スライスの比較 ----")
	{
		// スライスは比較可能では無い
		var x []int
		var y []int
		// fmt.Println(x == y) // これはエラー
		fmt.Println("x == nil:", x == nil) // true
		fmt.Println("y != nil:", y != nil) // false
	}

	fmt.Println("==== 3.2.1 len ====")
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println("len(x):", len(x))
		var y = []int{}
		fmt.Println("len{y}:", len(y))
	}

	fmt.Println("----- スライスのスライス -----")
	{
		var x [][]int
		fmt.Println(x)

		var y = [][]int{{0, 1}, {2, 3}, {4, 5}}
		fmt.Println(y)
		fmt.Println(y[1])    // [2 3]
		fmt.Println(y[1][1]) // 3
	}

	fmt.Println("----- lenの引数 -----")
	{
		var x [][]int
		fmt.Println("len(x):", len(x))
		var y = []int{2, 3, 4, 5}
		fmt.Println("len(y):", len(y))
		var z string = "abc"
		fmt.Println("len(z):", len(z))

		var a = 12
		fmt.Println("a:", a)
		// fmt.Println(len(a)) // エラー
	}

	fmt.Println("==== 3.2.2 append ====")
	{
		var x []int
		fmt.Println(x)
		x = append(x, 10)
		fmt.Println("append(x, 10):", x)
	}

	{
		var x = []int{1, 2, 3}
		x = append(x, 4)
		fmt.Println("append(x, 4):", x)
		x = append(x, 5, 6, 7)
		fmt.Println("append(x, 5, 6, 7):", x)

		y := []int{20, 30, 40}
		x = append(x, y...)
		fmt.Println("x:", x)
		fmt.Println("append(x, y...):", append(x, y...))
		fmt.Println("x:", x)
		// append(x, y) // エラー
	}

	fmt.Println("==== make ====")
	{
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
		fmt.Println("x[0], x[4]:", x[0], x[4])
	}

	{
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
		x = append(x, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}

	fmt.Println("----- キャパシティ（第3引数）を指定 -----")
	{
		x := make([]int, 5, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}

	{
		x := make([]int, 0, 10)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x)) // [] 0 10
	}

	{
		x := make([]int, 0, 10)
		x = append(x, 5, 6, 7, 8)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
	}

	fmt.Println("===== 3.2.5 スライスの生成方法の選択 =====")
	fmt.Println("----- 例3-2 -----")
	{
		var data []int // スライスのゼロ値nilで初期化される（nilスライス）
		fmt.Println(data == nil)

		x := []int{}          // 空のリテラルスライス
		fmt.Println(x == nil) // false
	}

	fmt.Println("----- 例3-3 -----")
	{
		// デフォルト値を指定したスライスの宣言
		data := []int{2, 4, 6, 8}
		fmt.Println(data)
	}

	fmt.Println("===== 3.2.6 スライスのスライス =====")

}
