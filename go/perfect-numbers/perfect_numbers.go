package perfect

import (
	"errors"
)

// Classification for type of number
type Classification int

// Classify type of natural number
const (
	ClassificationAbundant Classification = iota
	ClassificationDeficient
	ClassificationPerfect
)

// ErrOnlyPositive error if the value is not positive
var ErrOnlyPositive error = errors.New("only positive values allowed")

func factors(in int64) []int64 {
	f := []int64{}

	for i := int64(1); i <= in/2; i++ {
		if in%i == 0 {
			f = append(f, i)
		}
	}
	return f
}

func sum(in []int64) int64 {
	sum := int64(0)
	for _, v := range in {
		sum += v
	}
	return sum
}

// Classify classify the natural number
func Classify(in int64) (Classification, error) {
	if in < 1 {
		return 0, ErrOnlyPositive
	}

	s := sum(factors(in))

	switch {
	case s == in:
		return ClassificationPerfect, nil
	case s > in:
		return ClassificationAbundant, nil
	case s < in:
		return ClassificationDeficient, nil
	}

	return 0, nil
}
