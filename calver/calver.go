package calver

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/shazib-summar/go-calver/internal"
)

// Version is a Version object. To get the string representation of the version,
// use the String method.
type Version struct {
	// Format is the original format string.
	Format string
	// Major is the major version.
	Major string
	// Minor is the minor version.
	Minor string
	// Micro is the micro version.
	Micro string
	// Modifier is the modifier version.
	Modifier string
}

// NewVersion creates a new Version object from a format string and a version.
// The format string is expected to follow the conventions defined in
// ConventionsRegex.
//
// Example:
//
//	ver, err := NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.String()) // Rel-2025-07-14
func NewVersion(format string, version string) (*Version, error) {
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

	c := &Version{
		Format: originalFormat,
	}
	for i, lv := range re.SubexpNames() {
		if i == 0 {
			continue
		}
		switch lv {
		case internal.KeyMajor:
			c.Major = groups[i]
		case internal.KeyMinor:
			c.Minor = groups[i]
		case internal.KeyMicro:
			c.Micro = groups[i]
		case internal.KeyModifier:
			c.Modifier = groups[i]
		}
	}
	if c.Major == "" && c.Minor == "" && c.Micro == "" && c.Modifier == "" {
		return nil, fmt.Errorf(
			"malformed calver format: %s - "+
				"make sure to use at least one version",
			version,
		)
	}

	return c, nil
}

// String returns the Version object as a string. The string will be in the
// format of the original format string.
//
// Example:
//
//	ver, err := NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.String()) // Rel-2025-07-14
func (c *Version) String() string {
	out := c.Format
	versionParts := []string{c.Major, c.Minor, c.Micro, c.Modifier}
	for i, lv := range internal.ValidLevels {
		for _, con := range internal.ConventionsByLevel[lv] {
			if versionParts[i] != "" {
				out = strings.ReplaceAll(out, con, versionParts[i])
			}
		}
	}
	return out
}

// GetMajor returns the major version.
func (c *Version) GetMajor() string {
	return c.Major
}

// GetMinor returns the minor version.
func (c *Version) GetMinor() string {
	return c.Minor
}

// GetMicro returns the micro version.
func (c *Version) GetMicro() string {
	return c.Micro
}

// GetModifier returns the modifier version.
func (c *Version) GetModifier() string {
	return c.Modifier
}

// GetFormat returns the original format string.
func (c *Version) GetFormat() string {
	return c.Format
}

// IncMajor increments the major version. If the major version is 0 padded it
// will retain the 0 padding unless the major version if of the form 09 or 099
// or 0999 and so on.
func (c *Version) IncMajor() error {
	major, err := internal.IncWithPadding(c.Major)
	if err != nil {
		return err
	}
	c.Major = major
	return nil
}

// IncMinor increments the minor version. If the minor version is 0 padded it
// will retain the 0 padding unless the minor version if of the form 09 or 099
// or 0999 and so on.
func (c *Version) IncMinor() error {
	minor, err := internal.IncWithPadding(c.Minor)
	if err != nil {
		return err
	}
	c.Minor = minor
	return nil
}

// IncMicro increments the micro version. If the micro version is 0 padded it
// will retain the 0 padding unless the micro version if of the form 09 or 099
// or 0999 and so on.
func (c *Version) IncMicro() error {
	micro, err := internal.IncWithPadding(c.Micro)
	if err != nil {
		return err
	}
	c.Micro = micro
	return nil
}

// IncModifier increments the modifier version. If the modifier version is 0
// padded it will retain the 0 padding unless the modifier version if of the
// form 09 or 099 or 0999 and so on.
//
// Be careful with this. Only use if the modifier is a number.
func (c *Version) IncModifier() error {
	modifier, err := internal.IncWithPadding(c.Modifier)
	if err != nil {
		return err
	}
	c.Modifier = modifier
	return nil
}

// Series returns the series of the Version object. The series determined using
// the provided level. For example, if the level is major, the series will be
// the major version. If the level is minor, the series will be the major and
// minor version and so on.
//
// If no level or an unrecognized level is provided, the series will be the
// entire version string.
//
// Example:
//
//	ver, err := NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.Series("major"))    // Rel-2025
//	fmt.Println(ver.Series("minor"))    // Rel-2025-07
//	fmt.Println(ver.Series("micro"))    // Rel-2025-07-14
//	fmt.Println(ver.Series("modifier")) // Rel-2025-07-14-0
//	fmt.Println(ver.Series(""))         // Rel-2025-07-14
func (c *Version) Series(level string) string {
	level = strings.ToLower(level)
	if !slices.Contains(internal.ValidLevels, level) {
		return c.String()
	}

	conForLevel := ""
	for _, con := range internal.ConventionsByLevel[level] {
		if strings.Contains(c.Format, con) {
			conForLevel = con
			break
		}
	}
	if conForLevel == "" {
		return c.String()
	}

	parts := strings.Split(c.GetFormat(), conForLevel)
	if len(parts) == 1 {
		return c.String()
	}
	format := parts[0] + conForLevel

	for _, lv := range internal.ValidLevels {
		cons := internal.ConventionsByLevel[lv]
		for _, con := range cons {
			format = strings.ReplaceAll(format, con, getValueForLevel(c, lv))
		}
		if lv == level {
			break
		}
	}
	return format
}

func getValueForLevel(c *Version, level string) string {
	if level == internal.KeyMajor {
		return c.Major
	}
	if level == internal.KeyMinor {
		return c.Minor
	}
	if level == internal.KeyMicro {
		return c.Micro
	}
	if level == internal.KeyModifier {
		return c.Modifier
	}
	return ""
}
