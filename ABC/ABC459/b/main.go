package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(nextString())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()

	var result []byte

	for i := 0; i < n; i++ {
		s := nextString()
		if len(s) == 0 {
			continue
		}

		switch s[0] {
		case 'a', 'b', 'c':
			result = append(result, '2')
		case 'd', 'e', 'f':
			result = append(result, '3')
		case 'g', 'h', 'i':
			result = append(result, '4')
		case 'j', 'k', 'l':
			result = append(result, '5')
		case 'm', 'n', 'o':
			result = append(result, '6')
		case 'p', 'q', 'r', 's':
			result = append(result, '7')
		case 't', 'u', 'v':
			result = append(result, '8')
		case 'w', 'x', 'y', 'z':
			result = append(result, '9')
		}
	}

	fmt.Println(string(result))
}
