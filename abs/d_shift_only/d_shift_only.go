package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)

	list := make([]int,n)
	for i:=0 ; i<n ; i++ {
		fmt.Scan(&list[i])
	}

	count := 0
	for {
		for _, a := range list {
			if a%2 != 0 {
				fmt.Println(count)
				return
			}
		}

		for i := 0 ; i < n ; i ++ {
			list[i] /= 2
		}
		count ++
	}
}