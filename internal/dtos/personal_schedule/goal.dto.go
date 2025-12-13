package dtos

type Goals struct {
	Items      []GoalItemDTO `json:"items"`
	TotalGoals int32         `json:"total_goals"`
	TotalPages int32         `json:"total_pages"`
	PageSize   int32         `json:"page_size"`
	Page       int32         `json:"page"`
	HasPrev    bool          `json:"has_prev"`
	HasNext    bool          `json:"has_next"`
}

type UpsertGoalDTO struct {
	ID                  *string         `json:"id"`
	Name                string          `json:"name" binding:"required"`
	ShortDescriptions   *string         `json:"short_descriptions"`
	DetailedDescription *string         `json:"detailed_description"`
	StartDate           *int64          `json:"start_date"`
	EndDate             *int64          `json:"end_date"`
	StatusID            string          `json:"status_id" binding:"required"`
	DifficultyID        string          `json:"difficulty_id" binding:"required"`
	PriorityID          string          `json:"priority_id" binding:"required"`
	CategoryID          string          `json:"category_id" binding:"required"`
	Tasks               []UpsertTaskDTO `json:"tasks"`
}
type UpsertTaskDTO struct {
	ID          *string `json:"id"`
	Name        string  `json:"name" binding:"required"`
	IsCompleted bool    `json:"is_completed"`
}

type GoalItemDTO struct {
	ID                  string          `json:"id"`
	Name                string          `json:"name"`
	ShortDescriptions   *string         `json:"short_descriptions"`
	DetailedDescription *string         `json:"detailed_description"`
	StartDate           int64           `json:"start_date"`
	EndDate             int64           `json:"end_date"`
	Labels              []*LabelInfoDTO `json:"labels"`
	Category            *LabelInfoDTO   `json:"category"`
}

type GoalSimpleDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UpdateGoalLabelDTO struct {
	LabelType int32  `json:"label_type" binding:"required"`
	LabelID   string `json:"label_id" binding:"required,mongodb"`
}
