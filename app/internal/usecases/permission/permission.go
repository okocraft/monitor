package permission

import (
	"context"

	"github.com/Siroshun09/serrors"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/user"
	permission2 "github.com/okocraft/monitor/internal/repositories/permission"
	"github.com/okocraft/monitor/lib/ctxlib"
	"github.com/okocraft/monitor/lib/errlib"
)

type PermissionUsecase interface {
	CalculatePagePermissions(ctx context.Context) (permission.PagePermissions, error)
	HasPermission(ctx context.Context, userID user.ID, permission permission.Permission) (bool, error)
}

func NewPermissionUsecase(repo permission2.PermissionRepository) PermissionUsecase {
	return permissionUsecase{repo: repo}
}

type permissionUsecase struct {
	repo permission2.PermissionRepository
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

func (u permissionUsecase) HasPermission(ctx context.Context, userID user.ID, permission permission.Permission) (bool, error) {
	hasPermission, err := u.repo.HasPermission(ctx, userID, permission)
	if err != nil {
		return false, errlib.AsIs(err)
	}
	return hasPermission, nil
}
