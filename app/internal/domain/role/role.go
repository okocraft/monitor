package role

import (
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/handler/oapi"
	"time"
)

type ID = int32

type Role struct {
	ID        ID
	UUID      uuid.UUID
	Name      string
	Priority  int32
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (role Role) ToResponse() oapi.Role {
	return oapi.Role{
		Id:        role.UUID,
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
	}
}

var defaultRole = Role{
	ID:        0,
	UUID:      uuid.Nil,
	Name:      "Default",
	Priority:  0,
	CreatedAt: time.Time{},
	UpdatedAt: time.Time{},
}

func DefaultRole() Role {
	return defaultRole
}
