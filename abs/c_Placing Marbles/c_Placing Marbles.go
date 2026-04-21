package main

import "fmt"

func main(){
	var s string
	fmt.Scan(&s)

	count := 0

	for _, char := range s {
		if char == '1' {
			count ++
		}
	}
	fmt.Println(count)
}