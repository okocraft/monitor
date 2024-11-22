//go:build wireinject

package registry

import (
	"github.com/google/wire"
	"github.com/okocraft/monitor/internal/handler/oapi"
)

//go:generate go run github.com/google/wire/cmd/wire@v0.6.0

func NewHTTPHandler() (oapi.HTTPHandler, error) {
	wire.Build(
		handlerSet,
		oapi.NewHTTPHandler,
	)
	return oapi.HTTPHandler{}, nil
}
