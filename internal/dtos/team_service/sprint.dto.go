package dtos

type CreateSprintDTO struct {
	Name      string  `json:"name" binding:"required"`
	Goal      *string `json:"goal,omitempty"`
	StartDate string  `json:"start_date" binding:"required"`
	EndDate   string  `json:"end_date" binding:"required"`
}

type UpdateSprintDTO struct {
	Name      *string `json:"name,omitempty"`
	Goal      *string `json:"goal,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	EndDate   *string `json:"end_date,omitempty"`
}
