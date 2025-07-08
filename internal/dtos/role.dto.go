package dtos

type Role struct {
	RoleID      string `json:"role_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsRoot      bool   `json:"is_root"`
	TotalUsers  int32  `json:"total_users"`
	IsActive    bool   `json:"is_active"`
}

type RolesResponse struct {
	Items             []Role  `json:"items"`
	TotalRoles        int32   `json:"total_roles"`
	Root              int32   `json:"root"`
	NonRoot           int32   `json:"non_root"`
	RootPercentage    float32 `json:"root_percentage"`
	NonRootPercentage float32 `json:"non_root_percentage"`
	TotalItems        int32   `json:"total_items"`
	TotalPages        int32   `json:"total_pages"`
	PageSize          int32   `json:"page_size"`
	Page              int32   `json:"page"`
	HasPrev           bool    `json:"has_prev"`
	HasNext           bool    `json:"has_next"`
}

type RoleResponse struct {
	Data RoleDetail `json:"data"`
}

type RoleDetail struct {
	RoleID      string     `json:"role_id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	IsRoot      bool       `json:"is_root"`
	IsActive    bool       `json:"is_active"`
	Permissions []PermRole `json:"permissions"`
	CreatedAt   int64      `json:"created_at"`
	UpdatedAt   int64      `json:"updated_at"`
}

type PermRole struct {
	PermId      string `json:"perm_id"`
	PermName    string `json:"perm_name"`
	Description string `json:"description"`
}
