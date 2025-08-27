package internal

import (
	"fmt"
	"strconv"
)

func IncWithPadding(in string) (string, error) {
	if in == "" {
		return "", nil
	}
	current, err := strconv.Atoi(in)
	if err != nil {
		return "", fmt.Errorf("input is not a number: %w", err)
	}
	next := current + 1
	originalLen := len(in)
	nextStr := strconv.Itoa(next)

	if originalLen > len(nextStr) {
		return fmt.Sprintf("%0*d", originalLen, next), nil
	} else {
		return nextStr, nil
	}
}
