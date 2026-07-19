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
	h, w := nextInt(), nextInt()

	if w*10000 >= 25*h*h {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
