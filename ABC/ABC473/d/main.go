package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	out = bufio.NewWriter(os.Stdout)
	n   int
	k   int
	a   []int // 現在作成中の数列 A
)

// DFS（深さ優先探索）で idx 番目の要素 (A[idx]) を決定する関数
func dfs(idx int, currentSum int) {
	// 【ベースケース】N 個すべての要素を決め切ったとき
	if idx > n {
		if currentSum == k {
			// 条件を満たしたら標準出力に一括出力
			for i := 1; i <= n; i++ {
				out.WriteString(strconv.Itoa(a[i]))
				if i < n {
					out.WriteString(" ")
				}
			}
			out.WriteString("\n")
		}
		return
	}

	// 【枝払い】現在の合計がすでに K を超えていたらこれ以上探索しない
	if currentSum > k {
		return
	}

	// idx 番目の要素 A[idx] として試せる値を 0 から順番に試す（辞書順になる）
	// (k - currentSum) / idx は A[idx] が取り得る理論上の最大値
	maxVal := (k - currentSum) / idx

	for v := 0; v <= maxVal; v++ {
		a[idx] = v
		// 次の要素 A[idx+1] を決めるために再帰呼び出し
		// 合計値には idx * v を加算する
		dfs(idx+1, currentSum+idx*v)
	}
}

func main() {
	// 出力の高速化（大量に出力するため bufio.Writer が必須）
	defer out.Flush()

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	if !sc.Scan() {
		return
	}
	n, _ = strconv.Atoi(sc.Text())
	sc.Scan()
	k, _ = strconv.Atoi(sc.Text())

	// 1-based index で扱いやすくするためサイズ N+1 で確保
	a = make([]int, n+1)

	// A[1] から探索スタート（初期の合計値は 0）
	dfs(1, 0)
}
