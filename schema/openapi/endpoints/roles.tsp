import "../models/role.tsp";
import "../models/sort.tsp";
import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

@service
namespace MonitorAPI;

@tag("Role")
@route("/roles")
namespace RoleAPI {
  @get
  @operationId("getRoles")
  @doc("get all roles")
  op getUsersByIds(
    @query sorted_by?: SortableRoleDataType,
    @query sort_type?: SortType,
  ): {
    @doc("ok")
    @statusCode
    statusCode: 200;

    @body _: Role[];
  };
}
