package user

import (
	"strconv"
	"time"

	"github.com/gofrs/uuid/v5"
)

type ID int32

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

type User struct {
	ID         ID
	UUID       uuid.UUID
	NickName   string
	LastAccess time.Time
}
