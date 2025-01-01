package permission

type ID = int16

type Permission struct {
	ID           ID
	Name         string
	DefaultValue bool
}

type ValueMap map[ID]bool

func (valueMap ValueMap) IsTrue(id ID) bool {
	v, ok := valueMap[id]
	return ok && v
}

func (valueMap ValueMap) HasPermission(perm Permission) bool {
	if v, ok := valueMap[perm.ID]; ok {
		return v
	}
	return perm.DefaultValue
}

func (valueMap ValueMap) HasAllPermissions(perms ...Permission) bool {
	for _, perm := range perms {
		if !valueMap.HasPermission(perm) {
			return false
		}
	}
	return true
}

func (valueMap ValueMap) HasAnyPermissions(perms ...Permission) bool {
	for _, perm := range perms {
		if valueMap.HasPermission(perm) {
			return true
		}
	}
	return false
}
