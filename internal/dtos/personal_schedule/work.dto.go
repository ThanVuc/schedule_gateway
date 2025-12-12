package dtos

type LabelInfoDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	Color     string `json:"color"`
	LabelType int32  `json:"label_type"`
}

type NotificationPayloadDTO struct {
	ID         *string `json:"id" binding:"omitempty,mongodb"`
	TriggerAt  int64   `json:"trigger_at" binding:"required"` // Unix timestamp
	IsSendMail bool    `json:"is_send_mail"`
	IsActive   bool    `json:"is_active"`
	Link       *string `json:"link"`
}

type WorkLabelsDTO struct {
	Status     *LabelInfoDTO `json:"status"`
	Difficulty *LabelInfoDTO `json:"difficulty"`
	Priority   *LabelInfoDTO `json:"priority"`
	Type       *LabelInfoDTO `json:"type"`
}

type WorksResponseDTO struct {
	ID                  string          `json:"id"`
	Name                string          `json:"name"`
	ShortDescriptions   string          `json:"short_descriptions"`
	DetailedDescription string          `json:"detailed_description"`
	StartDate           int64           `json:"start_date"`
	EndDate             int64           `json:"end_date"`
	Goal                *string         `json:"goal"`
	Labels              []*LabelInfoDTO `json:"labels"`
	Category            *LabelInfoDTO   `json:"category"`
}

type RecoveryWorksDTO struct {
	TargetDate int64  `json:"target_date"`
	SourceDate *int64 `json:"source_date"`
}

type UpdateWorkLabelDTO struct {
	LabelType int32  `json:"label_type" binding:"required"`
	LabelID   string `json:"label_id" binding:"required,mongodb"`
}