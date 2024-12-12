package usecases

import (
	"context"
	"errors"
	"github.com/okocraft/monitor/internal/domain/me"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories"
	"github.com/okocraft/monitor/internal/repositories/database"
	"github.com/okocraft/monitor/lib/ctxlib"
	"github.com/okocraft/monitor/lib/errlib"
	"time"
)

type UserUsecase interface {
	FindUserIDBySub(ctx context.Context, sub string) (user.ID, error)
	SaveSubByLoginKey(ctx context.Context, loginKey int64, sub string) (user.ID, error)
	GetMe(ctx context.Context) (me.Me, error)
}

func NewUserUsecase(repo repositories.UserRepository, transaction database.Transaction) UserUsecase {
	return &userUsecase{
		repo:        repo,
		transaction: transaction,
	}
}

type userUsecase struct {
	repo        repositories.UserRepository
	transaction database.Transaction
}

func (u userUsecase) FindUserIDBySub(ctx context.Context, sub string) (user.ID, error) {
	userID, err := u.repo.GetUserIDBySub(ctx, sub)

	if errors.Is(err, user.NotFoundBySubError{Sub: sub}) {
		return 0, errlib.AsIs(err)
	}

	return userID, nil
}

func (u userUsecase) SaveSubByLoginKey(ctx context.Context, loginKey int64, sub string) (user.ID, error) {
	var result user.ID
	err := u.transaction.WithTx(ctx, func(ctx context.Context) error {
		userID, err := u.repo.GetUserIDByLoginKey(ctx, loginKey)
		if err != nil {
			return errlib.AsIs(err)
		}

		err = u.repo.DeleteLoginKeyByUserID(ctx, userID)
		if err != nil {
			return errlib.AsIs(err)
		}

		err = u.repo.SaveUserSub(ctx, userID, sub)
		if err != nil {
			return errlib.AsIs(err)
		}

		result = userID
		return nil
	})

	if err != nil {
		return 0, errlib.AsIs(err)
	}

	return result, nil
}

func (u userUsecase) GetMe(ctx context.Context) (me.Me, error) {
	id, ok := ctxlib.GetUserID(ctx)
	if !ok {
		return me.Me{}, user.NotFoundByIDError{ID: id}
	}

	usr, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		return me.Me{}, errlib.AsIs(err)
	}

	err = u.repo.UpdateLastAccessByID(ctx, id, time.Now())
	if err != nil {
		return me.Me{}, errlib.AsIs(err)
	}

	return me.Me{UUID: usr.UUID, NickName: usr.NickName}, nil
}