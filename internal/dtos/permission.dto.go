package dtos

type Permission struct {
	PermID      string `json:"perm_id"`
	PermName    string `json:"perm_name"`
	Description string `json:"description"`
	IsRoot      bool   `json:"is_root"`
}

type PermissionsResponse struct {
	Items             []Permission `json:"items"`
	TotalPermissions  int32        `json:"total_permissions"`
	Root              int32        `json:"root"`
	NonRoot           int32        `json:"non_root"`
	RootPercentage    float32      `json:"root_percentage"`
	NonRootPercentage float32      `json:"non_root_percentage"`
	TotalItems        int32        `json:"total_items"`
	TotalPages        int32        `json:"total_pages"`
	PageSize          int32        `json:"page_size"`
	Page              int32        `json:"page"`
	HasPrev           bool         `json:"has_prev"`
	HasNext           bool         `json:"has_next"`
}

type PermissionResponse struct {
	Data PermissionDetail `json:"data"`
}

type PermissionDetail struct {
	PermID      string   `json:"perm_id"`
	PermName    string   `json:"perm_name"`
	Description string   `json:"description"`
	IsRoot      bool     `json:"is_root"`
	CreatedAt   int64    `json:"created_at"`
	UpdatedAt   int64    `json:"updated_at"`
	Resource    Resource `json:"resource"`
	Action      []Action `json:"actions"`
}

type Resource struct {
	Id   string `json:"resource_id"`
	Name string `json:"resource_name"`
}

type Action struct {
	Id   string `json:"action_id"`
	Name string `json:"action_name"`
}
