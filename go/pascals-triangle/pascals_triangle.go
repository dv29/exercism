package pascal

func Triangle(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}

	t := make([][]int, n)
	t[0] = []int{1}

	for i := 1; i < n; i++ {
		t[i] = make([]int, i+1)
		t[i][0], t[i][i] = 1, 1

		for j := 1; j < i/2+1; j++ {
			t[i][j] = t[i-1][j-1] + t[i-1][j]
			t[i][i-j] = t[i][j]
		}
	}

	return t
}
