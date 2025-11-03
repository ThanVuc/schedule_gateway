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
