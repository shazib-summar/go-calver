package calver_test

import (
	"testing"

	"github.com/shazib-summar/go-calver"
	"github.com/stretchr/testify/assert"
)

func getVersion(version string, formats ...string) *calver.Version {
	ver, _ := calver.NewVersion(version, formats...)
	return ver
}

func TestCompare(t *testing.T) {
	tests := []struct {
		name    string
		version *calver.Version
		other   *calver.Version
		want    int
	}{
		{
			name: "1",
			version: getVersion(
				"2025-R1",
				"<YYYY>-R<DD>",
			),
			other: getVersion(
				"2025-R2",
				"<YYYY>-R<DD>",
			),
			want: -1,
		},
		{
			name: "2",
			version: getVersion(
				"2025-07-14",
				"<YYYY>-<MM>-<DD>",
			),
			other: getVersion(
				"2025-07-15",
				"<YYYY>-<MM>-<DD>",
			),
			want: -1,
		},
		{
			name: "3",
			version: getVersion(
				"2025-07-14",
				"<YYYY>-<MM>-<DD>",
			),
			other: getVersion(
				"2025-07-14",
				"<YYYY>-<MM>-<DD>",
			),
			want: 0,
		},
		{
			name: "4",
			version: getVersion(
				"2025-07-16",
				"<YYYY>-<MM>-<DD>",
			),
			other: getVersion(
				"2025-07-14",
				"<YYYY>-<MM>-<DD>",
			),
			want: 1,
		},
		{
			name:    "5",
			version: getVersion("2025-07-16", "<YYYY>-<MM>-<DD>"),
			other:   getVersion("2020-07-14", "<YYYY>-<MM>-<DD>"),
			want:    1,
		},
		{
			name: "6",
			version: getVersion(
				"2025-07-16",
				"<YYYY>-<MM>-<DD>",
			),
			other: getVersion(
				"2020-07-14",
				"<YYYY>-<MM>-<DD>",
			),
			want: 1,
		},
		{
			name: "7",
			version: getVersion(
				"2020.06.16",
				"<YYYY>.<MM>.<DD>",
			),
			other: getVersion(
				"2020.07.14",
				"<YYYY>.<MM>.<DD>",
			),
			want: -1,
		},
		{
			name: "8",
			version: getVersion(
				"2025-WW14",
				"<YYYY>-WW<DD>",
			),
			other: getVersion(
				"2025-WW15",
				"<YYYY>-WW<DD>",
			),
			want: -1,
		},
		{
			name: "9",
			version: getVersion(
				"RELEASE.2025-07-23T15-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			other: getVersion(
				"RELEASE.2025-07-22T15-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			want: 1,
		},
		{
			name: "10",
			version: getVersion(
				"RELEASE.2025-07-23T15-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			other: getVersion(
				"RELEASE.2025-07-23T15-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			want: 0,
		},
		{
			name: "11",
			version: getVersion(
				"RELEASE.2025-07-23T15-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			other: getVersion(
				"RELEASE.2025-07-23T15-54-03Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			want: -1,
		},
		{
			name: "12",
			version: getVersion(
				"RELEASE.2025-07-23T14-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			other: getVersion(
				"RELEASE.2025-07-23T15-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			want: -1,
		},
		{
			name: "13",
			version: getVersion(
				"RELEASE.2025-07-23T14-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			other: getVersion(
				"RELEASE.2025-07-23T15-54-02Z",
				"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			),
			want: -1,
		},
		{
			name: "14",
			version: getVersion(
				"2025-07-23T14-54-02Z",
				"<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			),
			other: getVersion(
				"2025-07-23T15-54-02Z",
				"<MAJOR>-<MINOR>-<MICRO>T<MODIFIER>Z",
			),
			want: -1,
		},
		{
			name: "15",
			version: getVersion(
				"20260723",
				"<YYYY><MM><DD>",
			),
			other: getVersion(
				"20250724",
				"<YYYY><MM><DD>",
			),
			want: 1,
		},
		{
			name: "16",
			version: getVersion(
				"20250724-alpha.2",
				"<YYYY><MM><DD>-alpha.<MODIFIER>",
			),
			other: getVersion(
				"20250724-alpha.1",
				"<YYYY><MM><DD>-alpha.<MODIFIER>",
			),
			want: 1,
		},
		{
			name: "17",
			version: getVersion(
				"20250724-eksbuild.16002300",
				"<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			),
			other: getVersion(
				"20250724-eksbuild.16004300",
				"<YYYY><MM><DD>-eksbuild.<MODIFIER>",
			),
			want: -1,
		},
		{
			name: "18",
			version: getVersion(
				"20250724-foobar.alpha",
				"<YYYY><MM><DD>-foobar.<MODIFIER>",
			),
			other: getVersion(
				"20250724-foobar.beta",
				"<YYYY><MM><DD>-foobar.<MODIFIER>",
			),
			want: -1,
		},
		{
			name: "19",
			version: getVersion(
				"2024-04-01",
				"<YYYY>-<MM>-<DD>",
			),
			other: getVersion(
				"2024.04.01",
				"<YYYY>.<MM>.<DD>",
			),
			want: 0,
		},
		{
			name: "20",
			version: getVersion(
				"2024/04/01",
				"<YYYY>/<MM>/<DD>",
			),
			other: getVersion(
				"2024|04",
				"<YYYY>|<MM>",
			),
			want: 1,
		},
		{
			name: "21",
			version: getVersion(
				"2024/04/01",
				"<YYYY>/<MM>/<DD>",
			),
			other: getVersion(
				"2024-R12",
				"<YYYY>-R<MINOR>",
			),
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.version.Compare(tt.other)
			assert.Equal(t, tt.want, got)
		})
	}
}
