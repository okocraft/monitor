package records

import (
	"time"

	"github.com/okocraft/monitor/internal/repositories/queries"
)

var UserRole1 = queries.UsersRole{
	UserID:    User1.ID,
	RoleID:    Role1.ID,
	UpdatedAt: time.Date(2025, 2, 26, 0, 0, 0, 0, time.UTC),
}

var UserRole2 = queries.UsersRole{
	UserID:    User2.ID,
	RoleID:    Role2.ID,
	UpdatedAt: time.Date(2025, 2, 27, 0, 0, 0, 0, time.UTC),
}

var UserRole3 = queries.UsersRole{
	UserID:    User3.ID,
	RoleID:    Role3.ID,
	UpdatedAt: time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC),
}
