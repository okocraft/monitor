package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/handler/oapi/auth"
	"github.com/okocraft/monitor/internal/handler/oapi/me"
)

var handlerSet = wire.NewSet(
	auth.NewAuthHandler,
	auth.NewGoogleAuthHandler,
	me.NewMeHandler,
)
