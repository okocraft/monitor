package errlib

import (
	"errors"

	"github.com/Siroshun09/serrors"
)

type detailedError struct {
	detail any
	err    error
}

func (e *detailedError) Error() string {
	return e.err.Error()
}

func (e *detailedError) Unwrap() error {
	return e.err
}

func New[T any](detail T, msg string) error {
	return &detailedError{detail: detail, err: serrors.New(msg)}
}

func Errorf[T any](detail T, format string, args ...any) error {
	return &detailedError{detail: detail, err: serrors.Errorf(format, args...)}
}

func WithDetail[T any](detail T, err error) error {
	return &detailedError{detail: detail, err: serrors.WithStackTrace(err)}
}

func GetDetail[T any](err error) *T {
	currentErr := err
	for {
		derr := getDetailError(currentErr)
		if derr == nil {
			return nil
		}

		detail, ok := derr.detail.(T)
		if !ok {
			currentErr = errors.Unwrap(currentErr)
		}

		return &detail
	}
}

func GetRawDetail(err error) any {
	derr := getDetailError(err)
	if derr == nil {
		return nil
	}

	return derr.detail
}

func getDetailError(err error) *detailedError {
	var detail *detailedError
	if errors.As(err, &detail) {
		return detail
	}

	return nil
}

func AsIs(err error) error {
	return err
}
