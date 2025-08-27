package calver

import (
	"fmt"
	"regexp"
	"strings"
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
//	calver, err := NewCalVer("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//		return err
//	}
//
// Do what you want with the calver object
//
//	fmt.Println(calver.String()) // Rel-2025-07-14
func NewCalVer(format string, version string) (*CalVer, error) {
	found := false
	for con := range ConventionsRegex {
		if strings.Contains(format, con) {
			found = true
			break
		}
	}
	if !found {
		return nil, fmt.Errorf("unrecognized calver format: %s", format)
	}
	originalFormat := format

	format = strings.ReplaceAll(format, ".", `\.`)
	for _, con := range ConventionPrecedence {
		if strings.Contains(format, con) {
			format = strings.ReplaceAll(format, con, ConventionsRegex[con])
		}
	}
	format = `^` + format + `$`
	re := regexp.MustCompile(format)
	groups := re.FindStringSubmatch(version)
	if len(groups) == 0 {
		return nil, fmt.Errorf("version %s does not match format: %s", version, format)
	}

	c := &CalVer{
		format: originalFormat,
	}
	for i, name := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		switch name {
		case "major":
			if c.major != "" {
				return nil, fmt.Errorf("malformed calver format: %s - make sure to only use one major version", version)
			}
			c.major = groups[i]
		case "minor":
			if c.minor != "" {
				return nil, fmt.Errorf("malformed calver format: %s - make sure to only use one minor version", version)
			}
			c.minor = groups[i]
		case "micro":
			if c.micro != "" {
				return nil, fmt.Errorf("malformed calver format: %s - make sure to only use one micro version", version)
			}
			c.micro = groups[i]
		case "modifier":
			if c.modifier != "" {
				return nil, fmt.Errorf("malformed calver format: %s - make sure to only use one modifier", version)
			}
			c.modifier = groups[i]
		}
	}
	if c.major == "" && c.minor == "" && c.micro == "" {
		return nil, fmt.Errorf("malformed calver format: %s - make sure to use at least one version", version)
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
	for _, con := range ConventionPrecedence {
		if strings.Contains(out, con) {
			switch con {
			case "<YYYY>", "<YY>", "<0Y>", "<MAJOR>":
				out = strings.ReplaceAll(out, con, c.major)
			case "<MM>", "<0M>", "<MINOR>":
				out = strings.ReplaceAll(out, con, c.minor)
			case "<WW>", "<0W>", "<DD>", "<0D>", "<MICRO>":
				out = strings.ReplaceAll(out, con, c.micro)
			case "<MODIFIER>":
				out = strings.ReplaceAll(out, con, c.modifier)
			}
		}
	}
	return out
}
