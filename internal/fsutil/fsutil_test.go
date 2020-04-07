package fsutil

import (
	"os"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParsePath(t *testing.T) {

	tests := []struct {
		name    string
		s       string
		want    string
		wantErr bool
	}{
		{
			name: "Success Case: path with ~/",
			s:    "~/dir",
			want: func() string {
				u, err := user.Current()
				require.NoError(t, err)
				return u.HomeDir + "/dir"
			}(),
		},
		{
			name: "Success Case: path with ./",
			s:    "./dir",
			want: func() string {
				p, err := os.Getwd()
				require.NoError(t, err)
				return p + "/dir"
			}(),
		},
		{
			name: "Success Case: path with ..",
			s:    "../dir",
			want: func() string {
				p, err := os.Getwd()
				require.NoError(t, err)
				return filepath.Clean(p + "/../dir")
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePath(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParsePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
