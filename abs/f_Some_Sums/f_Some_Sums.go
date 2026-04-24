package main

import "fmt"

func findSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func main(){
	var n,a,b int 
	fmt.Scan(&n,&a,&b)

	totalSum := 0
	for i := 1 ; i <= n ; i++ {
		dSum := findSumOfDigits(i)

		if a <= dSum && dSum <= b {
			totalSum += i
		}
	}
	fmt.Println(totalSum)

	// sum:=0

	// for i:=0 ; i<=n ; i++ {
	// 	digit_sum := 0
	// 	for j:=i ; j>0 ; j=j/10 {
	// 		digit_sum = digit_sum+(j%10)
	// 	}
	// 	if digit_sum>=a && digit_sum<=b {
	// 		sum = sum+i
	// 	}
	// }
	// fmt.Println(sum)
}
