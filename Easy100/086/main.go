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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	s := nextString()
	k := nextInt64()

	for i, c := range s {
		if c != '1' {
			if k <= int64(i) {
				break
			}
			fmt.Println(c - '0')
			return
		}
	}
	fmt.Println(1)
}
