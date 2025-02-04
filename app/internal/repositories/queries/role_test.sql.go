// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: role_test.sql

package queries

import (
	"context"
	"time"
)

const insertRoleWithIDForTest = `-- name: InsertRoleWithIDForTest :exec
INSERT INTO roles (id, name, priority, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
`

type InsertRoleWithIDForTestParams struct {
	ID        int32     `db:"id"`
	Name      string    `db:"name"`
	Priority  int32     `db:"priority"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (q *Queries) InsertRoleWithIDForTest(ctx context.Context, arg InsertRoleWithIDForTestParams) error {
	_, err := q.db.ExecContext(ctx, insertRoleWithIDForTest,
		arg.ID,
		arg.Name,
		arg.Priority,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const insertUserRoleForTest = `-- name: InsertUserRoleForTest :exec
INSERT INTO users_role (user_id, role_id, updated_at)
VALUES (?, ?, ?)
`

type InsertUserRoleForTestParams struct {
	UserID    int32     `db:"user_id"`
	RoleID    int32     `db:"role_id"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (q *Queries) InsertUserRoleForTest(ctx context.Context, arg InsertUserRoleForTestParams) error {
	_, err := q.db.ExecContext(ctx, insertUserRoleForTest, arg.UserID, arg.RoleID, arg.UpdatedAt)
	return err
}
