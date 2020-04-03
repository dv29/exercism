package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	var triplets []Triplet

	for a := min; a <= max; a++ {
		for b := a+1; b <= max; b++ {
			for c := b+1; c <= max; c++ {
				if a*a + b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}

	return triplets
}

func Sum(n int) []Triplet {
	var triplets []Triplet

	for a := 1; a <= n; a++ {
		for b := a+1; b <= n; b++ {
			for c := b+1; c <= n; c++ {
				if a+b+c == n &&  a*a + b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}

	return triplets
}


