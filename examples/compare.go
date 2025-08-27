package examples

import (
	"fmt"

	"github.com/shazib-summar/go-calver/calver"
)

func ex_compare() {
	format := "<YYYY>-Rel<MINOR>/<MICRO>"
	versionA := "2025-Rel07/14"
	versionB := "2025-Rel07/15"

	verA, err := calver.NewCalVer(format, versionA)
	if err != nil {
		panic(err)
	}

	verB, err := calver.NewCalVer(format, versionB)
	if err != nil {
		panic(err)
	}

	if verA.CompareOrPanic(verB) == 0 {
		fmt.Println("Versions are equal")
	} else if verA.CompareOrPanic(verB) > 0 {
		fmt.Println("Version A is greater than Version B")
	} else {
		fmt.Println("Version A is less than Version B")
	}
}
