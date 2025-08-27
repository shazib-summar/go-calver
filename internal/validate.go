package internal

import (
	"strings"

	"github.com/samber/lo"
)

// ValidateFormat reports if the format string is valid.
//
// The format string is valid if it contains only one convention of each level
// and contains at least one convention.
func ValidateFormat(format string) bool {
	if !lo.ContainsBy(ValidConventions, func(con string) bool {
		return strings.Contains(format, con)
	}) {
		return false
	}

	for _, lv := range ValidLevels {
		cons := ConventionsByLevel[lv]
		matches := 0
		for _, con := range cons {
			matches += strings.Count(format, con)
		}
		if matches > 1 {
			return false
		}
	}
	return true
}
