package request

type RoleMenus struct {
	RoleId  uint   `json:"role_id"`
	MenuIds []uint `json:"menu_ids"`
}
