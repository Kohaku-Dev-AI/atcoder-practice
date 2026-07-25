package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

// 💡 2つのスライスが完全一致しているか判定するヘルパー関数
func equal(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}

	q := make([]int, n)
	for i := 0; i < n; i++ {
		q[i] = nextInt()
	}

	// 1 から N までの初期順列を作る（例: [1, 2, 3]）
	cur := make([]int, n)
	for i := 0; i < n; i++ {
		cur[i] = i + 1
	}

	// P や Q が何番目に出てくるかを記録する変数
	pIdx, qIdx := -1, -1
	idx := 0

	// 順列を辞書順にすべて生成しながら順番をカウントする
	for {
		if equal(cur, p) {
			pIdx = idx
		}
		if equal(cur, q) {
			qIdx = idx
		}

		// 次の順列（辞書順で1つ後ろ）を作る。作れなくなったら終了
		if !nextPermutation(cur) {
			break
		}
		idx++
	}

	// P と Q の「間」にある個数を計算する
	// P より大きく Q より小さい ➔ |pIdx - qIdx| - 1
	ans := pIdx - qIdx
	if ans < 0 {
		ans = -ans
	}
	fmt.Println(ans - 1)
}

// 💡 辞書順で「次の順列」を生成する関数（next_permutation）
func nextPermutation(a []int) bool {
	n := len(a)
	i := n - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := n - 1
	for a[j] <= a[i] {
		j--
	}
	a[i], a[j] = a[j], a[i]
	for k, l := i+1, n-1; k < l; k, l = k+1, l-1 {
		a[k], a[l] = a[l], a[k]
	}
	return true
}
