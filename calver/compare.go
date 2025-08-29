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
//	ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%d\n", ver1.Compare(ver2)) // -1
//
// The comparison is done in the following order: major, minor, micro, modifier.
// The major, minor and micro are compared as integers. The modifier is compared
// as a string.
func (c *Version) Compare(other *Version) (int, error) {
	if c.Format != other.Format {
		return 0, fmt.Errorf("formats do not match: %s and %s", c.Format, other.Format)
	}

	if c.Major != "" && other.Major != "" {
		if c.Major != other.Major {
			majorCurrent, err := strconv.Atoi(c.Major)
			if err != nil {
				return 0, fmt.Errorf("major is not an integer: %s", c.Major)
			}
			majorOther, err := strconv.Atoi(other.Major)
			if err != nil {
				return 0, fmt.Errorf("major is not an integer: %s", other.Major)
			}
			if majorCurrent < majorOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.Minor != "" && other.Minor != "" {
		if c.Minor != other.Minor {
			minorCurrent, err := strconv.Atoi(c.Minor)
			if err != nil {
				return 0, fmt.Errorf("minor is not an integer: %s", c.Minor)
			}
			minorOther, err := strconv.Atoi(other.Minor)
			if err != nil {
				return 0, fmt.Errorf("minor is not an integer: %s", other.Minor)
			}
			if minorCurrent < minorOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.Micro != "" && other.Micro != "" {
		if c.Micro != other.Micro {
			microCurrent, err := strconv.Atoi(c.Micro)
			if err != nil {
				return 0, fmt.Errorf("micro is not an integer: %s", c.Micro)
			}
			microOther, err := strconv.Atoi(other.Micro)
			if err != nil {
				return 0, fmt.Errorf("micro is not an integer: %s", other.Micro)
			}
			if microCurrent < microOther {
				return -1, nil
			}
			return 1, nil
		}
	}

	if c.Modifier != "" && other.Modifier != "" {
		modifierCurrent, errModifierCurrent := strconv.Atoi(c.Modifier)
		modifierOther, errModifierOther := strconv.Atoi(other.Modifier)
		if errModifierCurrent != nil || errModifierOther != nil {
			return strings.Compare(c.Modifier, other.Modifier), nil
		}
		if modifierCurrent < modifierOther {
			return -1, nil
		}
		return 1, nil
	}

	return 0, nil
}

// CompareOrPanic is just Compare, but panics if there's an error.
//
// Example:
//
//	ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//	if err != nil {
//	    return err
//	}
//	ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
//	if err != nil {
//	    return err
//	}
//	fmt.Printf("%d\n", ver1.CompareOrPanic(ver2)) // -1
//
// This is useful when you are sure the comparison will succeed and do not want
// an error return.
func (c *Version) CompareOrPanic(other *Version) int {
	compare, err := c.Compare(other)
	if err != nil {
		panic(err)
	}
	return compare
}

// Equal returns true if the versions are equal, false otherwise.
// If Compare() returns an error, Equal will propagate it.
// Example:
//
// ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
//
//	if err != nil {
//	    return err
//	}
//
// ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// if err != nil
//
//	    return err
//	}
//
// fmt.Printf("%t\n", ver1.Equal(ver2)) // true
func (c *Version) Equal(other *Version) (bool, error) {
	compare, err := c.Compare(other)
	if err != nil {
		return false, err
	}
	return compare == 0, nil
}

// EqualOrPanic is just Equal, but panics if there's an error.
// // Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.EqualOrPanic(ver2)) // false
func (c *Version) EqualOrPanic(other *Version) bool {
	equal, err := c.Equal(other)
	if err != nil {
		panic(err)
	}
	return equal
}

// Less returns true if the current version is less than the other version.
// If Compare() returns an error, Less will propagate it.
// Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.Less(ver2)) // true
func (c *Version) Less(other *Version) (bool, error) {
	compare, err := c.Compare(other)
	if err != nil {
		return false, err
	}
	return compare == -1, nil
}

// LessOrPanic is just Less, but panics if there's an error.
// Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.LessOrPanic(ver2)) // true
func (c *Version) LessOrPanic(other *Version) bool {
	less, err := c.Less(other)
	if err != nil {
		panic(err)
	}
	return less
}

// Greater returns true if the current version is greater than the other version.
// If Compare() returns an error, Greater will propagate it.
// Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.Greater(ver2)) // false
func (c *Version) Greater(other *Version) (bool, error) {
	compare, err := c.Compare(other)
	if err != nil {
		return false, err
	}
	return compare == 1, nil
}

// GreaterOrPanic is just Greater, but panics if there's an error.
// Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.GreaterOrPanic(ver2)) // false
func (c *Version) GreaterOrPanic(other *Version) bool {
	greater, err := c.Greater(other)
	if err != nil {
		panic(err)
	}
	return greater
}

// LessOrEqual returns true if the current version is less than or equal to the other version.
// If Compare() returns an error, LessOrEqual will propagate it.
// Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.LessOrEqual(ver2)) // true
func (c *Version) LessOrEqual(other *Version) (bool, error) {
	compare, err := c.Compare(other)
	if err != nil {
		return false, err
	}
	return compare == -1 || compare == 0, nil
}

// LessOrEqualOrPanic is just LessOrEqual, but panics if there's an error.
// Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.LessOrEqualOrPanic(ver2)) // true

func (c *Version) LessOrEqualOrPanic(other *Version) bool {
	lessOrEqual, err := c.LessOrEqual(other)
	if err != nil {
		panic(err)
	}
	return lessOrEqual
}

// GreaterOrEqual returns true if the current version is greater than or equal to the other version.
// If Compare() returns an error, GreaterOrEqual will propagate it.
// Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.GreaterOrEqual(ver2)) // false

func (c *Version) GreaterOrEqual(other *Version) (bool, error) {
	compare, err := c.Compare(other)
	if err != nil {
		return false, err
	}
	return compare == 1 || compare == 0, nil
}

// GreaterOrEqualOrPanic is just GreaterOrEqual, but panics if there's an error.
// Example:
// // ver1, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-14")
// // if err != nil {
// //     return err
// // }
// // ver2, err := calver.NewVersion("Rel-<YYYY>-<0M>-<0D>", "Rel-2025-07-15")
// // if err != nil {
// //     return err
// // }
// // fmt.Printf("%t\n", ver1.GreaterOrEqualOrPanic(ver2)) // false
func (c *Version) GreaterOrEqualOrPanic(other *Version) bool {
	greaterOrEqual, err := c.GreaterOrEqual(other)
	if err != nil {
		panic(err)
	}
	return greaterOrEqual
}
