package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	a := nextInt()
	b := nextInt()
	m := nextInt()

	fridge := make([]int, a)
	for i := 0; i < a; i++ {
		fridge[i] = nextInt()
	}
	macrowave := make([]int, b)
	for i := 0; i < b; i++ {
		macrowave[i] = nextInt()
	}
	fridgeMin := slices.Min(fridge)
	macrowaveMin := slices.Min(macrowave)
	minPrice := fridgeMin + macrowaveMin
	for i := 0; i < m; i++ {
		x := nextInt()
		y := nextInt()
		c := nextInt()
		setPrice := fridge[x-1] + macrowave[y-1] - c
		if setPrice < minPrice {
			minPrice = setPrice
		}
	}
	fmt.Println(minPrice)
}
