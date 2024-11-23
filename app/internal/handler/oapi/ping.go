package oapi

import (
	"github.com/Siroshun09/logs"
	"github.com/okocraft/monitor/lib/httplib"
	"net/http"
)

type PingHandler struct {
}

func NewPingHandler() PingHandler {
	return PingHandler{}
}

func (h PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	if err := httplib.RenderOK(ctx, w, Pong{Name: "monitor-app-http", Message: "hello, world!"}); err != nil {
		logs.Error(ctx, err)
	}
}
