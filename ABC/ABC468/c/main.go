package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 階乗をあらかじめ計算しておく配列 (10! まで)
var fact [11]int64

func init() {
	fact[0] = 1
	for i := 1; i <= 10; i++ {
		fact[i] = fact[i-1] * int64(i)
	}
}

// 順列 P が辞書順で何番目か（0-indexed）を計算する関数
func permutationIndex(p []int) int64 {
	n := len(p)
	var index int64 = 0

	// 使用済みの数字を管理するフラグ（または配列）
	used := make([]bool, n+1)

	for i := 0; i < n; i++ {
		// p[i] より小さく、まだ使われていない数字が何個あるか数える
		smallerCount := 0
		for v := 1; v < p[i]; v++ {
			if !used[v] {
				smallerCount++
			}
		}

		// 残りの桁数分の組み合わせを足し合わせる
		remaining := n - 1 - i
		index += int64(smallerCount) * fact[remaining]

		// 使った数字をマーク
		used[p[i]] = true
	}

	return index
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	p := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		p[i], _ = strconv.Atoi(scanner.Text())
	}

	q := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		q[i], _ = strconv.Atoi(scanner.Text())
	}

	// index(Q) - index(P) - 1 が求める答え
	// （Q未満の個数から、P以下の個数を引くイメージ）
	ans := permutationIndex(q) - permutationIndex(p) - 1

	fmt.Println(ans)
}
