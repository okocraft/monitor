package user

import "strconv"

type NotFoundByIDError struct {
	ID ID
}

func (e NotFoundByIDError) Error() string {
	return "user not found: " + e.ID.String()
}

type NotFoundBySubError struct {
	Sub string
}

func (e NotFoundBySubError) Error() string {
	return "user not found with sub: " + e.Sub
}

type NotFoundByLoginKeyError struct {
	LoginKey int64
}

func (e NotFoundByLoginKeyError) Error() string {
	return "user not found with login key: " + strconv.FormatInt(e.LoginKey, 10)
}
