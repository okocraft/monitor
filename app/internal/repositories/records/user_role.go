package records

import (
	"github.com/okocraft/monitor/internal/repositories/queries"
	"time"
)

var UserRole1 = queries.UsersRole{
	UserID:    User1.ID,
	RoleID:    Role1.ID,
	UpdatedAt: time.Date(2025, 2, 26, 0, 0, 0, 0, time.UTC),
}
