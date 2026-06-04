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

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	shiritori := make([]string, n)
	for i := 0; i < n; i++ {
		shiritori[i] = nextString()
	}
	runes := []rune(shiritori[0])
	prevLastWord := runes[len(shiritori[0])-1]
	for i := 1; i < n; i++ {
		runes = []rune(shiritori[i])
		currentfastWord := runes[0]
		if prevLastWord != currentfastWord {
			fmt.Println("No")
			return
		}
		for j := 0; j < i; j++ {
			if shiritori[i] == shiritori[j] {
				fmt.Println("No")
				return
			}
		}
		prevLastWord = runes[len(shiritori[i])-1]
	}
	fmt.Println("Yes")
}
