package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	nextIntBig := func() (*big.Int, bool) {
		if !sc.Scan() {
			return nil, false
		}
		n := new(big.Int)
		return n.SetString(sc.Text(), 10)
	}

	a, okA := nextIntBig()
	b, okB := nextIntBig()

	if !okA || !okB {
		return
	}

	switch a.Cmp(b) {
	case 1:
		fmt.Println("GREATER")
	case -1:
		fmt.Println("LESS")
	default:
		fmt.Println("EQUAL")
	}
}
