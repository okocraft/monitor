import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi3";
import "./models/auth.tsp";
import "./endpoints/auth.tsp";
import "./endpoints/me.tsp";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

@service({
  title: "Monitor API",
})
@useAuth(AccessTokenAuth)
namespace MonitorAPI;
