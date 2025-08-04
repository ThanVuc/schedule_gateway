package dtos

type UpsertRoleRequestDTO struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	PermissionIds []string `json:"permission_ids"`
}
