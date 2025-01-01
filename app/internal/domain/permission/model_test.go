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
		valueMap ValueMap
		id       ID
		want     bool
	}{
		{
			name:     "false: empty",
			valueMap: ValueMap{},
			id:       TestPermissionAllowed.ID,
			want:     false,
		},
		{
			name: "false: different id",
			valueMap: ValueMap{
				TestPermissionAllowed.ID: true,
			},
			id:   0,
			want: false,
		},
		{
			name: "true",
			valueMap: ValueMap{
				TestPermissionAllowed.ID: true,
			},
			id:   1,
			want: true,
		},
		{
			name: "false",
			valueMap: ValueMap{
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
		valueMap ValueMap
		perm     Permission
		want     bool
	}{
		{
			name:     "true: empty (default allowed)",
			valueMap: ValueMap{},
			perm:     TestPermissionAllowed,
			want:     true,
		},
		{
			name:     "false: empty (default not allowed)",
			valueMap: ValueMap{},
			perm:     TestPermissionNotAllowed,
			want:     false,
		},
		{
			name: "true: different id (default allowed)",
			valueMap: ValueMap{
				0: false,
			},
			perm: TestPermissionAllowed,
			want: true,
		},
		{
			name: "false: different id (default not allowed)",
			valueMap: ValueMap{
				0: true,
			},
			perm: TestPermissionNotAllowed,
			want: false,
		},
		{
			name: "true (default not allowed)",
			valueMap: ValueMap{
				TestPermissionNotAllowed.ID: TestPermissionNotAllowed.DefaultValue,
			},
			perm: TestPermissionAllowed,
			want: true,
		},
		{
			name: "false (default allowed)",
			valueMap: ValueMap{
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
		valueMap ValueMap
		perms    []Permission
		want     bool
	}{
		// same cases as HasPermission
		{
			name:     "true: empty (default allowed)",
			valueMap: ValueMap{},
			perms:    []Permission{TestPermissionAllowed},
			want:     true,
		},
		{
			name:     "false: empty (default not allowed)",
			valueMap: ValueMap{},
			perms:    []Permission{TestPermissionNotAllowed},
			want:     false,
		},
		{
			name: "true: different id (default allowed)",
			valueMap: ValueMap{
				0: false,
			},
			perms: []Permission{TestPermissionAllowed},
			want:  true,
		},
		{
			name: "false: different id (default not allowed)",
			valueMap: ValueMap{
				0: true,
			},
			perms: []Permission{TestPermissionNotAllowed},
			want:  false,
		},
		{
			name: "true (default not allowed)",
			valueMap: ValueMap{
				TestPermissionNotAllowed.ID: TestPermissionNotAllowed.DefaultValue,
			},
			perms: []Permission{TestPermissionAllowed},
			want:  true,
		},
		{
			name: "false (default allowed)",
			valueMap: ValueMap{
				TestPermissionAllowed.ID: TestPermissionAllowed.DefaultValue,
			},
			perms: []Permission{TestPermissionNotAllowed},
			want:  false,
		},
		// multiple permissions
		{
			name:     "multiple permissions: empty -> default allowed",
			valueMap: ValueMap{},
			perms:    []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:     true,
		},
		{
			name: "multiple permissions: all true -> true",
			valueMap: ValueMap{
				TestPermissionAllowed.ID:    true,
				TestPermissionNotAllowed.ID: true,
			},
			perms: []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:  true,
		},
		{
			name: "multiple permissions: all false -> false",
			valueMap: ValueMap{
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
		valueMap ValueMap
		perms    []Permission
		want     bool
	}{
		// same cases as HasPermission
		{
			name:     "true: empty (default allowed)",
			valueMap: ValueMap{},
			perms:    []Permission{TestPermissionAllowed},
			want:     true,
		},
		{
			name:     "false: empty (default not allowed)",
			valueMap: ValueMap{},
			perms:    []Permission{TestPermissionNotAllowed},
			want:     false,
		},
		{
			name: "true: different id (default allowed)",
			valueMap: ValueMap{
				0: false,
			},
			perms: []Permission{TestPermissionAllowed},
			want:  true,
		},
		{
			name: "false: different id (default not allowed)",
			valueMap: ValueMap{
				0: true,
			},
			perms: []Permission{TestPermissionNotAllowed},
			want:  false,
		},
		{
			name: "true (default not allowed)",
			valueMap: ValueMap{
				TestPermissionNotAllowed.ID: TestPermissionNotAllowed.DefaultValue,
			},
			perms: []Permission{TestPermissionAllowed},
			want:  true,
		},
		{
			name: "false (default allowed)",
			valueMap: ValueMap{
				TestPermissionAllowed.ID: TestPermissionAllowed.DefaultValue,
			},
			perms: []Permission{TestPermissionNotAllowed},
			want:  false,
		},
		// multiple permissions
		{
			name:     "multiple permissions: empty -> default not allowed",
			valueMap: ValueMap{},
			perms:    []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:     false,
		},
		{
			name: "multiple permissions: all true -> true",
			valueMap: ValueMap{
				TestPermissionAllowed.ID:    true,
				TestPermissionNotAllowed.ID: true,
			},
			perms: []Permission{TestPermissionAllowed, TestPermissionNotAllowed},
			want:  true,
		},
		{
			name: "multiple permissions: all false -> false",
			valueMap: ValueMap{
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
