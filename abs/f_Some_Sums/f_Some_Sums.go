package main

import "fmt"

func main(){
	var n,a,b int 
	fmt.Scan(&n,&a,&b)
	sum:=0

	for i:=0 ; i<=n ; i++ {
		digit_sum := 0
		for j:=i ; j>0 ; j=j/10 {
			digit_sum = digit_sum+(j%10)
		}
		if digit_sum>=a && digit_sum<=b {
			sum = sum+i
		}
	}
	fmt.Println(sum)
}
