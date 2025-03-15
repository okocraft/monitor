package user

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/domain/me"
	"github.com/okocraft/monitor/internal/domain/user"
	"github.com/okocraft/monitor/internal/repositories/database"
	roleRepo "github.com/okocraft/monitor/internal/repositories/role"
	userRepo "github.com/okocraft/monitor/internal/repositories/user"
	"github.com/okocraft/monitor/lib/ctxlib"
	"github.com/okocraft/monitor/lib/errlib"
)

type UserUsecase interface {
	FindUserIDBySub(ctx context.Context, sub string) (user.ID, error)
	SaveSubByLoginKey(ctx context.Context, loginKey int64, sub string) (user.ID, error)
	GetMe(ctx context.Context) (me.Me, error)
	GetNicknameByID(ctx context.Context, id user.ID) (string, error)

	GetUsersWithRoleByUUIDs(ctx context.Context, uuids []uuid.UUID) ([]user.UserWithRole, error)
	SearchForUserUUIDs(ctx context.Context, params user.SearchParams) ([]uuid.UUID, error)
}

func NewUserUsecase(repo userRepo.UserRepository, transaction database.Transaction) UserUsecase {
	return &userUsecase{
		repo:        repo,
		transaction: transaction,
	}
}

type userUsecase struct {
	repo        userRepo.UserRepository
	roleRepo    roleRepo.RoleRepository
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

	return me.Me{UUID: usr.UUID, Nickname: usr.Nickname}, nil
}

func (u userUsecase) GetNicknameByID(ctx context.Context, id user.ID) (string, error) {
	nickname, err := u.repo.GetUserNicknameByID(ctx, id)
	if err != nil {
		return "", errlib.AsIs(err)
	}
	return nickname, nil
}

func (u userUsecase) GetUsersWithRoleByUUIDs(ctx context.Context, uuids []uuid.UUID) ([]user.UserWithRole, error) {
	users, err := u.repo.GetUsersWithRoleByUUIDs(ctx, uuids)
	if err != nil {
		return nil, errlib.AsIs(err)
	}
	return users, nil
}

func (u userUsecase) SearchForUserUUIDs(ctx context.Context, params user.SearchParams) ([]uuid.UUID, error) {
	ids, err := u.repo.SearchForUserUUIDs(ctx, params)
	if err != nil {
		return nil, errlib.AsIs(err)
	}
	return ids, nil
}
