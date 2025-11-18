package personalschedule_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/personalschedule"
	dtos "schedule_gateway/internal/dtos/personal_schedule"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/personal_schedule"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type WorkController struct {
	logger log.Logger
	client client.WorkClient
}

func NewWorkController() *WorkController {
	return &WorkController{
		logger: global.Logger,
		client: client.NewWorkClient(),
	}
}

func (wc *WorkController) UpsertWork(c *gin.Context) {
	req := wc.buildUpsertWorkRequest(c)
	if req == nil {
		return
	}
	upsertResp, err := wc.client.UpsertWork(c, req)
	if err != nil {
		wc.logger.Error("Connection error: ", "", zap.Error(err))
	}
	if upsertResp != nil && upsertResp.Error != nil {
		response.InternalServerError(c, upsertResp.Error.Message)
		return
	}
	if upsertResp == nil {
		response.InternalServerError(c, "Empty response from service")
		return
	}

	if !upsertResp.IsSuccess {
		response.InternalServerError(c, "Upsert Work Failed")
		return
	}

	response.Ok(c, "Upsert Work Successful", gin.H{
		"is_success": upsertResp.IsSuccess,
	})
}

func (wc *WorkController) buildUpsertWorkRequest(c *gin.Context) *personal_schedule.UpsertWorkRequest {
	var req personal_schedule.UpsertWorkRequest
	var dto dtos.UpsertWorkDTO

	UserID := c.GetString("user_id")
	if UserID == "" {
		response.BadRequest(c, "user_id is required")
		return nil
	}

	id := c.Param("id")
	if id != "" {
		req.Id = &id
	} else {
		req.Id = nil
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, "Invalid request body: "+err.Error())
		return nil
	}
	req.UserId = UserID
	req.Name = dto.Name
	req.ShortDescriptions = dto.ShortDescriptions
	req.DetailedDescription = dto.DetailedDescription
	req.StartDate = dto.StartDate
	req.EndDate = dto.EndDate
	req.StatusId = dto.StatusID
	req.DifficultyId = dto.DifficultyID
	req.PriorityId = dto.PriorityID
	req.TypeId = dto.TypeID
	req.CategoryId = dto.CategoryID
	req.GoalId = dto.GoalID
	req.NotificationIds = dto.NotificationIds
	
	req.SubTasks = make([]*personal_schedule.SubTaskPayload, len(dto.SubTasks))
	for i, subTaskDto := range dto.SubTasks {
		var subTaskID string
		if subTaskDto.ID != nil {
			subTaskID = *subTaskDto.ID
		}
		req.SubTasks[i] = &personal_schedule.SubTaskPayload{
			Id:          &subTaskID,
			Name:        subTaskDto.Name,
			IsCompleted: subTaskDto.IsCompleted,
		}
	}

	return &req
}
