package auditlog

import (
	"net"

	"github.com/okocraft/monitor/internal/domain/user"
)

type Operator struct {
	UserID user.ID
	Name   string
	IP     net.IP
}

type OperatorID int64
