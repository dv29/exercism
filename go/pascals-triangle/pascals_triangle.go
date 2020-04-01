package pascal

import (
	"math"
)

func Triangle(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}

	t := make([][]int, n)

	for i := 1; i <= n; i++ {
		x := i - 1
		t[x] = make([]int, i)
		mid := int(math.Round(float64(i) / 2.0))

		t[x][0] = 1

		for j := 2; j <= mid; j++ {
			y := j - 1
			if x > 1 && y > 0 {
				t[x][y] = t[x-1][y-1] + t[x-1][y]
			} else {
				t[x][y] = j
			}
		}

		for j := mid; j < i; j++ {
			t[i-1][j] = t[i-1][i-j-1]
		}
	}

	return t
}
