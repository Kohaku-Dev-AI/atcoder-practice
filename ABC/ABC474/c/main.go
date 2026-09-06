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

func main() {
	sc.Split(bufio.ScanWords)
	n, q := nextInt(), nextInt()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}
	for i := 0; i < q; i++ {
		a := nextInt()
		for j := 0; j < n; j++ {
			if p[j] == a {
				p = append(p[:j], p[j+1:]...)
				p = append(p, a)
			}
		}
	}
	for i, v := range p {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}
