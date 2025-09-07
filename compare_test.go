package calver_test

import (
	"testing"

	"github.com/shazib-summar/go-calver"
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
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    -1,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    1,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    1,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    -1,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.Parse(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.Parse(tt.format, tt.other)
			assert.NoError(t, err)
			got := ver.Compare(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		version string
		other   string
		want    bool
	}{
		{
			name:    "1",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R2",
			want:    false,
		},
		{
			name:    "2",
			format:  "<YYYY>-R<DD>",
			version: "2022-R1",
			other:   "2025-R1",
			want:    false,
		},
		{
			name:    "3",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2015-R1",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.Parse(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.Parse(tt.format, tt.other)
			assert.NoError(t, err)
			got := ver.Equal(other)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLessThan(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		version string
		other   string
		want    bool
	}{
		{
			name:    "1",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R2",
			want:    true,
		},
		{
			name:    "2",
			format:  "<YYYY>-R<DD>",
			version: "2022-R1",
			other:   "2025-R1",
			want:    true,
		},
		{
			name:    "3",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2015-R1",
			want:    false,
		},
		{
			name:    "4",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R1",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.Parse(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.Parse(tt.format, tt.other)
			assert.NoError(t, err)
			got := ver.LessThan(other)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGreaterThan(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		version string
		other   string
		want    bool
	}{
		{
			name:    "1",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R2",
			want:    false,
		},
		{
			name:    "2",
			format:  "<YYYY>-R<DD>",
			version: "2025-R0",
			other:   "2025-R1",
			want:    false,
		},
		{
			name:    "3",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2015-R1",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.Parse(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.Parse(tt.format, tt.other)
			assert.NoError(t, err)
			got := ver.GreaterThan(other)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLessThanOrEqual(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		version string
		other   string
		want    bool
	}{
		{
			name:    "1",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R2",
			want:    true,
		},
		{
			name:    "2",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R1",
			want:    true,
		},
		{
			name:    "3",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2015-R1",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.Parse(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.Parse(tt.format, tt.other)
			assert.NoError(t, err)
			got := ver.LessThanOrEqual(other)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGreaterThanOrEqual(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		version string
		other   string
		want    bool
	}{
		{
			name:    "1",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R2",
			want:    false,
		},
		{
			name:    "2",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R1",
			want:    true,
		},
		{
			name:    "3",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2015-R1",
			want:    true,
		},
		{
			name:    "4",
			format:  "<YYYY>-R<DD>",
			version: "2025-R1",
			other:   "2025-R2",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.Parse(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.Parse(tt.format, tt.other)
			assert.NoError(t, err)
			got := ver.GreaterThanOrEqual(other)
			assert.Equal(t, tt.want, got)
		})
	}
}
