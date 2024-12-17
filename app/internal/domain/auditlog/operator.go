package auditlog

import (
	"github.com/okocraft/monitor/internal/domain/user"
	"net"
)

type Operator struct {
	UserID user.ID
	Name   string
	IP     net.IP
}

type OperatorID int64
