package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()
	t := nextString()
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n; i++ {
		first := runes[0]
		runes = append(runes[1:], first)
		if string(runes) == t {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
