package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	n, d := nextInt(), nextInt()
	x := make(map[int]int)
	location := make([]int, n)
	for i := 0; i < n; i++ {
		location[i] = nextInt()
		x[location[i]]++
	}

	coords := make([]int, 0, len(x))
	for pos := range x {
		coords = append(coords, pos)
	}
	slices.Sort(coords)

	isOK := make(map[int]bool)

	for i, pos := range coords {
		if x[pos] > 1 {
			continue
		}

		valid := true

		if i > 0 {
			leftPos := coords[i-1]
			if pos-leftPos < d {
				valid = false
			}
		}

		if i < len(coords)-1 {
			rightPos := coords[i+1]
			if rightPos-pos < d {
				valid = false
			}
		}

		if valid {
			isOK[pos] = true
		}
	}

	var ans []int
	for i := 0; i < n; i++ {
		pos := location[i]
		if isOK[pos] {
			ans = append(ans, i+1) 
		}
	}
	fmt.Println(len(ans))

	if len(ans) == 0 {
		fmt.Println()
	} else {
		for i, id := range ans {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(id)
		}
		fmt.Println()
	}
}