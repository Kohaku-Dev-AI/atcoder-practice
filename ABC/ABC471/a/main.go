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
	a, b := nextInt(), nextInt()
	if a+b == 9 || a-b == 9 || a*b == 9 || (a/b == 9 && a%b == 0) {
		fmt.Println("Nine")
		return
	}
	fmt.Println("Nein")
}
