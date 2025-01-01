package records

import (
	"github.com/okocraft/monitor/internal/domain/permission"
)

var (
	TestPermissionAllowed = permission.Permission{
		ID:           1,
		Name:         "test.allowed",
		DefaultValue: true,
	}

	TestPermissionNotAllowed = permission.Permission{
		ID:           2,
		Name:         "test.not_allowed",
		DefaultValue: false,
	}
)
