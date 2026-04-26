package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	list := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&list[i])
	}

	newlist := make([]int, 0)
	for j := 0; j < n; j++ {
		exists := false
		for k := 0; k < len(newlist); k++ {
			if newlist[k] == list[j] {
				exists = true
				break
			}
		}
		if !exists {
			newlist = append(newlist, list[j])
		}
	}
	fmt.Println(len(newlist))
}
