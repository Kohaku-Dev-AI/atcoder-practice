package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Scannerは1行が長すぎるとエラーになることがあるので、バッファサイズを拡張
	const maxCapacity = 100000 + 10 // 制約の10^5に対応
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	if !scanner.Scan() {
		return
	}
	s := scanner.Text() // ここで読み込んだ文字列を取得

	// HasSuffix（後ろからチェック）を使うなら、単語は「そのまま」でOK！
	t := []string{"dream", "dreamer", "erase", "eraser"}

	for len(s) > 0 {
		matched := false
		for _, word := range t {
			if strings.HasSuffix(s, word) {
				// 後ろから単語の長さ分を切り落とす
				s = s[:len(s)-len(word)]
				matched = true
				break
			}
		}
		if !matched {
			fmt.Println("NO") // 問題文に合わせて大文字
			return
		}
	}
	fmt.Println("YES")
}
