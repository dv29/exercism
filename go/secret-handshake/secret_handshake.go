package secret

func Handshake(code uint) []string {
	var s []string

	if (code & 1) == 1 {
		s = append(s, "wink")
	}

	if (code & 2) == 2 {
		s = append(s, "double blink")
	}

	if (code & 4) == 4 {
		s = append(s, "close your eyes")
	}

	if (code & 8) == 8 {
		s = append(s, "jump")
	}

	if (code & 16) == 16 {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	return s
}
