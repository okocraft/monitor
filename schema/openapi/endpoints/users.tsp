import "../models/me.tsp";
import "../models/user.tsp";
import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";

using TypeSpec.Http;
using TypeSpec.Rest;
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
}
