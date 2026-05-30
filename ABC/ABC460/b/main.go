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
	t := nextInt()
	circle := make([][]int, t)
	for i := 0; i < t; i++ {
		circle[i] = make([]int, 6)
		for j := 0; j < 6; j++ {
			circle[i][j] = nextInt()
		}
	}
	for i := 0; i < t; i++ {
		if (circle[i][2]-circle[i][5])*(circle[i][2]-circle[i][5]) <= (circle[i][3]-circle[i][0])*(circle[i][3]-circle[i][0])+(circle[i][4]-circle[i][1])*(circle[i][4]-circle[i][1]) && (circle[i][3]-circle[i][0])*(circle[i][3]-circle[i][0])+(circle[i][4]-circle[i][1])*(circle[i][4]-circle[i][1]) <= (circle[i][2]+circle[i][5])*(circle[i][2]+circle[i][5]) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
