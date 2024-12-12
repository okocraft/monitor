package httplib

import (
	"encoding/json"
	"github.com/Siroshun09/serrors"
	"net/http"
)

func DecodeBody[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		var zero T
		return zero, serrors.WithStackTrace(err)
	}
	return v, nil
}
