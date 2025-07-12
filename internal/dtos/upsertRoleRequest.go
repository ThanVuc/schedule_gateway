package dtos

type UpsertRoleRequestDTO struct {
	RoleId        *string  `json:"role_id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIds []string `json:"permission_ids"`
}
