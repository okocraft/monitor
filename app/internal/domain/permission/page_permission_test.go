package permission

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type checkingValueMap struct {
	valueMap
	source   map[Permission]struct{}
	notFound map[Permission]struct{}
	unused   map[Permission]struct{}
}

func (valueMap checkingValueMap) HasPermission(permission Permission) bool {
	if _, ok := valueMap.source[permission]; ok {
		delete(valueMap.unused, permission)
	} else {
		valueMap.notFound[permission] = struct{}{}
	}
	return true // admin
}

func TestPagePermissionImpl(t *testing.T) {
	perms := impl.sourcePermissions
	source := make(map[Permission]struct{}, len(perms))
	notFound := make(map[Permission]struct{}, len(perms))
	unused := make(map[Permission]struct{}, len(perms))

	for _, v := range perms {
		source[v] = struct{}{}
		unused[v] = struct{}{}
	}

	vMap := checkingValueMap{
		source:   source,
		notFound: notFound,
		unused:   unused,
	}

	_ = impl.Calculate(vMap)

	for perm := range vMap.notFound {
		assert.Fail(t, perm.Name+" is not found in source permissions")
	}

	for perm := range vMap.unused {
		assert.Fail(t, perm.Name+" is not used")
	}
}
