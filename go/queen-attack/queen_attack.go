package queenattack

import (
	"errors"
)

func CanQueenAttack(w, b string) (bool, error) {
	if w == "" || b == "" || len(w) > 2 || len(b) > 2 {
		return false, errors.New("invalid input")
	}

	wRow := int(w[0]) - 96
	wCol := int(w[1]) - 48

	bRow := int(b[0]) - 96
	bCol := int(b[1]) - 48

	if wRow < 1 || wRow > 8 || wCol < 1 || wCol > 8 {
		return false, errors.New("invalid placement")
	}

	if bRow < 1 || bRow > 8 || bCol < 1 || bCol > 8 {
		return false, errors.New("invalid placement")
	}

	if bRow == wRow && bCol == wCol {
		return false, errors.New("invalid placement")
	}

	if bRow == wRow || bCol == wCol || (bRow - bCol) == (wRow - wCol) || (bRow + bCol) == (wRow + wCol) {
		return true, nil
	}

	return false, nil
}
