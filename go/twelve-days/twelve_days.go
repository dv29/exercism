package twelve

import (
	"fmt"
	"strings"
)

var deeds = []string{
	"twelve Drummers Drumming",
	"eleven Pipers Piping",
	"ten Lords-a-Leaping",
	"nine Ladies Dancing",
	"eight Maids-a-Milking",
	"seven Swans-a-Swimming",
	"six Geese-a-Laying",
	"five Gold Rings",
	"four Calling Birds",
	"three French Hens",
	"two Turtle Doves",
	"a Partridge in a Pear Tree",
}

var dayStrings = []string{
	"twelfth",
	"eleventh",
	"tenth",
	"ninth",
	"eighth",
	"seventh",
	"sixth",
	"fifth",
	"fourth",
	"third",
	"second",
	"first",
}

func firstPart(dayString string) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me:", dayString)
}

func Song() string {
	s := ""
	for i := 1; i < 12; i++ {
		s += Verse(i) + "\n"
	}
	s += Verse(12)
	return s
}

func Verse(i int) string {
	i--
	if i == 0 {
		return fmt.Sprintf("%s %s.", firstPart(dayStrings[11]), deeds[11])
	}

	secondPart := strings.Join(deeds[11-i:11], ", ")
	return fmt.Sprintf("%s %s, and %s.", firstPart(dayStrings[11-i]), secondPart, deeds[11])
}
