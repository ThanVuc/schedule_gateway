package personalschedule_controller

import (
	"fmt"
	"schedule_gateway/global"
	notification_client "schedule_gateway/internal/client/notification"
	client "schedule_gateway/internal/client/personalschedule"
	dtos "schedule_gateway/internal/dtos/personal_schedule"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/notification_service"
	"schedule_gateway/proto/personal_schedule"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type WorkController struct {
	logger             log.Logger
	client             client.WorkClient
	notificationClient notification_client.NotificationClient
}

func NewWorkController() *WorkController {
	return &WorkController{
		logger:             global.Logger,
		client:             client.NewWorkClient(),
		notificationClient: notification_client.NewNotificationClient(),
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
		if upsertResp.Error != nil && upsertResp.Error.ErrorCode != nil {
			response.ValidationError(c, upsertResp.Error.Message, utils.Int32PtrToString(upsertResp.Error.ErrorCode))
			return
		}
		fmt.Println("Error Message:", upsertResp.Error.Message)
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
	notifications := make([]*personal_schedule.WorkNotification, len(dto.Notifications))
	for i, notificationDto := range dto.Notifications {
		var notificationID string
		if notificationDto.ID != nil {
			notificationID = *notificationDto.ID
		}
		notifications[i] = &personal_schedule.WorkNotification{
			Id:         &notificationID,
			TriggerAt:  notificationDto.TriggerAt,
			IsSendMail: notificationDto.IsSendMail,
			IsActive:   notificationDto.IsActive,
			Link:       notificationDto.Link,
			ImgUrl:     notificationDto.ImgUrl,
		}
	}
	req.Notifications = notifications

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

func (wc *WorkController) GetWorks(c *gin.Context) {
	req := wc.buildGetWorksRequest(c)
	if req == nil {
		return
	}

	resp, err := wc.client.GetWorks(c, req)
	if err != nil {
		wc.logger.Error("Connection error: ", "", zap.Error(err))
		response.InternalServerError(c, "Error connecting to grpc service")
		return
	}
	if resp != nil && resp.Error != nil {
		if resp.Error != nil && resp.Error.ErrorCode != nil {
			response.InternalServerError(c, utils.Int32PtrToString(resp.Error.ErrorCode))
			return
		}
		response.InternalServerError(c, resp.Error.Message)
		return
	}

	result := make([]dtos.WorksResponseDTO, 0)
	if resp != nil && resp.Works != nil {
		for _, w := range resp.Works {
			result = append(result, wc.mapProtoToDTO(w))
		}
	}

	response.Ok(c, "Get Works Successful", gin.H{
		"works": result,
	})
}

func (wc *WorkController) mapProtoToDTO(p *personal_schedule.Work) dtos.WorksResponseDTO {
	mapLabel := func(l *personal_schedule.LabelInfo) *dtos.LabelInfoDTO {
		if l == nil {
			return nil
		}
		return &dtos.LabelInfoDTO{
			ID:        l.Id,
			Name:      l.Name,
			Key:       l.Key,
			Color:     l.Color,
			LabelType: l.LabelType,
		}
	}
	var goalName *string
	if p.Goal != nil {
		name := p.Goal.Name
		goalName = &name
	}
	sd := ""
	if p.ShortDescriptions != nil {
		sd = *p.ShortDescriptions
	}
	dd := ""
	if p.DetailedDescription != nil {
		dd = *p.DetailedDescription
	}

	labels := make([]*dtos.LabelInfoDTO, 0)
	if p.Labels != nil {
		if p.Labels.Status != nil {
			labels = append(labels, mapLabel(p.Labels.Status))
		}
		if p.Labels.Difficulty != nil {
			labels = append(labels, mapLabel(p.Labels.Difficulty))
		}
		if p.Labels.Priority != nil {
			labels = append(labels, mapLabel(p.Labels.Priority))
		}
		if p.Labels.Type != nil {
			labels = append(labels, mapLabel(p.Labels.Type))
		}

	}

	return dtos.WorksResponseDTO{
		ID:                  p.Id,
		Name:                p.Name,
		ShortDescriptions:   sd,
		DetailedDescription: dd,
		StartDate:           p.StartDate,
		EndDate:             p.EndDate,
		Goal:                goalName,
		Category:            mapLabel(p.Category),
		Labels:              labels,
	}
}

func (wc *WorkController) buildGetWorksRequest(c *gin.Context) *personal_schedule.GetWorksRequest {
	userID := c.GetString("user_id")
	if userID == "" {
		response.BadRequest(c, "user_id is required")
		return nil
	}

	search := c.Query("search")
	statusID := c.Query("status_id")
	difficultyID := c.Query("difficulty_id")
	priorityID := c.Query("priority_id")
	typeID := c.Query("type_id")
	categoryID := c.Query("category_id")
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	var startDate, endDate int64
	now := time.Now()
	fmt.Println("Current time UTC:", now)
	if startDateStr != "" {
		parsed, err := utils.ParseStringToInt64(startDateStr)
		if err != nil {
			response.BadRequest(c, "Invalid start_date format")
			return nil
		}
		startDate = parsed
	} else {
		startDate, _ = utils.StartAndEndOfDayTimestamp(now)
	}

	if endDateStr != "" {
		parsed, err := utils.ParseStringToInt64(endDateStr)
		if err != nil {
			response.BadRequest(c, "Invalid end_date format")
			return nil
		}
		endDate = parsed
	} else {
		t := time.UnixMilli(startDate)
		_, endDate = utils.StartAndEndOfDayTimestamp(t)
	}

	if startDate > endDate {
		response.BadRequest(c, "start_date must be before end_date")
		return nil
	}

	req := &personal_schedule.GetWorksRequest{
		UserId:       userID,
		Search:       &search,
		FromDate:     &startDate,
		ToDate:       &endDate,
		StatusId:     &statusID,
		DifficultyId: &difficultyID,
		PriorityId:   &priorityID,
		TypeId:       &typeID,
		CategoryId:   &categoryID,
	}

	return req
}

func (wc *WorkController) GetWork(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Work ID is required")
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		response.BadRequest(c, "user_id is required")
		return
	}

	req := &personal_schedule.GetWorkRequest{
		UserId: userID,
		WorkId: id,
	}

	workResp, err := wc.client.GetWork(c, req)
	if err != nil {
		wc.logger.Error("Connection error: ", "", zap.Error(err))
	}
	if workResp != nil && workResp.Error != nil {
		response.InternalServerError(c, utils.Int32PtrToString(workResp.Error.ErrorCode))
		return
	}
	if workResp == nil {
		response.InternalServerError(c, "Empty response from service")
		return
	}

	notificationsByWorkResp, err := wc.notificationClient.GetNotificationByWorkId(c, &common.IDRequest{
		Id: id,
	})

	if err != nil {
		wc.logger.Error("Connection error: ", "", zap.Error(err))
		response.InternalServerError(c, "Error connecting to notification service")
		return
	}

	if notificationsByWorkResp != nil && notificationsByWorkResp.Error != nil {
		response.InternalServerError(c, notificationsByWorkResp.Error.Message)
		return
	}

	workDetailDTO := wc.buildWorkDetailResponse(workResp.Work, notificationsByWorkResp.Notifications)

	response.Ok(c, "Get Work Successful", workDetailDTO)
}

func (wc *WorkController) buildWorkDetailResponse(work *personal_schedule.WorkDetail, notifications []*notification_service.WorkNotification) *dtos.WorkDetailsResponseDTO {
	var goalDTO *dtos.GoalSimpleDTO
	if work.Goal != nil {
		goalDTO = &dtos.GoalSimpleDTO{
			ID:   work.Goal.Id,
			Name: work.Goal.Name,
		}
	}

	notificationsDTO := make([]*dtos.NotificationDTO, 0)
	for _, n := range notifications {
		notificationsDTO = append(notificationsDTO, &dtos.NotificationDTO{
			ID:         n.Id,
			TriggerAt:  n.TriggerAt,
			IsSendMail: n.IsSendMail,
			IsActive:   n.IsActive,
			Link:       &n.Link,
		})
	}

	workDTO := &dtos.WorkDetailsResponseDTO{
		ID:                  work.Id,
		Name:                work.Name,
		ShortDescriptions:   utils.SafeString(work.ShortDescriptions),
		DetailedDescription: utils.SafeString(work.DetailedDescription),
		StartDate:           work.StartDate,
		EndDate:             work.EndDate,
		Goal:                goalDTO,
		Notifications:       notificationsDTO,
	}

	return workDTO
}

func (wc *WorkController) DeleteWork(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Work ID is required")
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		response.BadRequest(c, "user_id is required")
		return
	}

	req := &personal_schedule.DeleteWorkRequest{
		UserId: userID,
		WorkId: id,
	}
	deleteResp, err := wc.client.DeleteWork(c, req)
	if err != nil {
		wc.logger.Error("Connection error: ", "", zap.Error(err))
	}

	if deleteResp != nil && deleteResp.Error != nil {
		response.InternalServerError(c, utils.Int32PtrToString(deleteResp.Error.ErrorCode))
		return
	}
	if deleteResp == nil {
		response.InternalServerError(c, "Empty response from service")
		return
	}
	response.Ok(c, "Delete Work Successful", gin.H{
		"is_success": deleteResp.Success,
		"error":      deleteResp.Error,
	})
}

func (gc *GoalController) GetGoalsForDialog(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		response.BadRequest(c, "user_id is required")
		return
	}

	resp, err := gc.client.GetGoalForDiaglog(c, &personal_schedule.GetGoalsForDialogRequest{
		UserId: userID,
	})

	if err != nil {
		response.NotFound(c, "Error connecting to grpc service: "+err.Error())
		return
	}

	response.Ok(c, "Ok", resp.Goals)
}

