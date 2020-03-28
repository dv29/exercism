package matrix

import (
	"errors"
	// "fmt"
	"strconv"
	"strings"
)

type matrixType [][]int

type mtx interface {
	Rows() [][]int
	Cols() [][]int
}

func New(m string) (matrixType, error) {
	row := strings.Split(m, "\n")

	mx := make([][]int, len(row))
	for i, v := range row {
		col := strings.Split(v, " ")
		mx[i] = make([]int, len(col))

		for j, c := range col {
			val, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			mx[i][j] = val
		}
	}

	for x := 0; x < len(mx); x++ {
		if len(mx[0]) != len(mx[x]) {
			return nil, errors.New("mismatched length")
		}
	}

	return mx, nil
}

func (m matrixType) Rows() [][]int {
	y := make([][]int, len(m))
	for i := 0; i < len(m); i++ {
		y[i] = make([]int, len(m[0]))
		for j := 0; j < len(m[0]); j++ {
			y[i][j] = m[i][j]
		}
	}
	return y
}

func (m matrixType) Cols() [][]int {
	y := make([][]int, len(m[0]))

	for i := 0; i < len(m[0]); i++ {
		y[i] = make([]int, len(m))
	}

	for i := 0; i < len(y); i++ {
		for j := 0; j < len(y[0]); j++ {
			y[i][j] = m[j][i]
		}
	}
	return y
}

func (m *matrixType) Set(r, c, val int) bool {
	if r < 0 || r >= len((*m)[0]) || c < 0 || c >= len((*m)[0]) {
		return false
	}
	(*m)[r][c] = val
	return true
}
