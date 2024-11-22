package oapi

import (
	"encoding/json"
	"net/http"
)

type PingHandler struct {
}

func NewPingHandler() PingHandler {
	return PingHandler{}
}

func (h PingHandler) Ping(w http.ResponseWriter, _ *http.Request) {
	j := json.NewEncoder(w)
	err := j.Encode(Pong{Name: "monitor-app-http", Message: "hello, world!"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
