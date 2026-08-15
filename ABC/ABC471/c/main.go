package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt64() int64 {
	sc.Scan()
	i, _ := strconv.ParseInt(sc.Text(), 10, 64)
	return i
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	n := int(nextInt64())

	location := make([]int64, n)
	for i := 0; i < n; i++ {
		location[i] = nextInt64()
	}

	// 座標順にソート（例: [-10, -3, 2, 5, 8]）
	slices.Sort(location)

	// マイナス側で一番原点に近いインデックス (left)
	// プラス側で一番原点に近いインデックス (right)
	left := -1
	right := n

	for i := 0; i < n; i++ {
		if location[i] < 0 {
			left = i // 負の数のうち、最も右側（原点に近い）を保持
		}
	}
	for i := n - 1; i >= 0; i-- {
		if location[i] > 0 {
			right = i // 正の数のうち、最も左側（原点に近い）を保持
		}
	}

	var totalDistance int64 = 0
	var prevLocation int64 = 0

	// N個のクッキーをすべて拾うまでN回ループ
	for k := 0; k < n; k++ {
		distLeft := int64(1e18) // 十分大きい値で初期化
		distRight := int64(1e18)

		if left >= 0 {
			distLeft = abs(location[left] - prevLocation)
		}
		if right < n {
			distRight = abs(location[right] - prevLocation)
		}

		// 距離が等しい場合は「座標が小さい方＝left（マイナス側）」を優先する（問題文の指定）
		if distLeft <= distRight {
			totalDistance += distLeft
			prevLocation = location[left]
			left-- // 次に遠いマイナス側へ移動
		} else {
			totalDistance += distRight
			prevLocation = location[right]
			right++ // 次に遠いプラス側へ移動
		}
	}

	fmt.Println(totalDistance)
}
