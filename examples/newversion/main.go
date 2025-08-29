package main

import (
	"fmt"

	"github.com/shazib-summar/go-calver"
)

func main() {
	entries := []map[string]string{
		{
			"format":  "Rel-<YYYY>-<0M>-<0D>",
			"version": "Rel-2025-07-14",
		},
		{
			"format":  "<YYYY>.<0M>.<0D>",
			"version": "2025.07.14",
		},
		{
			"format":  "<YYYY>/<0M>/<0D>",
			"version": "2025/07/14",
		},
		{
			"format":  "<YYYY>-Rel<MINOR>",
			"version": "2025-Rel07",
		},
		{
			"format":  "<YYYY>-Rel<MINOR>/<MICRO>",
			"version": "2025-Rel07/14",
		},
	}

	for _, entry := range entries {
		ver, err := calver.NewVersion(entry["format"], entry["version"])
		if err != nil {
			panic(err)
		}
		fmt.Println(ver.String())
	}
}
