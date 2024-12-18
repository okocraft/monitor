import "./me.tsp";

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

model AccessTokenWithMe {
  @doc("the access token")
  access_token: string;

  @doc("the currently logged-in user info")
  me: Me;
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
