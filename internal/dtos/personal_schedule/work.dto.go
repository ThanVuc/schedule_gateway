package dtos

import "schedule_gateway/proto/personal_schedule"

type SubTaskPayloadDTO struct {
	ID          *string `json:"id" binding:"omitempty,mongodb"`
	Name        string  `json:"name" binding:"required"`
	IsCompleted bool    `json:"is_completed"`
}

type UpsertWorkDTO struct {
	Name                string  `json:"name" binding:"required"`
	ShortDescriptions   *string `json:"short_descriptions"`
	DetailedDescription *string `json:"detailed_description"`
	StartDate           *int64  `json:"start_date"`
	EndDate             int64   `json:"end_date" binding:"required"`
	StatusID            string  `json:"status_id" binding:"required,mongodb"`
	DifficultyID        string  `json:"difficulty_id" binding:"required,mongodb"`
	PriorityID          string  `json:"priority_id" binding:"required,mongodb"`
	TypeID              string  `json:"type_id" binding:"required,mongodb"`
	CategoryID          string  `json:"category_id" binding:"required,mongodb"`
	GoalID              *string `json:"goal_id"`

	SubTasks        []SubTaskPayloadDTO `json:"sub_tasks"`
	NotificationIds []string            `json:"notification_ids" binding:"omitempty,dive,mongodb"`
}

type WorksBySessionDTO struct {
	Morning   []*personal_schedule.Work `json:"morning"`
	Noon      []*personal_schedule.Work `json:"noon"`
	Afternoon []*personal_schedule.Work `json:"afternoon"`
	Night     []*personal_schedule.Work `json:"night"`
	Evernight []*personal_schedule.Work `json:"evernight"`
}
