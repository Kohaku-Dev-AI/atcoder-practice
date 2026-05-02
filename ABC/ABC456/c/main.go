package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod int64 = 998244353

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	const initialBufSize = 1024
	const maxBufSize = 1024 * 1024
	buf := make([]byte, initialBufSize)
	scanner.Buffer(buf, maxBufSize)

	scanner.Scan()
	s := scanner.Text()

	var ans int64
	var dpEnd int64 // i文字目で終わる有効な部分文字列の数

	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1] {
			// 直前と同じなら、長さ1の部分文字列しか作れない
			dpEnd = 1
		} else {
			// 直前と異なるなら、1つ前で終わる有効列をすべて1文字延長できる
			dpEnd = (dpEnd + 1) % mod
		}
		ans = (ans + dpEnd) % mod
	}

	fmt.Println(ans)
}