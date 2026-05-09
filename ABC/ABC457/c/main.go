package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	nextInt := func() int {
		sc.Scan()
		i, _ := strconv.Atoi(sc.Text())
		return i
	}

	n := nextInt()
	k := nextInt()
	matrix := make([][]int, n)

	for i := 0; i < n; i++ {
		l := nextInt()
		matrix[i] = make([]int, l)
		for j := 0; j < l; j++ {
			matrix[i][j] = nextInt()
		}
	}

	c := make([]int, n)
	count := 0
	for i := 0; i < n; i++ {
		c[i] = nextInt()
		count += c[i] * len(matrix[i])

	}

	b := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := 0; j < c[i]; j++ {
			b = append(b, matrix[i]...)
		}
	}
	fmt.Println(b[k-1])
}
