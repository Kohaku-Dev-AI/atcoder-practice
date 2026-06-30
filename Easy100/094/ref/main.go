package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

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

	for i := 0; i < (1 << 3); i++ {
		sum := digits[0]
		op := make([]string, 3)

		for j := 0; j < 3; j++ {
			if (i >> j & 1) == 1 {
				sum -= digits[j+1]
				op[j] = "-"
			} else {
				sum += digits[j+1]
				op[j] = "+"
			}
		}

		if sum == 7 {
			fmt.Printf("%d%s%d%s%d%s%d=7\n", digits[0], op[0], digits[1], op[1], digits[2], op[2], digits[3])
			return
		}
	}
}
