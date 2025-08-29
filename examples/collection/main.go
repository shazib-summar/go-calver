package main

import (
	"fmt"
	"sort"

	"github.com/shazib-summar/go-calver"
)

func main() {
	format := "<YYYY>.<0M>.<0D>"
	versions := []string{
		"2025.01.18",
		"2023.07.14",
		"2025.03.16",
		"2025.07.15",
		"2025.05.17",
		"2021.07.19",
	}

	coll, err := calver.NewCollection(versions, format)
	if err != nil {
		panic(err)
	}

	// The Collection object implements the sort.Interface interface
	sort.Sort(coll)

	for i, v := range coll {
		fmt.Printf("%d: %s\n", i, v.String())
	}

	// Get the smallest version
	smallest := coll[0]
	fmt.Println("Smallest version:", smallest.String())

	// Get the largest version
	largest := coll[len(coll)-1]
	fmt.Println("Largest version:", largest.String())
}
