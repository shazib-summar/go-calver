package calver

import (
	"strconv"
	"strings"
)

// Compare returns 0 if the versions are equal, -1 if the current version is
// less than the other version, and 1 if the current version is greater than the
// other version.
//
// Example:
//
//	ver1, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	ver2, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%d\n", ver1.Compare(ver2)) // -1
//
// The comparison is done in the following order: major, minor, micro, modifier.
// Major, minor and micro are compared as integers whereas the modifier is
// compared as integer if it is a number otherwise as a string.
func (c *Version) Compare(v *Version) int {
	res := compareStringInt(c.Major, v.Major)
	if res != 0 {
		return res
	}

	res = compareStringInt(c.Minor, v.Minor)
	if res != 0 {
		return res
	}

	res = compareStringInt(c.Micro, v.Micro)
	if res != 0 {
		return res
	}

	res = compareStringInt(c.Modifier, v.Modifier)
	if res != 0 {
		return res
	}

	return 0
}

// Equal reports whether the version is equal to the other version.
// Example:
//
//	ver1, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	ver2, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%t\n", ver1.Equal(ver2)) // false
func (c *Version) Equal(v *Version) bool {
	return c.Compare(v) == 0
}

// LessThan reports whether the version is less than the other version.
// Example:
//
//	ver1, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	ver2, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%t\n", ver1.LessThan(ver2)) // true
func (c *Version) LessThan(v *Version) bool {
	return c.Compare(v) < 0
}

// GreaterThan reports whether the version is greater than the other version.
// Example:
//
//	ver1, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	ver2, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%t\n", ver1.GreaterThan(ver2)) // false
func (c *Version) GreaterThan(v *Version) bool {
	return c.Compare(v) > 0
}

// LessThanOrEqual reports whether the version is less than or equal to the
// other version.
// Example:
//
//	ver1, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	ver2, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%t\n", ver1.LessThanOrEqual(ver2)) // true
func (c *Version) LessThanOrEqual(v *Version) bool {
	return c.Compare(v) <= 0
}

// GreaterThanOrEqual reports whether the version is greater than or equal to
// the other version.
// Example:
//
//	ver1, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	ver2, err := calver.Parse("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%t\n", ver1.GreaterThanOrEqual(ver2)) // false
func (c *Version) GreaterThanOrEqual(v *Version) bool {
	return c.Compare(v) >= 0
}

// compareStringInt compares two string integers as integers. It handles the
// case where one or both of the strings are empty. If the integer in the first
// string is larger it returns 1, if the integer in the second string is larger
// it returns -1, and if they are equal it returns 0.
func compareStringInt(a, b string) int {
	if a == b {
		return 0
	}
	if a == "" {
		return -1
	}
	if b == "" {
		return 1
	}

	aInt, errA := strconv.Atoi(a)
	bInt, errB := strconv.Atoi(b)
	if errA != nil || errB != nil {
		return strings.Compare(a, b)
	}

	if aInt < bInt {
		return -1
	}
	if aInt > bInt {
		return 1
	}
	return 0
}
