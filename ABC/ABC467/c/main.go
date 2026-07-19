package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt64()
	m := nextInt64()

	a := make([]int64, n)
	for i := int64(0); i < n; i++ {
		a[i] = nextInt64()
	}
	b := make([]int64, n-1)
	for i := int64(0); i < n-1; i++ {
		b[i] = nextInt64()
	}

	var count int64 = 0

	for i := int64(0); i < n-1; i++ {
		currentMod := (a[i] + a[i+1]) % m

		if currentMod != b[i] {
			a[i+1]++
			count++
		}
	}

	fmt.Println(count)
}
