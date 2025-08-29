package calver

// Collection is a slice of Version objects. It implements the sort.Interface
// interface.
type Collection []*Version

// NewCollection creates a new Collection from a list of versions and format(s).
// It returns an error if any of the versions do not match any of the formats.
//
// Example:
//
//	versions := []string{"2025-07-14", "2025-07-15"}
//	formats := []string{"<YYYY>-<MM>-<DD>"}
//	collection, err := calver.NewCollection(
//	    versions,
//	    formats...,
//	)
//	if err != nil {
//	    return err
//	}
//	sort.Sort(collection)
func NewCollection(versions []string, format ...string) (Collection, error) {
	collection := make(Collection, len(versions))
	for i, version := range versions {
		calver, err := NewVersion(version, format...)
		if err != nil {
			return nil, err
		}
		collection[i] = calver
	}
	return collection, nil
}

// Len returns the length of the collection.
func (c Collection) Len() int {
	return len(c)
}

// Less returns true if the version at index i is less than the version at index
// j.
func (c Collection) Less(i, j int) bool {
	return c[i].Compare(c[j]) < 0
}

// Swap swaps the elements with indexes i and j.
func (c Collection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
