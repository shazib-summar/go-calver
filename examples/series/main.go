package main

import (
	"fmt"

	"github.com/shazib-summar/go-calver"
)

func main() {
	ver, err := calver.Parse(
		"Rel-<YYYY>-<0M>-<0D>.<MODIFIER>",
		"Rel-2025-07-14.alpha",
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Version string: %s\n", ver.String())
	fmt.Printf("Series (major): %s\n", ver.Series("major"))
	fmt.Printf("Series (minor): %s\n", ver.Series("minor"))
	fmt.Printf("Series (micro): %s\n", ver.Series("micro"))
	fmt.Printf("Series (modifier): %s\n", ver.Series("modifier"))
}
