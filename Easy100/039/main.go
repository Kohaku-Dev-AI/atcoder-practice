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
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	leftMax := make([]int, n)
	leftMax[0] = a[0]
	for i := 1; i < n; i++ {
		leftMax[i] = leftMax[i-1]
		if a[i] > leftMax[i] {
			leftMax[i] = a[i]
		}
	}
	rightMax := make([]int, n)
	rightMax[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = rightMax[i+1]
		if a[i] > rightMax[i] {
			rightMax[i] = a[i]
		}
	}
	for i := 0; i < n; i++ {
		ans := 0
		if i == 0 {
			ans = rightMax[1]
		} else if i == n-1 {
			ans = leftMax[n-2]
		} else {
			l := leftMax[i-1]
			r := rightMax[i+1]
			if l > r {
				ans = l
			} else {
				ans = r
			}
		}
		fmt.Fprintln(out, ans)
	}
}
