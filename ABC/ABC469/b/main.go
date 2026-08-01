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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := nextString()
	totalCount := 0

	if n == 1 {
		if s[0] == 'x' {
			totalCount++
		}
		fmt.Println(totalCount)
		return
	}

	if n == 2 {
		if s[0] == 'x' && s[1] == 'x' {
			totalCount++
		}
		fmt.Println(totalCount)
		return
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			if s[i+1] == 'x' && s[0] == 'x' {
				totalCount++
			}
		} else if i == n-1 {
			if s[n-2] == 'x' && s[n-1] == 'x' {
				totalCount++
			}
		} else {
			if s[i-1] == 'x' && s[i+1] == 'x' && s[i] == 'x' {
				totalCount++
			}
		}
	}

	fmt.Println(totalCount)
}
