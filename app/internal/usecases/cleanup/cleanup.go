package cleanup

import (
	"context"

	"github.com/okocraft/monitor/internal/domain/cleanup"
	"github.com/okocraft/monitor/internal/repositories/auth"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/lib/errlib"
)

type CleanupUsecase interface {
	CleanupExpiredTokens(ctx context.Context, param cleanup.Param) (cleanup.DeletedTokenResult, error)
}

func NewCleanupUsecase(authRepo auth.AuthRepository, tx database.Transaction) CleanupUsecase {
	return &cleanupUsecase{
		authRepo: authRepo,
		tx:       tx,
	}
}

type cleanupUsecase struct {
	authRepo auth.AuthRepository
	tx       database.Transaction
}

func (u cleanupUsecase) CleanupExpiredTokens(ctx context.Context, param cleanup.Param) (cleanup.DeletedTokenResult, error) {
	result := &cleanup.DeletedTokenResult{}

	err := u.tx.WithTx(ctx, func(ctx context.Context) error {
		count, err := u.authRepo.DeleteExpiredAccessTokens(ctx, param.AccessTokenExpiredAt)
		if err != nil {
			return errlib.AsIs(err)
		}
		result.AccessTokens = count

		count, err = u.authRepo.DeleteAccessTokensByExpiredRefreshTokens(ctx, param.RefreshTokenExpiredAt)
		if err != nil {
			return errlib.AsIs(err)
		}
		result.AccessTokens += count

		count, err = u.authRepo.DeleteExpiredRefreshTokens(ctx, param.RefreshTokenExpiredAt)
		if err != nil {
			return errlib.AsIs(err)
		}
		result.RefreshTokens = count

		return nil
	})
	if err != nil {
		return cleanup.DeletedTokenResult{}, errlib.AsIs(err)
	}

	return *result, nil
}
