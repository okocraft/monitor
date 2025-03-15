package testutils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetProjectRoot(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "success",
			want:    "../../",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantDir, err := filepath.Abs(tt.want)
			require.NoError(t, err)
			got, err := GetProjectRoot()
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, wantDir, got)
		})
	}
}
