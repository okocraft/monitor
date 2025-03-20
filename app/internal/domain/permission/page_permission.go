package permission

import "github.com/okocraft/monitor/internal/handler/oapi"

type PagePermissions struct {
	Settings SettingPagePermissions
}

type SettingPagePermissions struct {
	Users bool
}

func (p PagePermissions) ToResponse() oapi.PagePermissions {
	return oapi.PagePermissions{
		Settings: oapi.SettingPagePermissions{
			Users: p.Settings.Users,
		},
	}
}

type PagePermissionCalculator interface {
	GetSourcePermissions() []Permission
	Calculate(valueMap ValueMap) PagePermissions
}

type pagePermissionCalculator struct {
	sourcePermissions []Permission
	calculator        func(valueMap ValueMap) PagePermissions
}

func (p pagePermissionCalculator) GetSourcePermissions() []Permission {
	return p.sourcePermissions
}

func (p pagePermissionCalculator) Calculate(valueMap ValueMap) PagePermissions {
	return p.calculator(valueMap)
}

var impl = pagePermissionCalculator{
	sourcePermissions: []Permission{
		UserList,
	},
	calculator: func(valueMap ValueMap) PagePermissions {
		return PagePermissions{
			Settings: SettingPagePermissions{
				Users: valueMap.HasPermission(UserList),
			},
		}
	},
}

func GetPagePermissionCalculator() PagePermissionCalculator {
	return impl
}
