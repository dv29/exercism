package palindrome

import "errors"

type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

func isPalindrom(x int) bool {
	if x < 0 {
		x *= -1
	}

	y := x
	reverse := 0

	for y > 0 {
		digit := y % 10
		reverse = reverse*10 + digit
		y /= 10
	}
	return reverse == x
}

func Products(fMin, fMax int) (pMin, pMax Product, err error) {
	if fMin > fMax {
		err = errors.New("fmin > fmax")
		return
	}

	for x := fMin; x <= fMax; x++ {
		for y := x; y <= fMax; y++ {
			product := x * y
			if isPalindrom(product) {
				if pMin.Product > product || pMin.Product == 0 {
					pMin.Product = product
					pMin.Factorizations = [][2]int{{x, y}}
				}
				if pMax.Product == product {
					pMax.Factorizations = append(pMax.Factorizations, [2]int{x, y})
				}
				if pMax.Product < product {
					pMax.Product = product
					pMax.Factorizations = [][2]int{{x, y}}
				}
			}
		}
	}
	if pMin.Product == 0 {
		err = errors.New("no palindromes")
	}
	return
}
