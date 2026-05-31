package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	topAndBottomBorder := strings.Repeat("#", w+2)
	fmt.Println(topAndBottomBorder)

	for i := 0; i < h; i++ {
		row := nextString()
		fmt.Printf("#%s#\n", row)
	}

	fmt.Println(topAndBottomBorder)
}
