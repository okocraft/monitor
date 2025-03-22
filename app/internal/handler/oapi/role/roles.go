package role

import (
	"net/http"

	"github.com/okocraft/monitor/internal/domain/permission"
	"github.com/okocraft/monitor/internal/domain/role"
	"github.com/okocraft/monitor/internal/domain/sort"
	"github.com/okocraft/monitor/internal/handler/oapi"
	permissionUsecase "github.com/okocraft/monitor/internal/usecases/permission"
	roleUsecase "github.com/okocraft/monitor/internal/usecases/role"
	"github.com/okocraft/monitor/lib/ctxlib"
	"github.com/okocraft/monitor/lib/errlib"
	"github.com/okocraft/monitor/lib/httplib"
	"github.com/okocraft/monitor/lib/null"
)

type RoleHandler struct {
	usecase           roleUsecase.RoleUsecase
	permissionUsecase permissionUsecase.PermissionUsecase
}

func NewRoleHandler(usecase roleUsecase.RoleUsecase, permissionUsecase permissionUsecase.PermissionUsecase) RoleHandler {
	return RoleHandler{usecase: usecase, permissionUsecase: permissionUsecase}
}

func (h RoleHandler) GetRoles(w http.ResponseWriter, r *http.Request, params oapi.GetRolesParams) {
	ctx := r.Context()

	userID, ok := ctxlib.GetUserID(ctx)
	if !ok {
		httplib.RenderUnauthorized(ctx, w, nil)
		return
	}

	if hasPermission, err := h.permissionUsecase.HasPermission(ctx, userID, permission.RoleList); err != nil {
		httplib.RenderError(ctx, w, errlib.AsIs(err))
		return
	} else if !hasPermission {
		httplib.RenderForbidden(ctx, w, nil)
		return
	}

	var sortedBy null.Optional[role.SortableDataType]
	if params.SortedBy != nil {
		converted, err := role.ConvertSortableDataTypeFromRequest(*params.SortedBy)
		if err != nil {
			httplib.RenderBadRequest(ctx, w, errlib.AsIs(err))
			return
		}
		sortedBy = null.FromValue(converted)
	}

	var sortType null.Optional[sort.Type]
	if params.SortedBy != nil {
		converted, err := sort.FromRequest(*params.SortType)
		if err != nil {
			httplib.RenderBadRequest(ctx, w, errlib.AsIs(err))
			return
		}
		sortType = null.FromValue(converted)
	}

	roles, err := h.usecase.GetAllRoles(ctx, sortedBy, sortType)
	if err != nil {
		httplib.RenderError(ctx, w, err)
		return
	}

	httplib.RenderOK(ctx, w, roles.ToResponse())
}
