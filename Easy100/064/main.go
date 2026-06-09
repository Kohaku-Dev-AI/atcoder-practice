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
	o := nextString()
	e := nextString()
	result := make([]rune, len(o)+len(e))

	for i, r := range o {
		result[i*2] = r
	}

	for i, r := range e {
		result[i*2+1] = r
	}
	fmt.Println(string(result))
}
