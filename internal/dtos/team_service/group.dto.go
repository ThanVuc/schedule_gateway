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

type ListGroupItemDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	MyRole      int32  `json:"my_role"`
	MemberTotal int32  `json:"member_total"`
	AvatarURL   string `json:"avatar_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type MemberDTO struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Role     int32  `json:"role"`
	JoinedAt string `json:"joined_at"`
}

type UpdateMemberRoleDTO struct {
	NewRole int32 `json:"new_role" binding:"required"`
}

type CreateInviteDTO struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  int32  `json:"role"`
}

type InviteDTO struct {
	Code      string `json:"code"`
	ExpiresAt string `json:"expires_at"`
	CreateAt  string `json:"created_at"`
}

type CodeDataDTO struct {
	Code string `json:"code" binding:"required"`
}


type PresignFileItemDTO struct {
	Index       int32  `json:"index"`
	ContentType string `json:"content_type"`
	FileName    string `json:"file_name"`
}

type GeneratePresignedURLsRequest struct {
	Files []PresignFileItemDTO `json:"files"`
}

type PresignedFileItemResponseDTO struct {
	Index        int32  `json:"index"`
	PresignedURL string `json:"presigned_url"`
}

type GeneratePresignedURLsResponse struct {
	Files []PresignedFileItemResponseDTO `json:"files"`
}
