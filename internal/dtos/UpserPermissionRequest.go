package dtos

type UpsertPermissionRequestDTO struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	ResourceId  string   `json:"resource_id"`
	ActionsIds  []string `json:"actions_ids"`
}
