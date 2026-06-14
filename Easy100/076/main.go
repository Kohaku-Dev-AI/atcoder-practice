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

type ProblemState struct {
	isAC    bool
	waCount int
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	problems := make([]ProblemState, n+1)

	totalAC := 0
	totalPenalty := 0

	for i := 0; i < m; i++ {
		p := nextInt()
		result := nextString()

		if problems[p].isAC {
			continue
		}

		if result == "AC" {
			problems[p].isAC = true
			totalAC++
			totalPenalty += problems[p].waCount
		} else {
			problems[p].waCount++
		}
	}
	fmt.Printf("%d %d\n", totalAC, totalPenalty)
}
