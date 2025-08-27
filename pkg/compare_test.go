package pkg_test

import (
	"testing"

	"github.com/shazib-summar/go-calver/pkg"
	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {

	tests := []struct {
		name    string
		format  string
		version string
		other   string
		want    int
	}{
		{
			name:    "1",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R2",
			want:    -1,
		},
		{
			name:    "2",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    -1,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    0,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    1,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    1,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    1,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    -1,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    -1,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    1,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    0,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    -1,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    -1,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			version, err := pkg.NewCalVer(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := pkg.NewCalVer(tt.format, tt.other)
			assert.NoError(t, err)
			got, err := version.Compare(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
