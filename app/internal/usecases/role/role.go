package role

import (
	"context"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/sort"
	roleRepo "github.com/okocraft/monitor/internal/repositories/role"
	"github.com/okocraft/monitor/lib/errlib"
	"github.com/okocraft/monitor/lib/null"
)

type RoleUsecase interface {
	GetAllRoles(ctx context.Context, sortedBy null.Optional[role.SortableDataType], sortType null.Optional[sort.Type]) (role.Roles, error)
}

func NewRoleUsecase(repo roleRepo.RoleRepository) RoleUsecase {
	return roleUsecase{repo: repo}
}

type roleUsecase struct {
	repo roleRepo.RoleRepository
}

func (u roleUsecase) GetAllRoles(ctx context.Context, sortedBy null.Optional[role.SortableDataType], sortType null.Optional[sort.Type]) (role.Roles, error) {
	roles, err := u.repo.GetAllRoles(ctx, sortedBy, sortType)
	if err != nil {
		return nil, errlib.AsIs(err)
	}
	return roles, nil
}
