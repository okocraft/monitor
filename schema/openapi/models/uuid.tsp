import "@typespec/openapi";

import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";

using TypeSpec.OpenAPI;

@format("uuid")
@extension("x-go-type", "uuid.UUID")
@extension(
  "x-go-type-import",
  #{ name: "uuid", path: "github.com/gofrs/uuid/v5" }
)
@doc("the UUID")
scalar UUID extends string;
