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
	if c.format != other.format {
		return 0, fmt.Errorf("formats do not match: %s and %s", c.format, other.format)
	}

	if c.major != "" && other.major != "" {
		if c.major != other.major {
			majorCurrent, _ := strconv.Atoi(c.major)
			majorOther, _ := strconv.Atoi(other.major)
			if majorCurrent == majorOther {
				return 0, nil
			}
			if majorCurrent < majorOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.minor != "" && other.minor != "" {
		if c.minor != other.minor {
			minorCurrent, _ := strconv.Atoi(c.minor)
			minorOther, _ := strconv.Atoi(other.minor)
			if minorCurrent == minorOther {
				return 0, nil
			}
			if minorCurrent < minorOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.micro != "" && other.micro != "" {
		if c.micro != other.micro {
			microCurrent, _ := strconv.Atoi(c.micro)
			microOther, _ := strconv.Atoi(other.micro)
			if microCurrent == microOther {
				return 0, nil
			}
			if microCurrent < microOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.modifier != "" && other.modifier != "" {
		return strings.Compare(c.modifier, other.modifier), nil
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
