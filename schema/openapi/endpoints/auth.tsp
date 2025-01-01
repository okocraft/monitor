import "../models/auth.tsp";
import "../models/me.tsp";
import "@typespec/http";
import "@typespec/rest";
import "@typespec/openapi";

using TypeSpec.Http;
using TypeSpec.Rest;
using TypeSpec.OpenAPI;

@service
namespace MonitorAPI;

@tag("Auth")
@route("/auth")
@useAuth(SkipAuth)
namespace AuthAPI {
  @route("/logout")
  @post
  @operationId("logout")
  @doc("Invalidate refresh_token and access_token")
  op logout(): {
    @doc("refresh token has been invalidated")
    @statusCode
    statusCode: 204;
  } | {
    @doc("invalid refresh token")
    @statusCode
    statusCode: 401;
  };

  @route("/refresh")
  @post
  @operationId("refreshAccessToken")
  @doc("Refresh access token")
  op refreshAccessToken(): {
    @doc("successfully refreshed access token")
    @statusCode
    statusCode: 200;

    @body _: AccessTokenWithMeAndPagePermissions;
  } | {
    @doc("invalid refresh token")
    @statusCode
    statusCode: 401;
  };

  @route("/google")
  namespace Google {
    @route("/link/{loginKey}")
    @post
    @operationId("linkWithGoogle")
    @doc("First login with Google Account")
    op linkWithGoogle(@path loginKey: string): {
      @doc("the URL of Google's account page")
      @statusCode
      statusCode: 200;

      @body _: GoogleLoginURL;
    } | {
      @doc("if Google login is not enabled")
      @statusCode
      statusCode: 406;
    };

    @route("/login")
    @post
    @operationId("loginWithGoogle")
    @doc("Login with Google Account")
    op loginWithGoogle(@body _: CurrentPage): {
      @doc("the URL of Google's account page")
      @statusCode
      statusCode: 200;

      @body _: GoogleLoginURL;
    } | {
      @doc("if Google login is not enabled")
      @statusCode
      statusCode: 406;
    };

    @route("/callback")
    @get
    @operationId("callbackFromGoogle")
    @doc("Callback from Google")
    op callbackFromGoogle(): {
      @doc("redirect to login result page")
      @statusCode
      statusCode: 307;
    } | {
      @doc("if Google login is not enabled")
      @statusCode
      statusCode: 406;
    };
  }
}
