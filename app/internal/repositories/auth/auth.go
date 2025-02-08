package auth

import (
	"context"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/internal/repositories/queries"
)

type AuthRepository interface {
	SaveRefreshToken(ctx context.Context, userID user.ID, jti uuid.UUID, ip []byte, userAgent string, createdAt time.Time) error
	SaveAccessToken(ctx context.Context, userID user.ID, refreshTokenID int64, jti uuid.UUID, createdAt time.Time) error
	GetUserIDAndRefreshTokenIDFromJTI(ctx context.Context, jti uuid.UUID) (user.ID, int64, error)
	InvalidateTokensByRefreshTokenID(ctx context.Context, refreshTokenID int64) error
	GetUserIDFromAccessTokenJTI(ctx context.Context, jti uuid.UUID) (user.ID, error)
	DeleteExpiredAccessTokens(ctx context.Context, expiredAt time.Time) (int64, error)
	DeleteAccessTokensByExpiredRefreshTokens(ctx context.Context, expiredAt time.Time) (int64, error)
	DeleteExpiredRefreshTokens(ctx context.Context, expiredAt time.Time) (int64, error)
}

func NewAuthRepository(db database.DB) AuthRepository {
	return authRepository{db: db}
}

type authRepository struct {
	db database.DB
}

func (r authRepository) SaveRefreshToken(ctx context.Context, userID user.ID, jti uuid.UUID, ip []byte, userAgent string, createdAt time.Time) error {
	q := r.db.Queries(ctx)
	err := q.InsertRefreshToken(ctx, queries.InsertRefreshTokenParams{UserID: int32(userID), Jti: jti.Bytes(), Ip: ip, UserAgent: userAgent, CreatedAt: createdAt})
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}

func (r authRepository) SaveAccessToken(ctx context.Context, userID user.ID, refreshTokenID int64, jti uuid.UUID, createdAt time.Time) error {
	q := r.db.Queries(ctx)
	err := q.InsertAccessToken(ctx, queries.InsertAccessTokenParams{UserID: int32(userID), RefreshTokenID: refreshTokenID, Jti: jti.Bytes(), CreatedAt: createdAt})
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}

func (r authRepository) GetUserIDAndRefreshTokenIDFromJTI(ctx context.Context, jti uuid.UUID) (user.ID, int64, error) {
	q := r.db.Queries(ctx)
	row, err := q.GetUserIDAndRefreshTokenIDByJTI(ctx, jti.Bytes())
	if err != nil {
		return 0, 0, err
	}

	return user.ID(row.UserID), row.ID, nil
}

func (r authRepository) InvalidateTokensByRefreshTokenID(ctx context.Context, refreshTokenID int64) error {
	q := r.db.Queries(ctx)
	err := q.DeleteRefreshTokenAndAccessToken(ctx, refreshTokenID)
	if err != nil {
		return database.NewDBErrorWithStackTrace(err)
	}
	return nil
}

func (r authRepository) GetUserIDFromAccessTokenJTI(ctx context.Context, jti uuid.UUID) (user.ID, error) {
	q := r.db.Queries(ctx)
	userID, err := q.GetUserIDByAccessTokenJTI(ctx, jti.Bytes())
	if err != nil {
		return 0, database.NewDBErrorWithStackTrace(err)
	}
	return user.ID(userID), nil
}

func (r authRepository) DeleteExpiredAccessTokens(ctx context.Context, expiredAt time.Time) (int64, error) {
	q := r.db.Queries(ctx)
	rows, err := q.DeleteExpiredAccessTokens(ctx, expiredAt)
	if err != nil {
		return 0, database.NewDBErrorWithStackTrace(err)
	}
	return rows, nil
}

func (r authRepository) DeleteAccessTokensByExpiredRefreshTokens(ctx context.Context, expiredAt time.Time) (int64, error) {
	q := r.db.Queries(ctx)
	rows, err := q.DeleteAccessTokensByExpiredRefreshTokens(ctx, expiredAt)
	if err != nil {
		return 0, database.NewDBErrorWithStackTrace(err)
	}
	return rows, nil
}

func (r authRepository) DeleteExpiredRefreshTokens(ctx context.Context, expiredAt time.Time) (int64, error) {
	q := r.db.Queries(ctx)
	rows, err := q.DeleteExpiredRefreshTokens(ctx, expiredAt)
	if err != nil {
		return 0, database.NewDBErrorWithStackTrace(err)
	}
	return rows, nil
}
