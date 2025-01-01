import "./me.tsp";
import "./permission.tsp";

@doc("Authorize by access tokens")
model AccessTokenAuth {
  type: "http";
  scheme: "bearer";
}

@doc("No authentication needed")
model SkipAuth {
  type: "http";
  scheme: "none";
}

model AccessTokenWithMeAndPagePermissions {
  @doc("the access token")
  access_token: string;

  @doc("the currently logged-in user info")
  me: Me;

  @doc("the page permissions")
  page_permissions: PagePermissions;
}

model GoogleLoginURL {
  @format("url")
  @doc("the Google's login page URL")
  redirect_url: string;
}

model CurrentPage {
  @format("url")
  @doc("the url of the page currently being viewed")
  url: string;
}
