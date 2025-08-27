package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncWithPadding(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{
			name: "1",
			in:   "1",
			want: "2",
		},
		{
			name: "2",
			in:   "01",
			want: "02",
		},
		{
			name: "3",
			in:   "",
			want: "",
		},
		{
			name: "4",
			in:   "09",
			want: "10",
		},
		{
			name: "5",
			in:   "099",
			want: "100",
		},
		{
			name: "6",
			in:   "0999",
			want: "1000",
		},
		{
			name: "7",
			in:   "199",
			want: "200",
		},
		{
			name: "8",
			in:   "999",
			want: "1000",
		},
		{
			name:    "9",
			in:      "abc",
			want:    "",
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := IncWithPadding(test.in)
			if test.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.want, got)
			}
		})
	}
}
