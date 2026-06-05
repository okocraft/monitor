package httplib

import (
	"encoding/json"
	"net/http"

	"github.com/Siroshun09/serrors/v2"
)

func DecodeBody[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		var zero T
		return zero, serrors.Wrap(err)
	}
	return v, nil
}
