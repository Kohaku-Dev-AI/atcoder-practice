package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// これを呼ぶことで、スペース・改行をすべて「区切り」として扱う
	scanner.Split(bufio.ScanWords)

	// 1. まず N を取る
	if !scanner.Scan() {
		return
	}
	n, _ := strconv.Atoi(scanner.Text())

	pt, px, py := 0, 0, 0
	ans := "Yes"

	// 2. N回ループを回して、t, x, y を順番に吸い込む
	for i := 0; i < n; i++ {
		// 時刻 t
		scanner.Scan()
		nt, _ := strconv.Atoi(scanner.Text())
		// 座標 x
		scanner.Scan()
		nx, _ := strconv.Atoi(scanner.Text())
		// 座標 y
		scanner.Scan()
		ny, _ := strconv.Atoi(scanner.Text())

		dt := nt - pt
		dist := abs(nx-px) + abs(ny-py)

		// 判定
		if dt < dist || (dt-dist)%2 != 0 {
			ans = "No"
			// ここで break してもいいですが、
			// scanner.Scan を最後まで回さないと、入力がターミナルに残る場合があります。
			break
		}
		pt, px, py = nt, nx, ny
	}
	fmt.Println(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
