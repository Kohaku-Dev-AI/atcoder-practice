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
	a, b := nextInt(), nextInt()

	ans := -1
	for price := 1; price <= 1300; price++ {
		tax8 := price * 8 / 100
		tax10 := price * 10 / 100

		if tax8 == a && tax10 == b {
			ans = price
			break
		}
	}
	fmt.Println(ans)
}
