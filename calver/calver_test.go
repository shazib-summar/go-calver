package calver_test

import (
	"testing"

	"github.com/shazib-summar/go-calver/calver"
	"github.com/stretchr/testify/assert"
)

func TestNewCalVer(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		version string
		wantErr bool
	}{
		{name: "1", format: "<YYYY>-R<DD>", version: "2025-R1", wantErr: false},
		{name: "2", format: "<YYYY>-R<0D>", version: "2025-R01", wantErr: false},
		{name: "3", format: "<YYYY>-R<DD>", version: "2025-R01", wantErr: false},
		{name: "4", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", wantErr: false},
		{name: "5", format: "<YYYY>.<MM>.<DD>", version: "2025.07.14", wantErr: false},
		{name: "6", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", wantErr: false},
		{name: "7", format: "<YY>-<MM>-<DD>", version: "2025-07-14", wantErr: true},
		{name: "8", format: "<0Y>.<0M>.<DD>", version: "18.04.6", wantErr: false},
		{name: "9", format: "<0Y>.<0M>.<DD>", version: "22.04.6", wantErr: false},
		{name: "10", format: "<0Y>.<0M>.<DD>", version: "22.4.6", wantErr: true},
		{name: "11", format: "<YYYY>-WW<DD>", version: "2025-WW14", wantErr: false},
		{name: "12", format: "<YYYY>-WW<0D>", version: "2025-WW04", wantErr: false},
		{name: "13", format: "<YYYY>-<MINOR>", version: "2025-14", wantErr: false},
		{name: "14", format: "<YYYY>-<MICRO>", version: "2025-14-12", wantErr: true},
		{name: "15", format: "<MAJOR>-<MINOR>-<MICRO>", version: "2025-14-12", wantErr: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := calver.NewCalVer(test.format, test.version)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestCalVerString(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		version string
		want    string
	}{
		{name: "1", format: "<YYYY>-R<DD>", version: "2025-R1", want: "2025-R1"},
		{name: "2", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", want: "2025-07-14"},
		{name: "3", format: "<YYYY>.<MM>.<DD>", version: "2025.07.14", want: "2025.07.14"},
		{name: "4", format: "<YY>-<MM>-<DD>", version: "25-07-14", want: "25-07-14"},
		{name: "5", format: "<0Y>.<0M>.<DD>", version: "18.04.6", want: "18.04.6"},
		{name: "6", format: "<YYYY>-WW<DD>", version: "2025-WW14", want: "2025-WW14"},
		{name: "7", format: "<YYYY>-WW<0D>", version: "2025-WW04", want: "2025-WW04"},
		{
			name:    "8",
			format:  "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			version: "RELEASE.2025-07-23T15-54-02Z",
			want:    "RELEASE.2025-07-23T15-54-02Z",
		},
		{name: "9", format: "<MAJOR>-WW<MINOR>", version: "2025-WW04", want: "2025-WW04"},
		{name: "10", format: "<MAJOR>-<YYY>-<MICRO>", version: "2025-<YYY>-12", want: "2025-<YYY>-12"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calver, err := calver.NewCalVer(test.format, test.version)
			assert.NoError(t, err)
			assert.Equal(t, test.want, calver.String())
		})
	}
}

func TestCalVerSeries(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		version string
		level   string
		want    string
	}{
		{name: "1", format: "<YYYY>-R<DD>", version: "2025-R1", level: "major", want: "2025"},
		{name: "2", format: "<YYYY>-R<DD>", version: "2025-R1", level: "minor", want: "2025-R1"},
		{name: "3", format: "<YYYY>-R<DD>", version: "2025-R1", level: "micro", want: "2025-R1"},
		{name: "4", format: "<YYYY>-R<DD>", version: "2025-R1", level: "modifier", want: "2025-R1"},
		{name: "5", format: "<YYYY>-R<DD>", version: "2025-R1", level: "", want: "2025-R1"},
		{name: "6", format: "<YYYY>-R<DD>", version: "2025-R1", level: "invalid", want: "2025-R1"},
		{name: "7", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", level: "major", want: "2025"},
		{name: "8", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", level: "minor", want: "2025-07"},
		{name: "9", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", level: "micro", want: "2025-07-14"},
		{name: "10", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", level: "modifier", want: "2025-07-14"},
		{name: "11", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", level: "", want: "2025-07-14"},
		{name: "12", format: "<YYYY>-<MM>-<DD>", version: "2025-07-14", level: "invalid", want: "2025-07-14"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calver, err := calver.NewCalVer(test.format, test.version)
			assert.NoError(t, err)
			assert.Equal(t, test.want, calver.Series(test.level))
		})
	}
}
