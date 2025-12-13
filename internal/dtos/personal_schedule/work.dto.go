package dtos

import "schedule_gateway/proto/personal_schedule"

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

type WorkDetailsResponseDTO struct {
	ID                  string                            `json:"id"`
	Name                string                            `json:"name"`
	ShortDescriptions   string                            `json:"short_descriptions"`
	DetailedDescription string                            `json:"detailed_description"`
	StartDate           int64                             `json:"start_date"`
	EndDate             int64                             `json:"end_date"`
	Goal                *GoalSimpleDTO                    `json:"goal"`
	Labels              *WorkLabelGroupDetail             `json:"labels"`
	SubTasks            *personal_schedule.SubTaskPayload `json:"sub_tasks"`
	Notifications       []*NotificationDTO                `protobuf:"bytes,6,opt,name=notifications,proto3" json:"notifications"`
}

type WorkLabelGroupDetail struct {
	Status     *personal_schedule.LabelInfo `protobuf:"bytes,1,opt,name=status,proto3" json:"status"`
	Difficulty *personal_schedule.LabelInfo `protobuf:"bytes,2,opt,name=difficulty,proto3" json:"difficulty"`
	Priority   *personal_schedule.LabelInfo `protobuf:"bytes,3,opt,name=priority,proto3" json:"priority"`
	Type       *personal_schedule.LabelInfo `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	Category   *personal_schedule.LabelInfo `protobuf:"bytes,5,opt,name=category,proto3" json:"category"`
}

type NotificationDTO struct {
	ID         string  `json:"id"`
	TriggerAt  int64   `json:"trigger_at"`
	IsSendMail bool    `json:"is_send_mail"`
	IsActive   bool    `json:"is_active"`
	Link       *string `json:"link"`
}
