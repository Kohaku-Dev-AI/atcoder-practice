package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(reader, &s)

	stack := []rune{}
	ans := 0

	for _, char := range s {
		if len(stack) > 0 && stack[len(stack)-1] != char {
			stack = stack[:len(stack)-1]
			ans += 2
		} else {
			stack = append(stack, char)
		}
	}

	fmt.Println(ans)
}
