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
	n := nextInt()
	a := make([]int, n)
	checkFlg := false
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		if a[i] == 2 {
			checkFlg = true
		}
	}

	if checkFlg == false {
		fmt.Println(-1)
		return
	}
	count := 0
	for i := 0; ; {
		count++
		if a[i] == 2 {
			fmt.Println(count)
			return
		}
		i = a[i] - 1
		if count > n {
			fmt.Println(-1)
			return
		}
	}
}
