package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	h := nextInt()
	w := nextInt()
	if h == 1 && w == 1 {
		fmt.Println(0)
		return
	}

	matrix := make([][]int, h)
	for i := range matrix {
		matrix[i] = make([]int, w)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if (h == 1 && j == 0) || (h == 1 && j == w-1) || (w == 1 && i == 0) || (w == 1 && i == h-1) {
				matrix[i][j] = 1
			} else if (h == 1 && j != 0) || (h == 1 && j != w-1) || (w == 1 && i != 0) || (w == 1 && i != h-1) {
				matrix[i][j] = 2
			} else if (i == 0 && j == 0) || (i == 0 && j == w-1) || (i == h-1 && j == 0) || (i == h-1 && j == w-1) {
				matrix[i][j] = 2
			} else if i == 0 || i == h-1 || j == 0 || j == w-1 {
				matrix[i][j] = 3
			} else {
				matrix[i][j] = 4
			}
		}
	}
	for i := 0; i < h; i++ {
		strRow := make([]string, w)
		for j := 0; j < w; j++ {
			strRow[j] = strconv.Itoa(matrix[i][j])
		}
		fmt.Println(strings.Join(strRow, " "))
	}
}
