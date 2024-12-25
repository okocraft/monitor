package records

import (
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"time"
)

var User1 = queries.InsertUserWithIDForTestParams{
	ID:         1,
	Uuid:       uuid.Must(uuid.FromString("88938686-ffa4-4d8f-8478-b8ade2b26e59")).Bytes(),
	Nickname:   "test_user_1",
	LastAccess: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
	CreatedAt:  time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
	UpdatedAt:  time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
}
