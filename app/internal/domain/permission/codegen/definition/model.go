package definition

type DefaultPermissionValue bool

const (
	allowedByDefault    DefaultPermissionValue = true
	restrictedByDefault DefaultPermissionValue = false
)

func (value DefaultPermissionValue) IsAllowedByDefault() bool {
	return value == allowedByDefault
}

type DefinedPermission struct {
	Name         string
	DefaultValue DefaultPermissionValue
}

type DefinedPermissions []DefinedPermission

func permission(name string, value DefaultPermissionValue) DefinedPermission {
	return DefinedPermission{name, value}
}

type PermissionDefiner struct {
	definedPermissions DefinedPermissions
}

func (def *PermissionDefiner) define(perms ...DefinedPermission) *PermissionDefiner {
	def.definedPermissions = append(def.definedPermissions, perms...)
	return def
}

func (def *PermissionDefiner) GetPermissions() DefinedPermissions {
	return def.definedPermissions
}

var definer = &PermissionDefiner{
	definedPermissions: make(DefinedPermissions, 0, 100),
}

func GetPermissions() DefinedPermissions {
	return definer.definedPermissions
}
