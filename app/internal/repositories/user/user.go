package user

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/domain/role"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id user.ID) (user.User, error)
	GetUserNicknameByID(ctx context.Context, id user.ID) (string, error)
	GetUserIDBySub(ctx context.Context, sub string) (user.ID, error)
	GetUserIDByLoginKey(ctx context.Context, loginKey int64) (user.ID, error)
	DeleteLoginKeyByUserID(ctx context.Context, id user.ID) error
	SaveUserSub(ctx context.Context, userID user.ID, sub string) error
	UpdateLastAccessByID(ctx context.Context, id user.ID, now time.Time) error

	GetUsersWithRoleByUUIDs(ctx context.Context, uuids []uuid.UUID) ([]user.UserWithRole, error)
}

func NewUserRepository(db database.DB) UserRepository {
	return userRepository{db: db}
}

type userRepository struct {
	db database.DB
}

func (r userRepository) GetUserByID(ctx context.Context, id user.ID) (user.User, error) {
	q := r.db.Queries(ctx)
	row, err := q.GetUserByID(ctx, int32(id))
	if errors.Is(err, sql.ErrNoRows) {
		return user.User{}, serrors.WithStackTrace(user.NotFoundByIDError{ID: id})
	} else if err != nil {
		return user.User{}, database.NewDBErrorWithStackTrace(err)
	}

	return user.User{ID: id, UUID: uuid.UUID(row.Uuid), Nickname: row.Nickname, LastAccess: row.LastAccess}, nil
}

func (r userRepository) GetUserNicknameByID(ctx context.Context, id user.ID) (string, error) {
	q := r.db.Queries(ctx)
	row, err := q.GetUserNicknameByID(ctx, int32(id))
	if errors.Is(err, sql.ErrNoRows) {
		return "", serrors.WithStackTrace(user.NotFoundByIDError{ID: id})
	} else if err != nil {
		return "", database.NewDBErrorWithStackTrace(err)
	}
	return row, nil
}

func (r userRepository) GetUserIDBySub(ctx context.Context, sub string) (user.ID, error) {
	q := r.db.Queries(ctx)
	userID, err := q.GetUserIDBySub(ctx, sub)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, serrors.WithStackTrace(user.NotFoundBySubError{Sub: sub})
	} else if err != nil {
		return 0, database.NewDBErrorWithStackTrace(err)
	}

	return user.ID(userID), nil
}

func (r userRepository) GetUserIDByLoginKey(ctx context.Context, loginKey int64) (user.ID, error) {
	q := r.db.Queries(ctx)
	userID, err := q.GetUserIDByLoginKey(ctx, loginKey)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, serrors.WithStackTrace(user.NotFoundByLoginKeyError{LoginKey: loginKey})
	} else if err != nil {
		return 0, database.NewDBErrorWithStackTrace(err)
	}

	return user.ID(userID), nil
}

func (r userRepository) DeleteLoginKeyByUserID(ctx context.Context, id user.ID) error {
	q := r.db.Queries(ctx)
	err := q.DeleteLoginKey(ctx, int32(id))
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}

func (r userRepository) SaveUserSub(ctx context.Context, userID user.ID, sub string) error {
	q := r.db.Queries(ctx)
	err := q.InsertSubForUserID(ctx, queries.InsertSubForUserIDParams{UserID: int32(userID), Sub: sub})
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}

func (r userRepository) UpdateLastAccessByID(ctx context.Context, id user.ID, now time.Time) error {
	q := r.db.Queries(ctx)
	err := q.UpdateUserLastAccessByID(ctx, queries.UpdateUserLastAccessByIDParams{ID: int32(id), LastAccess: now, UpdatedAt: now})
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}

func (r userRepository) GetUsersWithRoleByUUIDs(ctx context.Context, uuids []uuid.UUID) ([]user.UserWithRole, error) {
	q := r.db.Queries(ctx)
	rows, err := q.GetUsersWithRoleByUUIDs(ctx, queries.ToBytesSlice(uuids))
	if err != nil {
		return []user.UserWithRole{}, database.NewDBErrorWithStackTrace(err)
	}

	users := make([]user.UserWithRole, 0, len(rows))
	for _, row := range rows {
		var userRole role.Role
		if row.RoleID.Valid {
			userRole = role.Role{
				ID:        row.RoleID.Int32,
				Name:      row.RoleName.String,
				Priority:  row.RolePriority.Int32,
				CreatedAt: row.RoleCreatedAt.Time,
				UpdatedAt: row.RoleUpdatedAt.Time,
			}
		} else {
			userRole = role.DefaultRole()
		}
		users = append(users, user.UserWithRole{
			User: user.User{
				ID:         user.ID(row.UserID),
				UUID:       uuid.UUID(row.UserUuid),
				Nickname:   row.UserNickname,
				LastAccess: row.UserLastAccess,
				CreatedAt:  row.UserCreatedAt,
				UpdatedAt:  row.UserUpdatedAt,
			},
			Role: userRole,
		})
	}

	return users, nil
}
