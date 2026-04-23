package main

import "fmt"

func main(){
	var a,b,c,x int
	fmt.Scan(&a,&b,&c,&x)

	count := 0
	// 500円玉の枚数
	for i := 0 ; i <= a ; i ++ {
			// 100円玉の枚数
			for j := 0 ; j <= b ; j ++ {
					// 50円玉の枚数
					for k := 0 ; k <= c ; k++ {
						if 500*i + 100*j + 50*k == x {
							count ++
						} 
					}
				} 
			}
	fmt.Println(count)
}
