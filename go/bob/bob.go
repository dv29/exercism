// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
  "strings"
  "regexp"
)

var HasLetter = regexp.MustCompile(`[a-zA-Z]+`).MatchString

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
  remark = strings.TrimSpace(remark)

  if remark == "" {
    return "Fine. Be that way!"
  }

  if strings.ToUpper(remark) == remark && HasLetter(remark) {
    if remark[len(remark)-1] == '?' {
      return "Calm down, I know what I'm doing!"
    }
    return "Whoa, chill out!"
  }
  if remark[len(remark)-1] == '?' {
    return "Sure."
  }
	return "Whatever."
}
