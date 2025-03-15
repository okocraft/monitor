package role

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/handler/oapi"
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
		Priority:  role.Priority,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
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

type Roles []Role

func (roles Roles) ToResponse() []oapi.Role {
	res := make([]oapi.Role, len(roles))
	for i, role := range roles {
		res[i] = role.ToResponse()
	}
	return res
}
