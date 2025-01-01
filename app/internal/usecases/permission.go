package usecases

import (
	"context"
	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/repositories"
	"github.com/okocraft/monitor/lib/ctxlib"
	"github.com/okocraft/monitor/lib/errlib"
)

type PermissionUsecase interface {
	CalculatePagePermissions(ctx context.Context) (permission.PagePermissions, error)
}

func NewPermissionUsecase(repo repositories.PermissionRepository) PermissionUsecase {
	return permissionUsecase{repo: repo}
}

type permissionUsecase struct {
	repo repositories.PermissionRepository
}

func (u permissionUsecase) CalculatePagePermissions(ctx context.Context) (permission.PagePermissions, error) {
	userID, ok := ctxlib.GetUserID(ctx)
	if !ok {
		return permission.PagePermissions{}, serrors.New("userID not found in ctx")
	}

	calculator := permission.GetPagePermissionCalculator()
	valueMap, err := u.repo.GetPermissions(ctx, userID, calculator.GetSourcePermissions()...)
	if err != nil {
		return permission.PagePermissions{}, errlib.AsIs(err)
	}

	return calculator.Calculate(valueMap), nil
}
