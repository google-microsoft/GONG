package response

import "github.com/richardgong1987/server/model/system"

type SysMenusResponse struct {
	Menus []system.SysMenu `json:"menus"`
}

type SysBaseMenusResponse struct {
	Menus []system.SysBaseMenu `json:"menus"`
}

type SysBaseMenuResponse struct {
	Menu system.SysBaseMenu `json:"menu"`
}
