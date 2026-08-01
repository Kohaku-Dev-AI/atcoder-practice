package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	var s string
	fmt.Fscan(in, &n, &s)

	// 1. 各位置までの 'o' の累計数 (累積和) を計算する
	// C[i] : 先頭 i 文字に含まれる 'o' の数
	C := make([]int, n+1)
	// 2. 'x' が現れるインデックス (1-indexed) を記録する配列
	xPos := make([]int, 0, n)

	for i := 0; i < n; i++ {
		C[i+1] = C[i]
		if s[i] == 'o' {
			C[i+1]++
		} else {
			xPos = append(xPos, i+1) // 1-based の位置を記録
		}
	}

	// s[i] 以降で何番目の 'x' なのかを素早く引くために
	// 各位置 i から見て「それまでに 'x' が何個あったか」をカウント
	xCountBefore := make([]int, n+1)
	xCnt := 0
	for i := 0; i < n; i++ {
		if s[i] == 'x' {
			xCnt++
		}
		xCountBefore[i+1] = xCnt
	}

	// 3. 各 k (1 ～ N) について答えを求める
	for k := 1; k <= n; k++ {
		// 手元のライフ (先頭 k 個に含まれる 'o' の数)
		life := C[k]

		// k 個目を取った時点で、既に通過した 'x' の個数
		alreadyX := xCountBefore[k]

		// 追加で踏める 'x' の限界インデックス
		// これまでに出た 'x' の数 + 手元のライフ (life)
		targetXIndex := alreadyX + life

		var ans int
		if targetXIndex >= len(xPos) {
			// 踏める 'x' の上限が、全体にある 'x' の数以上なら
			// 列の最後まで全完走できる！
			ans = n
		} else {
			// targetXIndex 番目の 'x' (0-indexed) の位置まで進める
			// (その 'x' を踏んだ瞬間にお菓子を食べて終了)
			ans = xPos[targetXIndex]
		}

		fmt.Fprintln(out, ans)
	}
}