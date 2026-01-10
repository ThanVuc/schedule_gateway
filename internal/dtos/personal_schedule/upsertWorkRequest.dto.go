package dtos

type SubTaskPayloadDTO struct {
	ID          *string `json:"id" binding:"omitempty,mongodb"`
	Name        string  `json:"name" binding:"required"`
	IsCompleted bool    `json:"is_completed"`
}

type UpsertWorkDTO struct {
	Name                string                   `json:"name" binding:"required"`
	ShortDescriptions   *string                  `json:"short_descriptions"`
	DetailedDescription *string                  `json:"detailed_description"`
	StartDate           *int64                   `json:"start_date"`
	EndDate             int64                    `json:"end_date" binding:"required"`
	StatusID            string                   `json:"status_id" binding:"required,mongodb"`
	DifficultyID        string                   `json:"difficulty_id" binding:"required,mongodb"`
	PriorityID          string                   `json:"priority_id" binding:"required,mongodb"`
	TypeID              string                   `json:"type_id" binding:"required,mongodb"`
	CategoryID          string                   `json:"category_id" binding:"required,mongodb"`
	DraftID             *string                  `json:"draft_id" binding:"omitempty,mongodb"`
	GoalID              *string                  `json:"goal_id"`
	SubTasks            []SubTaskPayloadDTO      `json:"sub_tasks"`
	Notifications       []NotificationPayloadDTO `json:"notifications"`
	UpdateType          *int32                   `json:"update_type" binding:"omitempty"`
	RepeatStartDate     *int64                   `json:"repeat_start_date" binding:"omitempty"`
	RepeatEndDate       *int64                   `json:"repeat_end_date" binding:"omitempty"`
}
