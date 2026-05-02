package main

import (
	"fmt"
)

func main() {
	list1 := make([]int, 6)
	for i := 0; i < 6; i++ {
		fmt.Scan(&list1[i])
	}
	list2 := make([]int, 6)
	for j := 0; j < 6; j++ {
		fmt.Scan(&list2[j])
	}
	list3 := make([]int, 6)
	for k := 0; k < 6; k++ {
		fmt.Scan(&list3[k])
	}

	sum := 0
	var ans float64
	for l := 0; l < 6; l++ {
		if list1[l] == 4 {
			for m := 0; m < 6; m++ {
				if list2[m] == 5 {
					for n := 0; n < 6; n++ {
						if list3[n] == 6 {
							sum++
						}
					}
				} else if list2[m] == 6 {
					for n := 0; n < 6; n++ {
						if list3[n] == 5 {
							sum++
						}
					}
				}
			}
		}
		if list1[l] == 5 {
			for m := 0; m < 6; m++ {
				if list2[m] == 4 {
					for n := 0; n < 6; n++ {
						if list3[n] == 6 {
							sum++
						}
					}
				} else if list2[m] == 6 {
					for n := 0; n < 6; n++ {
						if list3[n] == 4 {
							sum++
						}
					}
				}
			}
		}
		if list1[l] == 6 {
			for m := 0; m < 6; m++ {
				if list2[m] == 4 {
					for n := 0; n < 6; n++ {
						if list3[n] == 5 {
							sum++
						}
					}
				} else if list2[m] == 5 {
					for n := 0; n < 6; n++ {
						if list3[n] == 4 {
							sum++
						}
					}
				}
			}
		}
	}
	if sum > 0 {
		ans = float64(sum) / 216
	} else {
		ans = 0
	}
	fmt.Println(ans)
}
