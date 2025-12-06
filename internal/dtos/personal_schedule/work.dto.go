package dtos

type LabelInfoDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Color     string `json:"color"`
	LabelType int32  `json:"label_type"`
}

type NotificationPayloadDTO struct {
	ID          *string `json:"id" binding:"omitempty,mongodb"`
	TriggerAt   int64   `json:"trigger_at" binding:"required"` // Unix timestamp
	IsEmailSent bool    `json:"is_email_sent"`
	IsActive    bool    `json:"is_active"`
	Link        *string `json:"link"`
}

type GoalInfoDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type SubTaskResponseDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"is_completed"`
}

type WorkLabelsDTO struct {
	Status     *LabelInfoDTO `json:"status"`
	Difficulty *LabelInfoDTO `json:"difficulty"`
	Priority   *LabelInfoDTO `json:"priority"`
	Type       *LabelInfoDTO `json:"type"`
}

type WorkLabelsDetailDTO struct {
	Status     *LabelInfoDTO `json:"status"`
	Difficulty *LabelInfoDTO `json:"difficulty"`
	Priority   *LabelInfoDTO `json:"priority"`
	Type       *LabelInfoDTO `json:"type"`
	Category   *LabelInfoDTO `json:"category"`
}

type WorksResponseDTO struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	ShortDescriptions   string        `json:"short_descriptions"`
	DetailedDescription string        `json:"detailed_description"`
	StartDate           int64         `json:"start_date"`
	EndDate             int64         `json:"end_date"`
	Goal                *string       `json:"goal"`
	Labels              WorkLabelsDTO `json:"labels"`
	Category            *LabelInfoDTO `json:"category"`
}

type WorksDetailDTO struct {
	ID                  string               `json:"id"`
	Name                string               `json:"name"`
	ShortDescriptions   string               `json:"short_descriptions"`
	DetailedDescription string               `json:"detailed_description"`
	StartDate           int64                `json:"start_date"`
	EndDate             int64                `json:"end_date"`
	Goal                *string              `json:"goal"`
	Labels              WorkLabelsDetailDTO  `json:"labels"`
	Category            *LabelInfoDTO        `json:"category"`
	SubTasks            []SubTaskResponseDTO `json:"sub_tasks"`
}
