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
	n, _ := nextInt(), nextInt()
	memberCounts := make(map[int]int)
	for i := 0; i < n; i++ {
		temp := nextInt()
		memberCounts[temp]++
	}
	maxMembers := 0
	for _, memberCount := range memberCounts {
		if memberCount > maxMembers {
			maxMembers = memberCount
		}
	}
	classCount := 0
	for _, memberCounts := range memberCounts {
		if memberCounts == maxMembers || memberCounts == maxMembers-1 {
			classCount++
		}
	}
	fmt.Println(classCount)
}
