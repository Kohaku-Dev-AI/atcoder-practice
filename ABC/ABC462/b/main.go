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
	giftTo := make([][]int, n)
	for i := 0; i < n; i++ {
		k := nextInt()
		giftTo[i] = make([]int, k)
		for j := 0; j < k; j++ {
			giftTo[i][j] = nextInt()
		}
	}
	for i := 0; i < n; i++ {
		count := 0
		var nums []int
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(giftTo[j]); k++ {
				if giftTo[j][k] == i+1 {
					count++
					nums = append(nums, j+1)
					break
				}
			}
		}
		fmt.Print(count)
		for _, v := range nums {
			fmt.Printf(" %d", v)
		}
		fmt.Println()
	}
}
