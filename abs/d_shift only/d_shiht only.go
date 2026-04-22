package main

import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	list := make([]int,n)
	for i:=0 ; i<n ; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Println(n)
	fmt.Println(list)
	count := 0
	for {
		isAllEven := true
		for _, a := range list {
			if a%2 != 0 {
				isAllEven = false
				break
			}
		}
		if isAllEven {
			for i := 0 ; i < n ; i++ {
				list[i] = list[i]/2
			} 
			count++
		} else {
			break
		}
		fmt.Println(list)
	}
	fmt.Println(count)
}