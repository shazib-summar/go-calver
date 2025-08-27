package internal

import (
	"fmt"

	"github.com/samber/lo"
)

const (
	KeyMajor    = "major"
	KeyMinor    = "minor"
	KeyMicro    = "micro"
	KeyModifier = "modifier"
)

// ValidLevels is used to determine the order of the identifiers when comparing
// two versions, so the order in this array matters. This is also used to
// validate the format string. The format string can only have a single
// convention of each level.
//
// All the identifiers are compared as integers except for the modifier which is
// compared as a string.
var ValidLevels = []string{
	KeyMajor,
	KeyMinor,
	KeyMicro,
	KeyModifier,
}

// ValidConventions is a list of valid conventions that may be used in the
// format string.
var ValidConventions = lo.Keys(ConventionsRegex)

// ConventionsRegex is a map of conventions to their regex. The conventions are
// taken the official docs at https://calver.org/#scheme
//
// CalVer and SemVer are different in nature. SemVer describes a strict schema,
// CalVer introduces a standard terminology and identifiers and lets you choose
// the format. The ConventionsRegex map describes the regex that may be used to
// extract the value against each identifier.
var ConventionsRegex = map[string]string{
	// Major
	"<YYYY>":  fmt.Sprintf(`(?P<%s>\d{4})`, KeyMajor),
	"<YY>":    fmt.Sprintf(`(?P<%s>\d{1,2})`, KeyMajor),
	"<0Y>":    fmt.Sprintf(`(?P<%s>\d{2})`, KeyMajor),
	"<MAJOR>": fmt.Sprintf(`(?P<%s>\d+)`, KeyMajor),

	// Minor
	"<MM>":    fmt.Sprintf(`(?P<%s>\d{1,2})`, KeyMinor),
	"<0M>":    fmt.Sprintf(`(?P<%s>\d{2})`, KeyMinor),
	"<MINOR>": fmt.Sprintf(`(?P<%s>\d+)`, KeyMinor),

	// Micro
	"<WW>":    fmt.Sprintf(`(?P<%s>\d{1,2})`, KeyMicro),
	"<0W>":    fmt.Sprintf(`(?P<%s>\d{2})`, KeyMicro),
	"<DD>":    fmt.Sprintf(`(?P<%s>\d{1,2})`, KeyMicro),
	"<0D>":    fmt.Sprintf(`(?P<%s>\d{2})`, KeyMicro),
	"<MICRO>": fmt.Sprintf(`(?P<%s>\d+)`, KeyMicro),

	// Modifier
	"<MODIFIER>": fmt.Sprintf(`(?P<%s>.*)`, KeyModifier),
}

// ConventionsByLevel groups the conventions by level.
var ConventionsByLevel = map[string][]string{
	KeyMajor: {
		"<YYYY>",
		"<YY>",
		"<0Y>",
		"<MAJOR>",
	},
	KeyMinor: {
		"<MM>",
		"<0M>",
		"<MINOR>",
	},
	KeyMicro: {
		"<WW>",
		"<0W>",
		"<DD>",
		"<0D>",
		"<MICRO>",
	},
	KeyModifier: {
		"<MODIFIER>",
	},
}
