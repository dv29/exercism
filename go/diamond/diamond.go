package diamond

import (
	"errors"
)

func initEmptyString(v int) []byte {
	str := make([]byte, v)
	for i := 0; i < v; i++ {
		str[i] = byte(' ')
	}
	return str
}

// Gen generates the diamond string
func Gen(char byte) (string, error) {
	if int(char) > 90 || int(char) < 65 {
		return "", errors.New("invalid character")
	}

	if int(char)-65 == 0 {
		return "A\n", nil
	}
	l := ((int(char) - 64) * 2) - 1

	strs := make([]string, l)

	mid := int(l / 2)
	end := mid * 2
	front := 0

	for i := int(l / 2); i >= 0; i-- {
		str := initEmptyString(l)
		c := byte(i + 65)

		str[end] = c
		str[front] = c
		end--
		front++
		strs[i] = string(str)
	}

	for i := l - 1; i > int(char)-65; i-- {
		strs[i] = strs[l-1-i]
	}

	s := ""
	for _, x := range strs {
		s += x + "\n"
	}
	return s, nil
}
