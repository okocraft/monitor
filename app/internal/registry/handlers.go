package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/handler/oapi"
)

var handlerSet = wire.NewSet(
	oapi.NewPingHandler,
)
