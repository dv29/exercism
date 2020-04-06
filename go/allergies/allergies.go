package allergies

import (
	"sort"
)

const (
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var allergyMap = map[string]uint{
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}

// Allergies returns the allergies for the score
func Allergies(score uint) []string {
	score %= 256
	str := []string{}
	for k, v := range allergyMap {
		if (score & v) != 0 {
			str = append(str, k)
		}
	}

	sort.Slice(str, func(i, j int) bool {
		return allergyMap[str[i]] < allergyMap[str[j]]
	})

	return str
}

// AllergicTo returns if the score is allergic to substance
func AllergicTo(score uint, substance string) bool {
	v := allergyMap[substance]
	return (score & v) == v
}
