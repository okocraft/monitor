package records

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/repositories/queries"
)

var User1 = queries.User{
	ID:         1,
	Uuid:       uuid.Must(uuid.FromString("88938686-ffa4-4d8f-8478-b8ade2b26e59")).Bytes(),
	Nickname:   "test_user_1",
	LastAccess: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
	CreatedAt:  time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
	UpdatedAt:  time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
}

var User2 = queries.User{
	ID:         2,
	Uuid:       uuid.Must(uuid.FromString("fe61251e-0b12-494b-a4d1-4107dfc5e353")).Bytes(),
	Nickname:   "test_user_2",
	LastAccess: time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC),
	CreatedAt:  time.Date(2024, 12, 27, 0, 0, 0, 0, time.UTC),
	UpdatedAt:  time.Date(2024, 12, 28, 0, 0, 0, 0, time.UTC),
}

var User3 = queries.User{
	ID:         3,
	Uuid:       uuid.Must(uuid.FromString("5b5475fd-f68b-427b-9d71-ad47be09b517")).Bytes(),
	Nickname:   "test_user_3",
	LastAccess: time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC),
	CreatedAt:  time.Date(2024, 12, 28, 0, 0, 0, 0, time.UTC),
	UpdatedAt:  time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC),
}
