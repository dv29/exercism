package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("invalid lengths")
	}

  distance := 0
  for i, v := range a {
    if v != rune(b[i]) {
      distance++
    }
  }
  return distance, nil
}
