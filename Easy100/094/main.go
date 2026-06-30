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
	input := nextString()
	digits := make([]int, 4)
	for i, r := range input {
		digits[i] = int(r - '0')
	}

	if digits[0]+digits[1]+digits[2]+digits[3] == 7 {
		fmt.Printf("%d+%d+%d+%d=7\n", digits[0], digits[1], digits[2], digits[3])
		return
	}

	if digits[0]+digits[1]+digits[2]-digits[3] == 7 {
		fmt.Printf("%d+%d+%d-%d=7\n", digits[0], digits[1], digits[2], digits[3])
		return
	}

	if digits[0]+digits[1]-digits[2]+digits[3] == 7 {
		fmt.Printf("%d+%d-%d+%d=7\n", digits[0], digits[1], digits[2], digits[3])
		return
	}

	if digits[0]+digits[1]-digits[2]-digits[3] == 7 {
		fmt.Printf("%d+%d-%d-%d=7\n", digits[0], digits[1], digits[2], digits[3])
		return
	}

	if digits[0]-digits[1]+digits[2]+digits[3] == 7 {
		fmt.Printf("%d-%d+%d+%d=7\n", digits[0], digits[1], digits[2], digits[3])
		return
	}

	if digits[0]-digits[1]-digits[2]+digits[3] == 7 {
		fmt.Printf("%d-%d-%d+%d=7\n", digits[0], digits[1], digits[2], digits[3])
		return
	}

	if digits[0]-digits[1]+digits[2]-digits[3] == 7 {
		fmt.Printf("%d-%d+%d-%d=7\n", digits[0], digits[1], digits[2], digits[3])
		return
	}

	if digits[0]-digits[1]-digits[2]-digits[3] == 7 {
		fmt.Printf("%d-%d-%d-%d=7\n", digits[0], digits[1], digits[2], digits[3])
		return
	}
}
