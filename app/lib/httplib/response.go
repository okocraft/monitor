package httplib

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Siroshun09/logs"
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/lib/ctxlib"
)

type Response interface {
	WriteTo(w http.ResponseWriter) error
}

func RenderOK(ctx context.Context, w http.ResponseWriter, res any) {
	if err := render(ctx, w, http.StatusOK, res); err != nil {
		logs.Error(ctx, err)
	}
}

func RenderCreated(ctx context.Context, w http.ResponseWriter, res any) {
	if err := render(ctx, w, http.StatusCreated, res); err != nil {
		logs.Error(ctx, err)
	}
}

func RenderNoContent(ctx context.Context, w http.ResponseWriter) {
	if err := render(ctx, w, http.StatusNoContent, nil); err != nil {
		logs.Error(ctx, err)
	}
}

func RenderRedirect(ctx context.Context, w http.ResponseWriter, r *http.Request, url string) {
	response := ctxlib.GetHTTPResponse(ctx)
	response.StatusCode = http.StatusTemporaryRedirect
	response.WriteHandlerInfo(2) // RenderRedirect -> caller

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func RenderBadRequest(ctx context.Context, w http.ResponseWriter, cause error) {
	if err := renderError(ctx, w, http.StatusBadRequest, cause); err != nil {
		logs.Error(ctx, err)
	}
}

func RenderUnauthorized(ctx context.Context, w http.ResponseWriter, cause error) {
	if err := renderError(ctx, w, http.StatusUnauthorized, cause); err != nil {
		logs.Error(ctx, err)
	}
}

func RenderForbidden(ctx context.Context, w http.ResponseWriter, cause error) {
	if err := renderError(ctx, w, http.StatusForbidden, cause); err != nil {
		logs.Error(ctx, err)
	}
}

func RenderNotFound(ctx context.Context, w http.ResponseWriter, cause error) {
	if err := renderError(ctx, w, http.StatusNotFound, cause); err != nil {
		logs.Error(ctx, err)
	}
}

func RenderNotAcceptable(ctx context.Context, w http.ResponseWriter, cause error) {
	if err := renderError(ctx, w, http.StatusNotAcceptable, cause); err != nil {
		logs.Error(ctx, err)
	}
}

func RenderError(ctx context.Context, w http.ResponseWriter, cause error) {
	if err := renderError(ctx, w, http.StatusInternalServerError, cause); err != nil {
		logs.Error(ctx, err)
	}
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
	if res != nil {
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
	}

	w.WriteHeader(statusCode)
	return nil
}

func renderError(ctx context.Context, w http.ResponseWriter, statusCode int, cause error) error {
	if err := write(w, statusCode, nil); err != nil {
		return err
	}

	response := ctxlib.GetHTTPResponse(ctx)
	response.StatusCode = statusCode
	response.Error = cause
	response.WriteHandlerInfo(3) // render -> RenderXXX -> caller

	return nil
}
