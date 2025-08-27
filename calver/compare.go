package calver

import (
	"fmt"
	"strconv"
	"strings"
)

// Compare returns 0 if the versions are equal, -1 if the current version is
// less than the other version, and 1 if the current version is greater than the
// other version. If the formats do not match, it returns an error.
//
// Example:
//
//	calver1, err := NewCalVer("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	calver2, err := NewCalVer("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%d\n", calver1.Compare(calver2)) // -1
//
// The comparison is done in the following order: major, minor, micro, modifier.
// The major, minor and micro are compared as integers. The modifier is compared
// as a string.
func (c *CalVer) Compare(other *CalVer) (int, error) {
	if c.Format != other.Format {
		return 0, fmt.Errorf("formats do not match: %s and %s", c.Format, other.Format)
	}

	if c.Major != "" && other.Major != "" {
		if c.Major != other.Major {
			majorCurrent, _ := strconv.Atoi(c.Major)
			majorOther, _ := strconv.Atoi(other.Major)
			if majorCurrent == majorOther {
				return 0, nil
			}
			if majorCurrent < majorOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.Minor != "" && other.Minor != "" {
		if c.Minor != other.Minor {
			minorCurrent, _ := strconv.Atoi(c.Minor)
			minorOther, _ := strconv.Atoi(other.Minor)
			if minorCurrent == minorOther {
				return 0, nil
			}
			if minorCurrent < minorOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.Micro != "" && other.Micro != "" {
		if c.Micro != other.Micro {
			microCurrent, _ := strconv.Atoi(c.Micro)
			microOther, _ := strconv.Atoi(other.Micro)
			if microCurrent == microOther {
				return 0, nil
			}
			if microCurrent < microOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.Modifier != "" && other.Modifier != "" {
		return strings.Compare(c.Modifier, other.Modifier), nil
	}

	return 0, nil
}

// CompareOrPanic is just Compare, but panics if there's an error.
//
// Example:
//
//	calver1, err := NewCalVer("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	calver2, err := NewCalVer("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%d\n", calver1.CompareOrPanic(calver2)) // -1
//
// This is useful when you are sure the comparison will succeed and do not want
// an error return.
func (c *CalVer) CompareOrPanic(other *CalVer) int {
	compare, err := c.Compare(other)
	if err != nil {
		panic(err)
	}
	return compare
}
