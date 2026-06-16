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
	n := nextInt()
	if n == 1 {
		fmt.Println("Yes")
		return
	}
	nonDecreasingFlag := true
	prevValue := nextInt()
	for i := 1; i < n; i++ {
		currentValue := nextInt()
		if !nonDecreasingFlag {
			continue
		}

		if currentValue > prevValue {
			currentValue--
		} else if currentValue < prevValue {
			nonDecreasingFlag = false
			continue
		}
		prevValue = currentValue
	}
	if nonDecreasingFlag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
