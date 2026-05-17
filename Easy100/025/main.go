package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

// 2つのスライスが完全に一致するか判定するヘルパー関数
func isEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	// 順列 P の読み込み
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}

	// 順列 Q の読み込み
	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[i] = nextInt()
	}

	// 辞書順にすべての順列を生成するための変数
	var allPermutations [][]int
	used := make([]bool, n+1) // 1〜N の数字が使用中かどうかを記録するフラグ
	var current []int

	// 再帰関数による順列の全列挙（DFS）
	var dfs func()
	dfs = func() {
		// 現在の長さが N に達したら、1つの順列が完成
		if len(current) == n {
			// current の中身が変わっていくので、コピーをとって保存する
			temp := make([]int, n)
			copy(temp, current)
			allPermutations = append(allPermutations, temp)
			return
		}

		// 1 から N までの数字を辞書順に試していく
		for i := 1; i <= n; i++ {
			if !used[i] {
				used[i] = true
				current = append(current, i)

				dfs() // 次のマスを決めに行く（再帰）

				// 元に戻す（バックトラッキング）
				current = current[:len(current)-1]
				used[i] = false
			}
		}
	}

	// 順列生成スタート
	dfs()

	// P と Q がそれぞれ何番目（1始まり）にあるか探す
	a, b := 0, 0
	for idx, perm := range allPermutations {
		if isEqual(perm, p) {
			a = idx + 1
		}
		if isEqual(perm, q) {
			b = idx + 1
		}
	}

	// |a - b| の絶対値を計算して出力
	ans := int(math.Abs(float64(a - b)))
	fmt.Println(ans)
}