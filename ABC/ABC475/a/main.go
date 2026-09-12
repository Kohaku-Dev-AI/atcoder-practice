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
	s := nextString()
	result := ""
	for i, v := range s {
		if i == len(s)-1 {
			result = result + string(v)
			break
		}
		result = result + string(v) + "o"
	}
	fmt.Println(result)
}
