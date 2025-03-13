package definition

// role permissions
var _ = definer.define(
	permission("role.list", restrictedByDefault),
)
