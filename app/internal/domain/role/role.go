package role

import (
	"github.com/okocraft/monitor/internal/handler/oapi"
	"time"
)

type ID = int32

type Role struct {
	ID        ID
	Name      string
	Priority  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (role Role) ToResponse() oapi.Role {
	return oapi.Role{
		Id:        role.ID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
	}
}

var defaultRole = Role{
	ID:        0,
	Name:      "Default",
	Priority:  0,
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
}

func DefaultRole() Role {
	return defaultRole
}
