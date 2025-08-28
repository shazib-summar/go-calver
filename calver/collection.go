package calver

// Collection is a collection of Version objects. It implements the
// sort.Interface interface.
type Collection []*Version

// NewCollection creates a new Collection from a format and a list of versions.
// It returns an error if any of the versions do not match the format.
func NewCollection(format string, versions ...string) (Collection, error) {
	collection := make(Collection, len(versions))
	for i, version := range versions {
		calver, err := NewVersion(format, version)
		if err != nil {
			return nil, err
		}
		collection[i] = calver
	}
	return collection, nil
}

func (c Collection) Len() int {
	return len(c)
}

func (c Collection) Less(i, j int) bool {
	return c[i].CompareOrPanic(c[j]) < 0
}

func (c Collection) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
