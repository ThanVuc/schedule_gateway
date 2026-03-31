package dtos

type CreateWorkDTO struct {
	SprintID    *string `json:"sprint_id",omitempty`
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description,omitempty"`
}

type UpdateWorkDTO struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	AssigneeID  *string `json:"assignee_id,omitempty"`
	Status      *int32  `json:"status,omitempty"`
	StoryPoint  *int32  `json:"story_point,omitempty"`
	DueDate     *string `json:"due_date,omitempty"`
	Priority    *int32  `json:"priority,omitempty"`
	Version     *int32  `json:"version" binding:"required"`
}

type CreateChecklistItemDTO struct {
	Name string `json:"name" binding:"required"`
}

type UpdateChecklistItemDTO struct {
	Name        *string `json:"name,omitempty"`
	IsCompleted *bool   `json:"is_completed,omitempty"`
}

type CreateCommentDTO struct {
	Content string `json:"content" binding:"required"`
}

type UpdateCommentDTO struct {
	Content string `json:"content" binding:"required"`
}
