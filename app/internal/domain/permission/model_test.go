package permission

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	TestPermissionAllowed = Permission{
		ID:           1,
		Name:         "test.allowed",
		DefaultValue: true,
	}

	TestPermissionNotAllowed = Permission{
		ID:           2,
		Name:         "test.not_allowed",
		DefaultValue: false,
	}
)

func TestValueMap_IsTrue(t *testing.T) {
	tests := []struct {
		name     string
		valueMap valueMap
		id       ID
		want     bool
	}{
		{
			name:     "false: empty",
			valueMap: valueMap{},
			id:       TestPermissionAllowed.ID,
			want:     false,
		},
		{
			name: "false: different id",
			valueMap: valueMap{
				TestPermissionAllowed.ID: true,
			},
			id:   0,
			want: false,
		},
		{
			name: "true",
			valueMap: valueMap{
				TestPermissionAllowed.ID: true,
			},
			id:   1,
			want: true,
		},
		{
			name: "false",
			valueMap: valueMap{
				TestPermissionAllowed.ID: false,
			},
			id:   1,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.valueMap.IsTrue(tt.id))
		})
	}
}

func TestValueMap_HasPermission(t *testing.T) {
	tests := []struct {
		name     string
		valueMap valueMap
		perm     Permission
		want     bool
	}{
		{
			name:     "true: empty (default allowed)",
			valueMap: valueMap{},
			perm:     TestPermissionAllowed,
			want:     true,
		},
		{
			name:     "false: empty (default not allowed)",
			valueMap: valueMap{},
			perm:     TestPermissionNotAllowed,
			want:     false,
		},
		{
			name: "true: different id (default allowed)",
			valueMap: valueMap{
				0: false,
			},
			perm: TestPermissionAllowed,
			want: true,
		},
		{
			name: "false: different id (default not allowed)",
			valueMap: valueMap{
				0: true,
			},
			perm: TestPermissionNotAllowed,
			want: false,
		},
		{
			name: "true (default not allowed)",
			valueMap: valueMap{
				TestPermissionNotAllowed.ID: TestPermissionNotAllowed.DefaultValue,
			},
			perm: TestPermissionAllowed,
			want: true,
		},
		{
			name: "false (default allowed)",
			valueMap: valueMap{
				TestPermissionAllowed.ID: TestPermissionAllowed.DefaultValue,
			},
			perm: TestPermissionNotAllowed,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.valueMap.HasPermission(tt.perm))
		})
	}
}

func TestValueMap_HasAnyPermissions(t *testing.T) {
	tests := []struct {
		name     string
		valueMap valueMap
		perms    []Permission
		want     bool
	}{
		// same cases as HasPermission
		{
			name:     "true: empty (default allowed)",
			valueMap: valueMap{},
			perms:    []Permission{TestPermissionAllowed},
			want:     true,
		},
		{
			name:     "false: empty (default not allowed)",
			valueMap: valueMap{},
			perms:    []Permission{TestPermissionNotAllowed},
			want:     false,
		},
		{
			name: "true: different id (default allowed)",
			valueMap: valueMap{
				0: false,
			},
			perms: []Permission{TestPermissionAllowed},
			want:  true,
		},
		{
			name: "false: different id (default not allowed)",
			valueMap: valueMap{
				0: true,
			},
			perms: []Permission{TestPermissionNotAllowed},
			want:  false,
		},
		{
			name: "true (default not allowed)",
			valueMap: valueMap{
				TestPermissionNotAllowed.ID: TestPermissionNotAllowed.DefaultValue,
			},
			perms: []Permission{TestPermissionAllowed},
			want:  true,
		},
		{
			name: "false (default allowed)",
			valueMap: valueMap{
				TestPermissionAllowed.ID: TestPermissionAllowed.DefaultValue,
			},
			perms: []Permission{TestPermissionNotAllowed},
			want:  false,
		},
		// multiple permissions
		{
			name:     "multiple permissions: empty -> default allowed",
			valueMap: valueMap{},
			perms:    []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:     true,
		},
		{
			name: "multiple permissions: all true -> true",
			valueMap: valueMap{
				TestPermissionAllowed.ID:    true,
				TestPermissionNotAllowed.ID: true,
			},
			perms: []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:  true,
		},
		{
			name: "multiple permissions: all false -> false",
			valueMap: valueMap{
				TestPermissionAllowed.ID:    false,
				TestPermissionNotAllowed.ID: false,
			},
			perms: []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.valueMap.HasAnyPermissions(tt.perms...))
		})
	}
}

func TestValueMap_HasAllPermissions(t *testing.T) {
	tests := []struct {
		name     string
		valueMap valueMap
		perms    []Permission
		want     bool
	}{
		// same cases as HasPermission
		{
			name:     "true: empty (default allowed)",
			valueMap: valueMap{},
			perms:    []Permission{TestPermissionAllowed},
			want:     true,
		},
		{
			name:     "false: empty (default not allowed)",
			valueMap: valueMap{},
			perms:    []Permission{TestPermissionNotAllowed},
			want:     false,
		},
		{
			name: "true: different id (default allowed)",
			valueMap: valueMap{
				0: false,
			},
			perms: []Permission{TestPermissionAllowed},
			want:  true,
		},
		{
			name: "false: different id (default not allowed)",
			valueMap: valueMap{
				0: true,
			},
			perms: []Permission{TestPermissionNotAllowed},
			want:  false,
		},
		{
			name: "true (default not allowed)",
			valueMap: valueMap{
				TestPermissionNotAllowed.ID: TestPermissionNotAllowed.DefaultValue,
			},
			perms: []Permission{TestPermissionAllowed},
			want:  true,
		},
		{
			name: "false (default allowed)",
			valueMap: valueMap{
				TestPermissionAllowed.ID: TestPermissionAllowed.DefaultValue,
			},
			perms: []Permission{TestPermissionNotAllowed},
			want:  false,
		},
		// multiple permissions
		{
			name:     "multiple permissions: empty -> default not allowed",
			valueMap: valueMap{},
			perms:    []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:     false,
		},
		{
			name: "multiple permissions: all true -> true",
			valueMap: valueMap{
				TestPermissionAllowed.ID:    true,
				TestPermissionNotAllowed.ID: true,
			},
			perms: []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:  true,
		},
		{
			name: "multiple permissions: all false -> false",
			valueMap: valueMap{
				TestPermissionAllowed.ID:    false,
				TestPermissionNotAllowed.ID: false,
			},
			perms: []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.valueMap.HasAllPermissions(tt.perms...)
			assert.Equal(t, tt.want, got)
		})
	}
}
