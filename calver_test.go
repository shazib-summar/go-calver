package calver_test

import (
	"testing"

	"github.com/shazib-summar/go-calver"
	"github.com/stretchr/testify/assert"
)

func TestNewVersion(t *testing.T) {
	tests := []struct {
		name    string
		formats []string
		version string
		wantErr bool
	}{
		{name: "1", formats: []string{"<YYYY>-R<DD>"}, version: "2025-R1", wantErr: false},
		{name: "2", formats: []string{"<YYYY>-R<0D>"}, version: "2025-R01", wantErr: false},
		{name: "3", formats: []string{"<YYYY>-R<DD>"}, version: "2025-R01", wantErr: false},
		{name: "4", formats: []string{"<YYYY>-<MM>-<DD>"}, version: "2025-07-14", wantErr: false},
		{name: "5", formats: []string{"<YYYY>.<MM>.<DD>"}, version: "2025.07.14", wantErr: false},
		{name: "6", formats: []string{"<YYYY>-<MM>-<DD>"}, version: "2025-07-14", wantErr: false},
		{name: "7", formats: []string{"<YY>-<MM>-<DD>"}, version: "2025-07-14", wantErr: true},
		{name: "8", formats: []string{"<0Y>.<0M>.<DD>"}, version: "18.04.6", wantErr: false},
		{name: "9", formats: []string{"<0Y>.<0M>.<DD>"}, version: "22.04.6", wantErr: false},
		{name: "10", formats: []string{"<0Y>.<0M>.<DD>"}, version: "22.4.6", wantErr: true},
		{name: "11", formats: []string{"<YYYY>-WW<DD>"}, version: "2025-WW14", wantErr: false},
		{name: "12", formats: []string{"<YYYY>-WW<0D>"}, version: "2025-WW04", wantErr: false},
		{name: "13", formats: []string{"<YYYY>-<MINOR>"}, version: "2025-14", wantErr: false},
		{name: "14", formats: []string{"<YYYY>-<MICRO>"}, version: "2025-14-12", wantErr: true},
		{name: "15", formats: []string{"<MAJOR>-<MINOR>-<MICRO>"}, version: "2025-14-12", wantErr: false},
		{name: "16", formats: []string{"v<MAJOR>-<MINOR>-<MICRO>"}, version: "v2025-14-12", wantErr: false},
		{name: "17", formats: []string{"v<MAJOR>-<MINOR>-<MICRO>"}, version: "2025-14-12", wantErr: true},
		{
			name: "18",
			formats: []string{
				"<YYYY>-<MM>-<DD>",
				"<YYYY>.<MM>.<DD>",
			},
			version: "2025.07.14",
			wantErr: false,
		},
		{
			name:    "19",
			formats: []string{"<YYYY>.<MM>.<DD>", "<YYYY>/<MM>/<DD>"},
			version: "2025-07-14",
			wantErr: true,
		},
		{
			name:    "20",
			formats: []string{"<YYYY>.<MM>.<DD>", "<YYYY>/<MM>/<DD>"},
			version: "2025.07",
			wantErr: true, // missing day
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := calver.NewVersion(test.version, test.formats...)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestVersionString(t *testing.T) {
	tests := []struct {
		name    string
		formats []string
		version string
		want    string
	}{
		{name: "1", formats: []string{"<YYYY>-R<DD>"}, version: "2025-R1", want: "2025-R1"},
		{name: "2", formats: []string{"<YYYY>-<MM>-<DD>"}, version: "2025-07-14", want: "2025-07-14"},
		{name: "3", formats: []string{"<YYYY>.<MM>.<DD>"}, version: "2025.07.14", want: "2025.07.14"},
		{name: "4", formats: []string{"<YY>-<MM>-<DD>"}, version: "25-07-14", want: "25-07-14"},
		{name: "5", formats: []string{"<0Y>.<0M>.<DD>"}, version: "18.04.6", want: "18.04.6"},
		{name: "6", formats: []string{"<YYYY>-WW<DD>"}, version: "2025-WW14", want: "2025-WW14"},
		{name: "7", formats: []string{"<YYYY>-WW<0D>"}, version: "2025-WW04", want: "2025-WW04"},
		{
			name:    "8",
			formats: []string{"RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z"},
			version: "RELEASE.2025-07-23T15-54-02Z",
			want:    "RELEASE.2025-07-23T15-54-02Z",
		},
		{name: "9", formats: []string{"<MAJOR>-WW<MINOR>"}, version: "2025-WW04", want: "2025-WW04"},
		{name: "10", formats: []string{"<MAJOR>-<YYY>-<MICRO>"}, version: "2025-<YYY>-12", want: "2025-<YYY>-12"},
		{
			name:    "11",
			formats: []string{"v<YYYY><0M><0D>"},
			version: "v20250723",
			want:    "v20250723",
		},
		{
			name:    "12",
			formats: []string{"<YYYY>.<MM>.<DD>", "<YYYY>/<MM>/<DD>"},
			version: "2025.07.14",
			want:    "2025.07.14",
		},
		{
			name:    "13",
			formats: []string{"<YYYY>.<MM>.<DD>", "<YYYY>/<MM>/<DD>", "<YYYY>-<MM>-<DD>"},
			version: "2025-07-14",
			want:    "2025-07-14",
		},
		{
			name: "14",
			formats: []string{
				"<YYYY>",
				"<YYYY>-<MM>",
				"<YYYY>-<MM>-<DD>",
			},
			version: "2025-07-14",
			want:    "2025-07-14",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calver, err := calver.NewVersion(test.version, test.formats...)
			assert.NoError(t, err)
			assert.Equal(t, test.want, calver.String())
		})
	}
}

func TestVersionSeries(t *testing.T) {
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
		{name: "12", format: "v<YYYY>-<MM>-<DD>", version: "v2025-07-14", level: "invalid", want: "v2025-07-14"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			calver, err := calver.NewVersion(test.version, test.format)
			assert.NoError(t, err)
			assert.Equal(t, test.want, calver.Series(test.level))
		})
	}
}
