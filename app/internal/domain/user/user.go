package user

import (
	"strconv"
	"time"

	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/handler/oapi"

	"github.com/gofrs/uuid/v5"
)

type ID int32

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

type User struct {
	ID         ID
	UUID       uuid.UUID
	Nickname   string
	LastAccess time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserWithRole struct {
	User
	Role role.Role
}

func (u UserWithRole) ToResponse() oapi.User {
	return oapi.User{
		Id:         u.UUID,
		Nickname:   u.Nickname,
		LastAccess: u.LastAccess,
		Role:       u.Role.ToResponse(),
		CreatedAt:  u.CreatedAt,
		UpdatedAt:  u.UpdatedAt,
	}
}
