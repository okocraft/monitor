package httplib

import (
	"context"
	"encoding/json"
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/lib/ctxlib"
	"net/http"
)

type Response interface {
	WriteTo(w http.ResponseWriter) error
}

func RenderOK(ctx context.Context, w http.ResponseWriter, res any) error {
	if err := render(ctx, w, http.StatusOK, res); err != nil {
		return err
	}
	return nil
}

func RenderCreated(ctx context.Context, w http.ResponseWriter, res any) error {
	if err := render(ctx, w, http.StatusCreated, res); err != nil {
		return err
	}
	return nil
}

func render(ctx context.Context, w http.ResponseWriter, statusCode int, res any) error {
	if err := write(w, statusCode, res); err != nil {
		return err
	}

	response := ctxlib.GetHTTPResponse(ctx)
	response.StatusCode = statusCode
	response.WriteHandlerInfo(3) // render -> RenderXXX -> caller

	return nil
}

func write(w http.ResponseWriter, statusCode int, res any) error {
	w.WriteHeader(statusCode)

	if res == nil {
		return nil
	}

	str, ok := res.(string)
	if ok {
		w.Header().Set("Content-Type", "text/plain")
		if _, err := w.Write([]byte(str)); err != nil {
			return serrors.WithStackTrace(err)
		}
		return nil
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		return serrors.WithStackTrace(err)
	}
	return nil
}
