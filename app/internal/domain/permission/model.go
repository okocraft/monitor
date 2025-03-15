package permission

type ID = int16

type Permission struct {
	ID           ID
	Name         string
	DefaultValue bool
}

type ValueMapSource = map[ID]bool

type ValueMap interface {
	IsTrue(id ID) bool
	HasPermission(perm Permission) bool
	HasAllPermissions(perms ...Permission) bool
	HasAnyPermissions(perms ...Permission) bool
	Len() int
	Iter(yield func(id ID, v bool) bool)
}

func EmptyValueMap() ValueMap {
	return valueMap{}
}

func NewValueMap(v ValueMapSource) ValueMap {
	return valueMap(v)
}

type valueMap map[ID]bool

func (valueMap valueMap) IsTrue(id ID) bool {
	v, ok := valueMap[id]
	return ok && v
}

func (valueMap valueMap) HasPermission(perm Permission) bool {
	if valueMap.IsTrue(Admin.ID) { // special case
		return true
	}

	if v, ok := valueMap[perm.ID]; ok {
		return v
	}
	return perm.DefaultValue
}

func (valueMap valueMap) HasAllPermissions(perms ...Permission) bool {
	if valueMap.IsTrue(Admin.ID) { // special case
		return true
	}

	for _, perm := range perms {
		if !valueMap.HasPermission(perm) {
			return false
		}
	}
	return true
}

func (valueMap valueMap) HasAnyPermissions(perms ...Permission) bool {
	if valueMap.IsTrue(Admin.ID) { // special case
		return true
	}

	for _, perm := range perms {
		if valueMap.HasPermission(perm) {
			return true
		}
	}
	return false
}

func (valueMap valueMap) Len() int {
	return len(valueMap)
}

func (valueMap valueMap) Iter(yield func(id ID, v bool) bool) {
	for k, v := range valueMap {
		if !yield(k, v) {
			return
		}
	}
}
