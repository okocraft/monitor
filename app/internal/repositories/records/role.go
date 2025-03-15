package records

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/repositories/queries"
)

var Role1 = queries.Role{
	ID:        1,
	Uuid:      uuid.Must(uuid.FromString("5655f692-119c-443a-9d12-080d50794a6f")).Bytes(),
	Name:      "test_role_1",
	Priority:  1,
	CreatedAt: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2024, 12, 26, 0, 0, 0, 0, time.UTC),
}

var Role2 = queries.Role{
	ID:        2,
	Uuid:      uuid.Must(uuid.FromString("043317b7-a81e-410d-a655-ab23bdc2558f")).Bytes(),
	Name:      "test_role_2",
	Priority:  2,
	CreatedAt: time.Date(2024, 12, 27, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2024, 12, 28, 0, 0, 0, 0, time.UTC),
}

var Role3 = queries.Role{
	ID:        3,
	Uuid:      uuid.Must(uuid.FromString("c4bf7afc-a462-4856-9af9-a635776bafa4")).Bytes(),
	Name:      "test_role_3",
	Priority:  3,
	CreatedAt: time.Date(2024, 12, 29, 0, 0, 0, 0, time.UTC),
	UpdatedAt: time.Date(2024, 12, 30, 0, 0, 0, 0, time.UTC),
}
