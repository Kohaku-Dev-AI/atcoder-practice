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
	n, k := nextInt(), nextInt()

	ans := k
	for i := 2; i <= n; i++ {
		ans = ans * (k-1)
	}
	fmt.Println(ans)
}
