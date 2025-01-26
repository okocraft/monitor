// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: permission.sql

package queries

import (
	"context"
	"strings"
)

const checkPermissionByUserID = `-- name: CheckPermissionByUserID :one
SELECT is_allowed
FROM roles_permissions
         INNER JOIN users_role ON user_id = ?
WHERE roles_permissions.role_id = users_role.role_id
  AND permission_id = ?
`

type CheckPermissionByUserIDParams struct {
	UserID       int32 `db:"user_id"`
	PermissionID int16 `db:"permission_id"`
}

func (q *Queries) CheckPermissionByUserID(ctx context.Context, arg CheckPermissionByUserIDParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkPermissionByUserID, arg.UserID, arg.PermissionID)
	var is_allowed bool
	err := row.Scan(&is_allowed)
	return is_allowed, err
}

const getPermissionsByUserID = `-- name: GetPermissionsByUserID :many
SELECT permission_id, is_allowed
FROM roles_permissions
         INNER JOIN users_role ON user_id = ?
WHERE roles_permissions.role_id = users_role.role_id
  AND permission_id IN (/*SLICE:permission_ids*/?)
`

type GetPermissionsByUserIDParams struct {
	UserID        int32   `db:"user_id"`
	PermissionIds []int16 `db:"permission_ids"`
}

type GetPermissionsByUserIDRow struct {
	PermissionID int16 `db:"permission_id"`
	IsAllowed    bool  `db:"is_allowed"`
}

func (q *Queries) GetPermissionsByUserID(ctx context.Context, arg GetPermissionsByUserIDParams) ([]GetPermissionsByUserIDRow, error) {
	query := getPermissionsByUserID
	var queryParams []interface{}
	queryParams = append(queryParams, arg.UserID)
	if len(arg.PermissionIds) > 0 {
		for _, v := range arg.PermissionIds {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:permission_ids*/?", strings.Repeat(",?", len(arg.PermissionIds))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:permission_ids*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetPermissionsByUserIDRow
	for rows.Next() {
		var i GetPermissionsByUserIDRow
		if err := rows.Scan(&i.PermissionID, &i.IsAllowed); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
