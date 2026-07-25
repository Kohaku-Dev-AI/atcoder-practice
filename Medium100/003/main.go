package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	const maxBuf = 1000000
	buf := make([]byte, maxBuf)
	sc.Buffer(buf, maxBuf)

	if !sc.Scan() {
		return
	}
	s := sc.Text()
	n := len(s) + 1

	left := make([]int64, n)
	right := make([]int64, n)

	for i := 0; i < n-1; i++ {
		if s[i] == '<' {
			left[i+1] = left[i] + 1
		} else {
			left[i+1] = 0
		}
	}

	for i := n - 2; i >= 0; i-- {
		if s[i] == '>' {
			right[i] = right[i+1] + 1
		} else {
			right[i] = 0
		}
	}

	var ans int64 = 0
	for i := 0; i < n; i++ {
		if left[i] > right[i] {
			ans += left[i]
		} else {
			ans += right[i]
		}
	}

	fmt.Println(ans)
}
