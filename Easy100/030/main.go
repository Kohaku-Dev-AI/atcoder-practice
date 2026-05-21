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
	x1, y1, x2, y2 := nextInt(), nextInt(), nextInt(), nextInt()
	dx := x2 - x1
	dy := y2 - y1
	x3 := x2 - dy // 点2から「左（マイナス方向）」へ
	y3 := y2 + dx

	x4 := x1 - dy // 点1から「左（マイナス方向）」へ
	y4 := y1 + dx
	fmt.Println(x3, y3, x4, y4)
}
