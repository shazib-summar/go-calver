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
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			got, err := ver.Compare(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCompareOrPanicNonPanicking(t *testing.T) {
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
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.Equal(t, nil, rec)
			}()
			got := ver.CompareOrPanic(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCompareOrPanicPanicking(t *testing.T) {
	tests := []struct {
		name    string
		format1 string
		format2 string
		version string
		other   string
	}{
		{
			name:    "1",
			format1: "<YYYY>-R<DD>",
			format2: "<YYYY>-<MM>-<DD>",
			version: "2025-R1",
			other:   "2025-07-15",
		},
		{
			name:    "2",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY>-WW<DD>",
			version: "2025-07-14",
			other:   "2025-WW15",
		},
		{
			name:    "3",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "2025-07-14",
			other:   "20250724-alpha.2",
		},
		{
			name:    "4",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "2025-07-16",
			other:   "20250724-eksbuild.16002300",
		},
		{
			name:    "5",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "2025-07-16",
			other:   "RELEASE.2025-07-23T14-54-02Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format1, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format2, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.NotEqual(t, nil, rec)
			}()
			_ = ver.CompareOrPanic(other)
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    false,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    false,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    false,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    false,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    false,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    false,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    false,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    false,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			got, err := ver.Equal(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEqualOrPanicNonPanicking(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    false,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    false,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    false,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    false,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    false,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    false,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    false,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    false,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.Equal(t, nil, rec)
			}()
			got := ver.EqualOrPanic(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestEqualOrPanicPanicking(t *testing.T) {
	tests := []struct {
		name    string
		format1 string
		format2 string
		version string
		other   string
	}{
		{
			name:    "1",
			format1: "<YYYY>-R<DD>",
			format2: "<YYYY>-<MM>-<DD>",
			version: "2025-R1",
			other:   "2025-07-15",
		},
		{
			name:    "2",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY>-WW<DD>",
			version: "2025-07-14",
			other:   "2025-WW15",
		},
		{
			name:    "3",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "2025-07-14",
			other:   "20250724-alpha.2",
		},
		{
			name:    "4",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "2025-07-16",
			other:   "20250724-eksbuild.16002300",
		},
		{
			name:    "5",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "2025-07-16",
			other:   "RELEASE.2025-07-23T14-54-02Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format1, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format2, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.NotEqual(t, nil, rec)
			}()
			_ = ver.EqualOrPanic(other)
		})
	}
}

func TestLess(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    true,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    true,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    true,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    false,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    true,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    false,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    false,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    true,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			got, err := ver.Less(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLessOrPanicNonPanicking(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    true,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    true,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    true,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    false,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    true,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    false,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    false,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    true,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.Equal(t, nil, rec)
			}()
			got := ver.LessOrPanic(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLessOrPanicPanicking(t *testing.T) {
	tests := []struct {
		name    string
		format1 string
		format2 string
		version string
		other   string
	}{
		{
			name:    "1",
			format1: "<YYYY>-R<DD>",
			format2: "<YYYY>-<MM>-<DD>",
			version: "2025-R1",
			other:   "2025-07-15",
		},
		{
			name:    "2",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY>-WW<DD>",
			version: "2025-07-14",
			other:   "2025-WW15",
		},
		{
			name:    "3",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "2025-07-14",
			other:   "20250724-alpha.2",
		},
		{
			name:    "4",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "2025-07-16",
			other:   "20250724-eksbuild.16002300",
		},
		{
			name:    "5",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "2025-07-16",
			other:   "RELEASE.2025-07-23T14-54-02Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format1, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format2, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.NotEqual(t, nil, rec)
			}()
			_ = ver.LessOrPanic(other)
		})
	}
}

func TestGreater(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    false,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    true,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    true,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    false,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    false,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    true,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    false,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    true,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    true,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    false,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			got, err := ver.Greater(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGreaterOrPanicNonPanicking(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    false,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    true,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    true,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    false,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    false,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    true,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    false,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    true,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    true,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    false,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.Equal(t, nil, rec)
			}()
			got := ver.GreaterOrPanic(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGreaterOrPanicPanicking(t *testing.T) {
	tests := []struct {
		name    string
		format1 string
		format2 string
		version string
		other   string
	}{
		{
			name:    "1",
			format1: "<YYYY>-R<DD>",
			format2: "<YYYY>-<MM>-<DD>",
			version: "2025-R1",
			other:   "2025-07-15",
		},
		{
			name:    "2",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY>-WW<DD>",
			version: "2025-07-14",
			other:   "2025-WW15",
		},
		{
			name:    "3",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "2025-07-14",
			other:   "20250724-alpha.2",
		},
		{
			name:    "4",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "2025-07-16",
			other:   "20250724-eksbuild.16002300",
		},
		{
			name:    "5",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "2025-07-16",
			other:   "RELEASE.2025-07-23T14-54-02Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format1, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format2, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.NotEqual(t, nil, rec)
			}()
			_ = ver.GreaterOrPanic(other)
		})
	}
}

func TestGreaterOrEqual(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    false,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    true,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    true,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    false,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    false,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    true,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    false,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    true,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    true,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    false,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			got, err := ver.GreaterOrEqual(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGreaterOrEqualOrPanicNonPanicking(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    false,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    true,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    true,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    false,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    false,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    true,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    false,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    false,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    true,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    true,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    false,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.Equal(t, nil, rec)
			}()
			got := ver.GreaterOrEqualOrPanic(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGreaterOrEqualOrPanicPanicking(t *testing.T) {
	tests := []struct {
		name    string
		format1 string
		format2 string
		version string
		other   string
	}{
		{
			name:    "1",
			format1: "<YYYY>-R<DD>",
			format2: "<YYYY>-<MM>-<DD>",
			version: "2025-R1",
			other:   "2025-07-15",
		},
		{
			name:    "2",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY>-WW<DD>",
			version: "2025-07-14",
			other:   "2025-WW15",
		},
		{
			name:    "3",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "2025-07-14",
			other:   "20250724-alpha.2",
		},
		{
			name:    "4",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "2025-07-16",
			other:   "20250724-eksbuild.16002300",
		},
		{
			name:    "5",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "2025-07-16",
			other:   "RELEASE.2025-07-23T14-54-02Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format1, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format2, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.NotEqual(t, nil, rec)
			}()
			_ = ver.GreaterOrEqualOrPanic(other)
		})
	}
}

func TestLessOrEqual(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    true,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    true,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    true,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    false,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    true,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    false,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    false,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    true,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			got, err := ver.LessOrEqual(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLessOrEqualOrPanicNonPanicking(t *testing.T) {
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
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-15",
			want:    true,
		},
		{
			name:    "3",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-14",
			other:   "2025-07-14",
			want:    true,
		},
		{
			name:    "4",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2025-07-14",
			want:    false,
		},
		{
			name:    "5",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "6",
			format:  "<YYYY>-<MM>-<DD>",
			version: "2025-07-16",
			other:   "2020-07-14",
			want:    false,
		},
		{
			name:    "7",
			format:  "<YYYY>.<MM>.<DD>",
			version: "2020.06.16",
			other:   "2020.07.14",
			want:    true,
		},
		{
			name:    "8",
			format:  "<YYYY>-WW<DD>",
			version: "2025-WW14",
			other:   "2025-WW15",
			want:    true,
		},
		{
			name:    "9",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-22T15-54-02Z",
			want:    false,
		},
		{
			name:    "10",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "11",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-03Z",
			want:    true,
		},
		{
			name:    "12",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "13",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T14-54-02Z",
			other:   "RELEASE.2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "14",
			format:  "<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			version: "2025-07-23T14-54-02Z",
			other:   "2025-07-23T15-54-02Z",
			want:    true,
		},
		{
			name:    "15",
			format:  "<YYYY><MM><DD>",
			version: "20260723",
			other:   "20250724",
			want:    false,
		},
		{
			name:    "16",
			format:  "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "20250724-alpha.2",
			other:   "20250724-alpha.1",
			want:    false,
		},
		{
			name:    "17",
			format:  "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "20250724-eksbuild.16002300",
			other:   "20250724-eksbuild.16004300",
			want:    true,
		},
		{
			name:    "18",
			format:  "<YYYY><MM><DD>-foobar.<MODIFIER>",
			version: "20250724-foobar.alpha",
			other:   "20250724-foobar.beta",
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.Equal(t, nil, rec)
			}()
			got := ver.LessOrEqualOrPanic(other)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLessOrEqualOrPanicPanicking(t *testing.T) {
	tests := []struct {
		name    string
		format1 string
		format2 string
		version string
		other   string
	}{
		{
			name:    "1",
			format1: "<YYYY>-R<DD>",
			format2: "<YYYY>-<MM>-<DD>",
			version: "2025-R1",
			other:   "2025-07-15",
		},
		{
			name:    "2",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY>-WW<DD>",
			version: "2025-07-14",
			other:   "2025-WW15",
		},
		{
			name:    "3",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-alpha.<MODIFIER>",
			version: "2025-07-14",
			other:   "20250724-alpha.2",
		},
		{
			name:    "4",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			version: "2025-07-16",
			other:   "20250724-eksbuild.16002300",
		},
		{
			name:    "5",
			format1: "<YYYY>-<MM>-<DD>",
			format2: "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "2025-07-16",
			other:   "RELEASE.2025-07-23T14-54-02Z",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ver, err := calver.NewVersion(tt.format1, tt.version)
			assert.NoError(t, err)
			other, err := calver.NewVersion(tt.format2, tt.other)
			assert.NoError(t, err)
			defer func() {
				rec := recover()
				assert.NotEqual(t, nil, rec)
			}()
			_ = ver.LessOrEqualOrPanic(other)
		})
	}
}