package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	// ここに処理を書く

	if a != b && b == c {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
