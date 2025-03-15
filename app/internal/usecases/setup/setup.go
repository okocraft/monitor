package setup

import (
	"context"
	"crypto/rand"
	"math"
	"math/big"
	"time"

	"github.com/Siroshun09/serrors"
	"github.com/gofrs/uuid/v5"
	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/user"
	permissionRepo "github.com/okocraft/monitor/internal/repositories/permission"
	roleRepo "github.com/okocraft/monitor/internal/repositories/role"
	userRepo "github.com/okocraft/monitor/internal/repositories/user"
	"github.com/okocraft/monitor/lib/errlib"
)

type SetupUsecase interface {
	IsFreshDatabase(ctx context.Context) (bool, error)
	CreateInitialAdminRole(ctx context.Context) (role.Role, error)
	CreateInitialAdminUser(ctx context.Context, adminRoleID role.ID) (user.User, error)
	CreateLoginKeyForAdminUser(ctx context.Context, userID user.ID) (string, error)
}

func NewSetupUsecase(roleRepo roleRepo.RoleRepository, userRepo userRepo.UserRepository, permissionRepo permissionRepo.PermissionRepository) SetupUsecase {
	return setupUsecase{
		roleRepo:       roleRepo,
		userRepo:       userRepo,
		permissionRepo: permissionRepo,
	}
}

type setupUsecase struct {
	roleRepo       roleRepo.RoleRepository
	userRepo       userRepo.UserRepository
	permissionRepo permissionRepo.PermissionRepository
}

func (u setupUsecase) IsFreshDatabase(ctx context.Context) (bool, error) {
	exists, err := u.roleRepo.ExistsRoleByID(ctx, 1) // checks by existing admin role
	if err != nil {
		return false, errlib.AsIs(err)
	}
	return exists, nil
}

func (u setupUsecase) CreateInitialAdminRole(ctx context.Context) (role.Role, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return role.Role{}, errlib.AsIs(err)
	}

	now := time.Now()
	r := role.Role{
		ID:        1,
		UUID:      id,
		Name:      "Administrator",
		Priority:  math.MaxInt32,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = u.roleRepo.CreateRoleWithIDIfNotExists(ctx, r)
	if err != nil {
		return role.Role{}, errlib.AsIs(err)
	}

	valueMap := permission.NewValueMap(map[permission.ID]bool{
		permission.Admin.ID: true,
	})

	err = u.permissionRepo.SaveRolePermissions(ctx, r.ID, valueMap)
	if err != nil {
		return role.Role{}, errlib.AsIs(err)
	}

	return r, nil
}

func (u setupUsecase) CreateInitialAdminUser(ctx context.Context, adminRoleID role.ID) (user.User, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return user.User{}, serrors.WithStackTrace(err)
	}
	now := time.Now()
	usr := user.User{
		ID:         1,
		UUID:       id,
		Nickname:   "Admin",
		LastAccess: now,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	err = u.userRepo.CreateUserWithIDIfNotExist(ctx, usr)
	if err != nil {
		return user.User{}, errlib.AsIs(err)
	}

	err = u.roleRepo.SaveUserRole(ctx, usr.ID, adminRoleID, now)
	if err != nil {
		return user.User{}, errlib.AsIs(err)
	}

	return usr, nil
}

func (u setupUsecase) CreateLoginKeyForAdminUser(ctx context.Context, userID user.ID) (string, error) {
	key, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		return "", serrors.WithStackTrace(err)
	}

	err = u.userRepo.SaveLoginKeyForUserID(ctx, userID, key.Int64(), time.Now())
	if err != nil {
		return "", errlib.AsIs(err)
	}
	return key.Text(16), nil
}
