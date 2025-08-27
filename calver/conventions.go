package calver

// ConventionsRegex is a map of conventions to their regex. The conventions are
// taken the official docs at https://calver.org/#scheme
//
// CalVer and SemVer are different in nature. SemVer describes a strict schema,
// CalVer introduces a standard terminology and identifiers and lets you choose
// the format. The ConventionsRegex map describes the regex that may be used to
// extract the value against each identifier.
var ConventionsRegex = map[string]string{
	"<YYYY>":     `(?P<major>\d{4})`,
	"<YY>":       `(?P<major>\d{1,2})`,
	"<0Y>":       `(?P<major>\d{2})`,
	"<MAJOR>":    `(?P<major>\d+)`,
	"<MM>":       `(?P<minor>\d{1,2})`,
	"<0M>":       `(?P<minor>\d{2})`,
	"<MINOR>":    `(?P<minor>\d+)`,
	"<WW>":       `(?P<micro>\d{1,2})`,
	"<0W>":       `(?P<micro>\d{2})`,
	"<DD>":       `(?P<micro>\d{1,2})`,
	"<0D>":       `(?P<micro>\d{2})`,
	"<MICRO>":    `(?P<micro>\d+)`,
	"<MODIFIER>": `(?P<modifier>.*)`,
}

// ConventionPrecedence is the precedence of the conventions. The higher the
// index, the higher the precedence.
//
// The precedence is used to determine the order of the identifiers when
// comparing two versions. Note that all the identifiers are compared as
// integers except for the modifier which is compared as a string.
var ConventionPrecedence = []string{
	"<YYYY>",
	"<YY>",
	"<0Y>",
	"<MAJOR>",
	"<MM>",
	"<0M>",
	"<MINOR>",
	"<WW>",
	"<0W>",
	"<DD>",
	"<0D>",
	"<MICRO>",
	"<MODIFIER>",
}
