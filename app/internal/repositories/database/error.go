package database

type DBError struct {
	err error
}

func (e DBError) Error() string {
	return e.err.Error()
}

func (e DBError) Unwrap() error {
	return e.err
}

func NewDBError(err error) DBError {
	return DBError{err: err}
}
