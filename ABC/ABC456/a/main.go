package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)

	if 3 <= x && x <= 18 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
