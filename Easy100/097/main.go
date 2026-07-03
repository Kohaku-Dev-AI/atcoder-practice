package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

type Shop struct{ Price, Stock int }

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	shops := make([]Shop, n)
	for i := 0; i < n; i++ {
		shops[i].Price = nextInt()
		shops[i].Stock = nextInt()
	}
	sort.Slice(shops, func(i, j int) bool {
		return shops[i].Price < shops[j].Price
	})

	var totalCost int64 = 0
	curretAmount := 0

	for i := 0; i < n; i++ {

		rem := m - curretAmount
		if rem == 0 {
			break
		}

		buy := min(shops[i].Stock, rem)
		totalCost += int64(shops[i].Price) * int64(buy)
		curretAmount += buy
	}

	fmt.Println(totalCost)
}
