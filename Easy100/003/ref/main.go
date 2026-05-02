package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)

	var s string
	fmt.Scan(&s)

	passedTotal := 0
	passedRankB := 0

	for i := 0; i < n; i++ {
		canPass := false

		switch s[i] {
		case 'a':
			if passedTotal < a+b {
				canPass = true
			}
		case 'b':
			if passedTotal < a+b && passedRankB < b {
				canPass = true
				passedRankB++
			}
		default:

		}
		if canPass {
			passedTotal++
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

	/*
		sumAB := 0
		sumB := 0
		for j := 0; j < n; j++ {
			switch s[j] {
			case 'a':
				if sumAB < a+b {
					sumAB++
					fmt.Println("Yes")
				} else {
					fmt.Println("No")
				}
			case 'b':
				if sumAB < a+b && sumB < b {
					sumAB++
					sumB++
					fmt.Println("Yes")
				} else {
					fmt.Println("No")
				}
			case 'c':
				fmt.Println("No")
			}
		}
	*/
}
