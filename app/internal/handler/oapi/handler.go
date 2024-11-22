package oapi

var _ ServerInterface = (*HTTPHandler)(nil)

type HTTPHandler struct {
	PingHandler
}

func NewHTTPHandler(
	pingHandler PingHandler,
) HTTPHandler {
	return HTTPHandler{
		PingHandler: pingHandler,
	}
}
