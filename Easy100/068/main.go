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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

type Restrant struct {
	id    int
	city  string
	point int
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	restrants := make([]Restrant, n)
	for i := 0; i < n; i++ {
		restrants[i].city = nextString()
		restrants[i].point = nextInt()
		restrants[i].id = i + 1
	}

	sort.Slice(restrants, func(i, j int) bool {
		if restrants[i].city != restrants[j].city {
			return restrants[i].city < restrants[j].city
		}
		return restrants[i].point > restrants[j].point
	})

	for _, r := range restrants {
		fmt.Println(r.id)
	}
}
