package raindrops

import "strconv"

func Convert(input int) string {
	str := ""
	if input%3 == 0 {
		str += "Pling"
	}

	if input%5 == 0 {
		str += "Plang"
	}

	if input%7 == 0 {
		str += "Plong"
	}

	if str == "" {
		str += strconv.Itoa(input)
	}

	return str
}
