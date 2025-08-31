package calver

import "fmt"

// Collection is a collection of Version objects. It implements the
// sort.Interface interface.
type Collection []*Version

// NewCollection creates a new Collection from a format string and a list of
// versions. It will return an error if any of the versions do not match the
// format.
//
// Example:
//
//	collection, err := calver.NewCollection(
//	    "<YYYY>.<0M>.<0D>",
//	    "2025.01.18",
//	    "2023.07.14",
//	    "2025.03.16",
//	)
//	if err != nil {
//	    return err
//	}
//
// This is the same as calling `NewCollectionWithOptions(versions,
// WithFormat(format))`. Note that `WithFormat` option can take a list of
// formats.
func NewCollection(format string, versions ...string) (Collection, error) {
	return NewCollectionWithOptions(versions, WithFormat(format))
}

func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Less(i, j int) bool {
	return c[i].Compare(c[j]) < 0
}

func (c Collection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// NewCollectionWithOptions creates a new `Collection` from a list of versions and
// a list of parse options. It will return an error if any of the versions do
// not match (any of) the format or if no options are provided.
//
// Example:
//
//	collection, err := calver.NewCollectionWithOptions(
//	    []string{"2025.01.18", "2023.07.14", "2025.03.16"},
//	    calver.WithFormat("<YYYY>.<0M>.<0D>"),
//	)
//	if err != nil {
//	    return err
//	}
//
// If there are more than one possible format that may match the version string,
// this function can be used with the `WithFormat` option.
//
// Example:
//
//	formats := []string{
//	    "<YYYY>.<0M>.<0D>",
//	    "<YYYY>.<0M>.<0D>-<MODIFIER>",
//	}
//	collection, err := calver.NewCollectionWithOptions(
//	    []string{"2025.01.18", "2023.07.14", "2025.03.16"},
//	    calver.WithFormat(formats...),
//	)
//	if err != nil {
//	    return err
//	}
func NewCollectionWithOptions(versions []string, opts ...parseOption) (Collection, error) {
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

	collection := make(Collection, len(versions))
	for i, version := range versions {
		calver, err := ParseWithOptions(version, opts...)
		if err != nil {
			return nil, err
		}
		collection[i] = calver
	}
	return collection, nil
}
