package definition

// user permissions
var _ = definer.define(
	permission("user.list", restrictedByDefault),
	permission("user.create", restrictedByDefault),
	permission("user.edit.nickname", restrictedByDefault),
	permission("user.delete", restrictedByDefault),
)

// mypage permissions
var _ = definer.define(
	permission("mypage.view", allowedByDefault),
	permission("mypage.edit.nickname", allowedByDefault),
)
