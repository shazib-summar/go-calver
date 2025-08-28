package calver_test

import (
	"sort"
	"testing"

	"github.com/shazib-summar/go-calver/calver"
	"github.com/stretchr/testify/assert"
)

func TestNewCollection(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		versions []string
		wantErr  bool
	}{
		{
			name:     "1",
			format:   "<YYYY>-R<DD>",
			versions: []string{"2025-R1", "2025-R2", "2025-R3", "2025-R20", "2025-R35"},
			wantErr:  false,
		},
		{
			name:     "2",
			format:   "<YYYY>-<MM>-<DD>",
			versions: []string{"2025-07-14", "2025-07-15", "2025-07-16", "2025-07-17", "2025-07-18"},
			wantErr:  false,
		},
		{
			name:     "3",
			format:   "<YYYY>.<MM>.<DD>",
			versions: []string{"2025.07.14", "2025.07.15", "2025.07.16", "2025.07.17", "2025.07.18"},
			wantErr:  false,
		},
		{
			name:     "4",
			format:   "<YYYY>-WW<DD>",
			versions: []string{"2025-WW14", "2025-WW15", "2025-WW16", "2025-WW17", "2025-WW18"},
			wantErr:  false,
		},
		{
			name:   "5",
			format: "RELEASE.<YYYY>-<0M>-<0D>T<MODIFIER>Z",
			versions: []string{
				"RELEASE.2025-07-23T15-54-02Z",
				"RELEASE.2025-07-23T15-54-03Z",
				"RELEASE.2025-07-23T15-54-04Z",
				"RELEASE.2025-07-23T15-54-05Z",
				"RELEASE.2025-07-23T15-54-06Z",
			},
			wantErr: false,
		},
		{
			name:     "6",
			format:   "<YYYY><MM><DD>",
			versions: []string{"20250723", "20250724", "20250725", "20250726", "20250727"},
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection, err := calver.NewCollection(tt.format, tt.versions...)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Greater(t, collection.Len(), 0)
			}
		})
	}
}

func TestSortCollection(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		versions []string
		want     []string
	}{
		{
			name:     "1",
			format:   "<YYYY>-R<DD>",
			versions: []string{"2025-R35", "2025-R3", "2025-R2", "2025-R1", "2025-R20"},
			want:     []string{"2025-R1", "2025-R2", "2025-R3", "2025-R20", "2025-R35"},
		},
		{
			name:     "2",
			format:   "<YYYY>-<MM>-<DD>",
			versions: []string{"2025-07-18", "2025-07-17", "2025-07-16", "2025-07-15", "2025-07-14"},
			want:     []string{"2025-07-14", "2025-07-15", "2025-07-16", "2025-07-17", "2025-07-18"},
		},
		{
			name:     "3",
			format:   "<YYYY>.<MM>.<DD>",
			versions: []string{"2025.07.18", "2025.07.17", "2025.07.16", "2025.07.15", "2025.07.14"},
			want:     []string{"2025.07.14", "2025.07.15", "2025.07.16", "2025.07.17", "2025.07.18"},
		},
		{
			name:     "4",
			format:   "<YY>-<MM>-<DD>",
			versions: []string{"25-07-18", "25-07-21", "25-07-16", "25-07-15", "25-07-14"},
			want:     []string{"25-07-14", "25-07-15", "25-07-16", "25-07-18", "25-07-21"},
		},
		{
			name:     "5",
			format:   "<YYYY>-Rel<MINOR>/<MICRO>",
			versions: []string{"2025-Rel07/14", "2025-Rel07/15", "2025-Rel07/16", "2025-Rel07/17", "2025-Rel07/18"},
			want:     []string{"2025-Rel07/14", "2025-Rel07/15", "2025-Rel07/16", "2025-Rel07/17", "2025-Rel07/18"},
		},
		{
			name:     "6",
			format:   "<YYYY><MM><DD>",
			versions: []string{"20240811", "20240711", "20250711", "20251130", "20250826"},
			want:     []string{"20240711", "20240811", "20250711", "20250826", "20251130"},
		},
		{
			name:     "7",
			format:   "<YYYY><MM><DD>-alpha.<MODIFIER>",
			versions: []string{"20220721-alpha.1", "20210922-alpha.2", "20210318-alpha.3", "20260121-alpha.4", "20210721-alpha.5"},
			want:     []string{"20210318-alpha.3", "20210721-alpha.5", "20210922-alpha.2", "20220721-alpha.1", "20260121-alpha.4"},
		},
		{
			name:   "8",
			format: "<YYYY>/<MM>/<DD>-eksbuild.<MODIFIER>",
			versions: []string{
				"2025/07/24-eksbuild.16002300",
				"2025/07/24-eksbuild.16004300",
				"2025/07/24-eksbuild.16001300",
			},
			want: []string{
				"2025/07/24-eksbuild.16001300",
				"2025/07/24-eksbuild.16002300",
				"2025/07/24-eksbuild.16004300",
			},
		},
		{
			name:   "9",
			format: "<YYYY><MM><DD>-foobar.<MODIFIER>",
			versions: []string{
				"20250724-foobar.alpha",
				"20250724-foobar.beta",
				"20250724-foobar.gamma",
				"20250724-foobar.delta",
				"20250724-foobar.epsilon",
			},
			want: []string{
				"20250724-foobar.alpha",
				"20250724-foobar.beta",
				"20250724-foobar.delta",
				"20250724-foobar.epsilon",
				"20250724-foobar.gamma",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collection, err := calver.NewCollection(tt.format, tt.versions...)
			assert.NoError(t, err)
			sort.Sort(collection)
			for i, v := range collection {
				assert.Equal(t, tt.want[i], v.String())
			}
		})
	}
}
