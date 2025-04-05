import "../models/me.tsp";
import "../models/user.tsp";
import "../models/sort.tsp";
import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";

using TypeSpec.Http;
using TypeSpec.OpenAPI;

@service
namespace MonitorAPI;

@tag("User")
@route("/users")
namespace UserAPI {
  @post
  @operationId("getUsersByIds")
  @doc("get users by ids")
  op getUsersByIds(@body _: UUID[]): {
    @doc("ok")
    @statusCode
    statusCode: 200;

    @body _: User[];
  } | {
    @doc("too many ids are specified (should be <= 50)")
    @statusCode
    statusCode: 400;
  } | {
    @doc("the request contains invalid user ids")
    @statusCode
    statusCode: 404;
  };

  @get
  @operationId("searchUsers")
  @doc("get user ids")
  @route("/search")
  op searchUsers(
    @query
    nickname?: string,

    @query
    last_access_before?: utcDateTime,

    @query
    last_access_after?: utcDateTime,

    @query
    role_id?: UUID,

    @query
    sorted_by?: SortableUserDataType,

    @query
    sort_type?: SortType,
  ): {
    @doc("ok")
    @statusCode
    statusCode: 200;

    @body _: UUID[];
  };
}
