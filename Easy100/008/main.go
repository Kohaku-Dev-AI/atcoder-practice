package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	combinedStr := a + b
	ab, err := strconv.Atoi(combinedStr)
	if err != nil {
		fmt.Println("数値に変換できません")
		return
	}

	isSquare := false
	for i := 1; i*i <= ab; i++ {
		if ab == i*i {
			isSquare = true
			break
		}
	}

	if isSquare == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
