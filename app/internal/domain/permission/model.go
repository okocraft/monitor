package permission

type PermissionID = int16

type Permission struct {
	ID           PermissionID
	Name         string
	DefaultValue bool
}
