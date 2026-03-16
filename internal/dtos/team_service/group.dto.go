package dtos

type CreateGroupDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
}

type GroupDTO struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Owner        *SimpleUserDTO `json:"owner"`
	MyRole       int32          `json:"my_role"`
	ActiveSprint string         `json:"active_sprint"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
}

type SimpleUserDTO struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type GroupDetailDTO struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	Owner        *SimpleUserDTO `json:"owner"`
	MyRole       int32          `json:"my_role"`
	ActiveSprint string         `json:"active_sprint"`
	Avatar       string         `json:"avatar"`
	MembersTotal int32          `json:"members_total"`
	CreatedAt    string         `json:"created_at"`
	UpdatedAt    string         `json:"updated_at"`
}
