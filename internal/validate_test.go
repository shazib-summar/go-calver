package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateFormat(t *testing.T) {
	tests := []struct {
		name   string
		format string
		want   bool
	}{
		{name: "1", format: "<YYYY>-<MM>-<DD>", want: true},
		{name: "2", format: "<YYYY>-<MM>-<DD>-<MM>", want: false},
		{name: "2", format: "<YYYY>-<YYYY>", want: false},
		{name: "3", format: "<MAJOR>-<MAJOR>", want: false},
		{name: "4", format: "<MAJOR>-<MINOR>-<MICRO>", want: true},
		{name: "5", format: "<MAJOR>-<MINOR>-<MICRO>-<MICRO>", want: false},
		{name: "6", format: "foobar", want: false},
		{name: "7", format: "foobar-<MICRO>", want: true},
		{name: "8", format: "foobar-<YYYY>", want: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := ValidateFormat(test.format)
			assert.Equal(t, test.want, got)
		})
	}
}
