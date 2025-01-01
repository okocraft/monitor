package records

import (
	"github.com/okocraft/monitor/internal/repositories/queries"
	"time"
)

var Role1 = queries.InsertRoleWithIDForTestParams{
	ID:        1,
	Name:      "test_role_1",
	Priority:  1,
	CreatedAt: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
}