func (wc *WorkController) GetRecoveryWorks(c *gin.Context) {
	userID := c.GetString("user_id")
	if userID == "" {
		response.BadRequest(c, "user_id is required")
		return
	}
	var dto dtos.RecoveryWorksDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		wc.logger.Error("Failed to bind JSON: ", "", zap.Error(err))
		response.BadRequest(c, "Invalid request body: "+err.Error())
		return
	}

	targetDate := dto.TargetDate
	var sourceDate int64
	if dto.SourceDate != nil && *dto.SourceDate != 0 {
		sourceDate = *dto.SourceDate
	} else {
		t := time.Unix(targetDate, 0).AddDate(0, 0, -1)
		sourceDate, _ = utils.StartAndEndOfDayTimestamp(t)
	}
	req := &personal_schedule.GetRecoveryWorksRequest{
		UserId:     userID,
		TargetDate: targetDate,
		SourceDate: sourceDate,
	}
	resp, err := wc.client.GetRecoveryWorks(c, req)
	if err != nil {
		wc.logger.Error("Connection error: ", "", zap.Error(err))
		response.InternalServerError(c, "Error connecting to grpc service")
		return
	}
	response.Ok(c, "Get Recovery Works Successful", resp.Works)
}

func (wc *WorkController) UpdateWorkLabel(c *gin.Context) {
	userID := c.GetString("user_id")
	workID := c.Param("id")

	var dto dtos.UpdateWorkLabelDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		wc.logger.Error("Failed to bind JSON: ", "", zap.Error(err))
		response.BadRequest(c, "Invalid body")
		return
	}

	req := &personal_schedule.UpdateWorkLabelRequest{
		UserId:    userID,
		WorkId:    workID,
		LabelType: dto.LabelType,
		LabelId:   dto.LabelID,
	}

	resp, err := wc.client.UpdateWorkLabel(c, req)
	if err != nil {
		wc.logger.Error("Connection error: ", "", zap.Error(err))
		response.InternalServerError(c, "Error connecting to grpc service")
		return
	}
	if resp != nil && resp.Error != nil {
		response.InternalServerError(c, resp.Error.Message)
		return
	}

	response.Ok(c, "Updated", gin.H{"is_success": true})
}
