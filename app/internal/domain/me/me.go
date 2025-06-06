package me

import (
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/handler/oapi"
)

type Me struct {
	UUID     uuid.UUID
	Nickname string
	RoleUUID uuid.UUID
	RoleName string
}

func (m Me) ToResponse() oapi.Me {
	return oapi.Me{
		Uuid:     m.UUID,
		Nickname: m.Nickname,
		RoleUuid: m.RoleUUID,
		RoleName: m.RoleName,
	}
}
