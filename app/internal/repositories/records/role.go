package records

import (
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/repositories/queries"
	"time"
)

var Role1 = queries.Role{
	ID:        1,
	Uuid:      uuid.Must(uuid.FromString("5655f692-119c-443a-9d12-080d50794a6f")).Bytes(),
	Name:      "test_role_1",
	Priority:  1,
	CreatedAt: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
}
