package calver

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/shazib-summar/go-calver/internal"
)

// Version is the object representing a CalVer version. To get the string
// representation of the Version, use the String method.
type Version struct {
	// Format is the original format string. If multiple formats were provided,
	// this will be the format that matched the version string.
	Format string
	// Major is the major version. This is guaranteed to be a number.
	Major string
	// Minor is the minor version. This is guaranteed to be a number.
	Minor string
	// Micro is the micro version. This is guaranteed to be a number.
	Micro string
	// Modifier is the modifier version. This can be a number or a string.
	Modifier string
}

type parseOptions struct {
	formats []string
}

type parseOption func(*parseOptions)

// WithFormat is a parse option that specifies the format string that should be
// used to parse the version string.
//
// Example:
//
//	ver, err := ParseWithOptions("2025.07.14", WithFormat("<YYYY>.<0M>.<0D>"))
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.String()) // 2025.07.14
//
// If there are more than one possible format that may match the version string,
// this function can be used with the WithFormat option.
//
// Example:
//
//	formats := []string{
//	    "<YYYY>.<0M>.<0D>",
//	    "<YYYY>.<0M>.<0D>-<MODIFIER>",
//	}
//	ver, err := ParseWithOptions("2025.07.14", WithFormat(formats...))
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.String()) // 2025.07.14
//
// If multiple formats are provided, the format that matches the version string
// will be used. For example, in the following code, the format
// `"Rel-<YYYY>-<0M>-<0D>"` will be used at it matches the version string while
// the other formats do not.
//
//	ver, err := ParseWithOptions(
//	    "Rel-2025-07-14",
//	    WithFormat(
//	        "Rel-<YYYY>",
//	        "Rel-<YYYY>-<0M>",
//	        "Rel-<YYYY>-<0M>-<0D>",
//	    )
//	)
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.String()) // Rel-2025-07-14
func WithFormat(formats ...string) parseOption {
	return func(options *parseOptions) {
		options.formats = formats
	}
}

// Parse creates a new Version object from a format string and a version. The
// format string is expected to follow the conventions defined in
// ConventionsRegex.
//
// Example:
//
//	ver, err := Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.String()) // Rel-2025-07-14
//
// This is the same as calling `ParseWithOptions(version, WithFormat(format))`
func Parse(format string, version string) (*Version, error) {
	return ParseWithOptions(version, WithFormat(format))
}

// ParseWithOptions creates a new Version object from a version string and a
// list of parse options.
//
// Example:
//
//	ver, err := ParseWithOptions(
//	    "Rel-2025-07-14",
//	    WithFormat("Rel-<YYYY>-<0M>-<0D>"),
//	)
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.String()) // Rel-2025-07-14
//
// If there are more than one possible format that may match the version string,
// this function can be used with the WithFormat option.
//
// Example:
//
//	formats := []string{
//	    "Rel-<YYYY>-<0M>-<0D>",
//	    "<YYYY>.<0M>.<0D>",
//	}
//	ver, err := ParseWithOptions("Rel-2025-07-14", WithFormat(formats...))
//	if err != nil {
//	    return err
//	}
//	fmt.Println(ver.String()) // Rel-2025-07-14
func ParseWithOptions(version string, opts ...parseOption) (*Version, error) {
	if len(opts) == 0 {
		return nil, fmt.Errorf("at least one parseOption is required")
	}
	o := &parseOptions{}
	for _, opt := range opts {
		opt(o)
	}

	if len(o.formats) == 0 {
		return nil, fmt.Errorf("no format provided")
	}

	var matchingFormat string
	var re *regexp.Regexp
	var groups []string
	for _, f := range o.formats {
		f_ := f // save a copy before modifying it
		if !internal.ValidateFormat(f) {
			return nil, fmt.Errorf("invalid format: %s", f)
		}

		f = strings.ReplaceAll(f, ".", `\.`)
		for _, con := range internal.ValidConventions {
			f = strings.ReplaceAll(f, con, internal.ConventionsRegex[con])
		}

		f = `^` + f + `$`
		currRe := regexp.MustCompile(f)
		currGroups := currRe.FindStringSubmatch(version)
		if len(currGroups) > len(groups) {
			matchingFormat = f_
			re = currRe
			groups = currGroups
		}
	}

	if len(groups) == 0 {
		return nil, fmt.Errorf(
			"version %q does not match any of the provided formats: %q",
			version,
			strings.Join(o.formats, ", "),
		)
	}

	c := &Version{
		Format: matchingFormat,
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
//	ver, err := Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
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
// will retain the 0 padding unless the major version is of the form 09 or 099
// or 0999 and so on.
func (c *Version) IncMajor() error {
	major, _ := internal.IncWithPadding(c.Major)
	c.Major = major
	return nil
}

// IncMinor increments the minor version. If the minor version is 0 padded it
// will retain the 0 padding unless the minor version is of the form 09 or 099
// or 0999 and so on.
func (c *Version) IncMinor() error {
	minor, _ := internal.IncWithPadding(c.Minor)
	c.Minor = minor
	return nil
}

// IncMicro increments the micro version. If the micro version is 0 padded it
// will retain the 0 padding unless the micro version is of the form 09 or 099
// or 0999 and so on.
func (c *Version) IncMicro() error {
	micro, _ := internal.IncWithPadding(c.Micro)
	c.Micro = micro
	return nil
}

// IncModifier increments the modifier version. If the modifier version is 0
// padded it will retain the 0 padding unless the modifier version is of the
// form 09 or 099 or 0999 and so on.
//
// It will return an error if the modifier is not a number.
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
//	ver, err := Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
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
