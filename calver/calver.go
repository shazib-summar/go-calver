package calver

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/shazib-summar/go-calver/internal"
)

// CalVer is a CalVer object.
type CalVer struct {
	format   string
	major    string
	minor    string
	micro    string
	modifier string
}

// NewCalVer creates a new CalVer object from a format string and a version. The
// format string is expected to follow the conventions defined in
// ConventionsRegex.
//
// Example:
//
//	ver, err := NewCalVer("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//		return err
//	}
//	fmt.Println(ver.String()) // Rel-2025-07-14
func NewCalVer(format string, version string) (*CalVer, error) {
	if !internal.ValidateFormat(format) {
		return nil, fmt.Errorf("invalid format: %s", format)
	}

	originalFormat := format
	format = strings.ReplaceAll(format, ".", `\.`)
	for _, con := range internal.ValidConventions {
		format = strings.ReplaceAll(format, con, internal.ConventionsRegex[con])
	}

	format = `^` + format + `$`
	re := regexp.MustCompile(format)
	groups := re.FindStringSubmatch(version)
	if len(groups) == 0 {
		return nil, fmt.Errorf(
			"version %s does not match format: %s",
			version,
			format,
		)
	}

	c := &CalVer{
		format: originalFormat,
	}
	for i, lv := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		switch lv {
		case internal.KeyMajor:
			c.major = groups[i]
		case internal.KeyMinor:
			c.minor = groups[i]
		case internal.KeyMicro:
			c.micro = groups[i]
		case internal.KeyModifier:
			c.modifier = groups[i]
		}
	}
	if c.major == "" && c.minor == "" && c.micro == "" && c.modifier == "" {
		return nil, fmt.Errorf(
			"malformed calver format: %s - "+
				"make sure to use at least one version",
			version,
		)
	}

	return c, nil
}

// String returns the CalVer object as a string. The string will be in the
// format of the original format string.
//
// Example:
//
//	calver, err := NewCalVer("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//		return err
//	}
//	fmt.Println(calver.String()) // Rel-2025-07-14
func (c *CalVer) String() string {
	out := c.format
	versionParts := []string{c.major, c.minor, c.micro, c.modifier}
	for i, lv := range internal.ValidLevels {
		for _, con := range internal.ConventionsByLevel[lv] {
			if versionParts[i] != "" {
				out = strings.ReplaceAll(out, con, versionParts[i])
			}
		}
	}
	return out
}
