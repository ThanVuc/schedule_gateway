package dtos

import "schedule_gateway/proto/personal_schedule"

type Goals struct {
	Items      []*personal_schedule.Goal `json:"items"`
	TotalGoals int32                     `json:"total_goals"`
	TotalPages int32                     `json:"total_pages"`
	PageSize   int32                     `json:"page_size"`
	Page       int32                     `json:"page"`
	HasPrev    bool                      `json:"has_prev"`
	HasNext    bool                      `json:"has_next"`
}

type UpsertGoalDTO struct {
	ID                  *string         `json:"id"` // Dùng con trỏ để phân biệt rỗng
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
