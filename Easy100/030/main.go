package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	x1,y1,x2,y2 := nextInt(),nextInt(),nextInt(),nextInt()
	var x3,y3,x4,y4 int int int int 
	if (y1 == y2){
		x3 = x2
		y3 = y1 + (x2 -x1)
		x4 = x1
		y4 = y3
	}	
}