import "../models/me.tsp";
import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

@service
namespace MonitorAPI;

@tag("Me")
@route("/me")
namespace MeAPI {
  @get
  @operationId("getMe")
  @doc("get own user info")
  op getMe(): {
    @doc("ok")
    @statusCode
    statusCode: 200;

    @body _: Me;
  };
}
