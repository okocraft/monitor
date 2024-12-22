package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Siroshun09/serrors"
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
		return user.User{}, asDBError(err)
	}

	return user.User{ID: id, UUID: uuid.UUID(row.Uuid), Nickname: row.Nickname, LastAccess: row.LastAccess}, nil
}

func (r userRepository) GetUserNicknameByID(ctx context.Context, id user.ID) (string, error) {
	q := r.db.Queries(ctx)
	row, err := q.GetUserNicknameByID(ctx, int32(id))
	if errors.Is(err, sql.ErrNoRows) {
		return "", serrors.WithStackTrace(user.NotFoundByIDError{ID: id})
	} else if err != nil {
		return "", asDBError(err)
	}
	return row, nil
}

func (r userRepository) GetUserIDBySub(ctx context.Context, sub string) (user.ID, error) {
	q := r.db.Queries(ctx)
	userID, err := q.GetUserIDBySub(ctx, sub)
	if errors.Is(err, sql.ErrNoRows) {
		return 0, serrors.WithStackTrace(user.NotFoundBySubError{Sub: sub})
	} else if err != nil {
		return 0, asDBError(err)
	}

	return user.ID(userID), nil
}

func (r userRepository) GetUserIDByLoginKey(ctx context.Context, loginKey int64) (user.ID, error) {
	q := r.db.Queries(ctx)
	userID, err := q.GetUserIDByLoginKey(ctx, loginKey)

	if errors.Is(err, sql.ErrNoRows) {
		return 0, serrors.WithStackTrace(user.NotFoundByLoginKeyError{LoginKey: loginKey})
	} else if err != nil {
		return 0, asDBError(err)
	}

	return user.ID(userID), nil
}

func (r userRepository) DeleteLoginKeyByUserID(ctx context.Context, id user.ID) error {
	q := r.db.Queries(ctx)
	err := q.DeleteLoginKey(ctx, int32(id))
	if err != nil {
		return asDBError(err)
	}
	return nil
}

func (r userRepository) SaveUserSub(ctx context.Context, userID user.ID, sub string) error {
	q := r.db.Queries(ctx)
	err := q.InsertSubForUserID(ctx, queries.InsertSubForUserIDParams{UserID: int32(userID), Sub: sub})
	if err != nil {
		return asDBError(err)
	}
	return nil
}

func (r userRepository) UpdateLastAccessByID(ctx context.Context, id user.ID, now time.Time) error {
	q := r.db.Queries(ctx)
	err := q.UpdateUserLastAccessByID(ctx, queries.UpdateUserLastAccessByIDParams{ID: int32(id), LastAccess: now, UpdatedAt: now})
	if err != nil {
		return asDBError(err)
	}
	return nil
}
