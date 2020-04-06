package railfence

// Encode encodes the given string
func Encode(s string, i int) string {
	strs := make([]string, i)
	selector := 0
	swingType := 0

	for _, v := range s {
		strs[selector] += string(v)
		if selector == i-1 {
			swingType = 1
		} else if selector == 0 {
			swingType = 0
		}

		if swingType == 0 {
			selector++
		} else {
			selector--
		}
	}

	str := ""
	for _, v := range strs {
		str += v
	}

	return str
}

// Decode decodes the given string
func Decode(s string, rails int) string {
	strsLens := make([]int, rails)
	selector := 0
	swingType := 0

	for range s {
		strsLens[selector]++
		if selector == rails-1 {
			swingType = 1
		} else if selector == 0 {
			swingType = 0
		}

		if swingType == 0 {
			selector++
		} else {
			selector--
		}
	}

	strs := make([]string, rails)
	lastStop := 0
	for i := 0; i < rails; i++ {
		nextStop := lastStop + strsLens[i]
		strs[i] = s[lastStop:nextStop]
		lastStop = nextStop
	}

	selector = 0
	str := ""
	for range s {
		str += string(strs[selector][0])

		strs[selector] = strs[selector][1:]

		if selector == rails-1 {
			swingType = 1
		} else if selector == 0 {
			swingType = 0
		}

		if swingType == 0 {
			selector++
		} else {
			selector--
		}
	}

	return str
}
