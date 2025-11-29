package personalschedule_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/personalschedule"
	dtos "schedule_gateway/internal/dtos/personal_schedule"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/personal_schedule"
	"time"

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

	if req.StartDate != nil {
		if *req.StartDate > req.EndDate {
			response.BadRequest(c, "start_date must be before end_date")
			return nil
		}
	}

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

	result := dtos.WorksBySessionDTO{
		Morning:   []*personal_schedule.Work{},
		Noon:      []*personal_schedule.Work{},
		Afternoon: []*personal_schedule.Work{},
		Night:     []*personal_schedule.Work{},
		Evernight: []*personal_schedule.Work{},
	}

	for _, work := range resp.Works {
		h := time.Unix(work.GetStartDate(), 0).Hour()
		switch {
		case h >= 0 && h < 10:
			result.Morning = append(result.Morning, work)
		case h >= 10 && h < 14:
			result.Noon = append(result.Noon, work)
		case h >= 14 && h < 18:
			result.Afternoon = append(result.Afternoon, work)
		case h >= 18 && h < 22:
			result.Night = append(result.Night, work)
		default:
			result.Evernight = append(result.Evernight, work)
		}
	}

	response.Ok(c, "Ok", result)
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
		t := time.Unix(startDate, 0)
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

	wworkResp, err := wc.client.GetWork(c, req)
	if err != nil {
		wc.logger.Error("Connection error: ", "", zap.Error(err))
	}
	if wworkResp != nil && wworkResp.Error != nil {
		response.InternalServerError(c, utils.Int32PtrToString(wworkResp.Error.ErrorCode))
		return
	}
	if wworkResp == nil {
		response.InternalServerError(c, "Empty response from service")
		return
	}
	response.Ok(c, "Get Work Successful", wworkResp.Work)
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
        response.NotFound(c, "Error connecting to grpc service: " + err.Error())
		return
	}

	response.Ok(c, "Ok", resp.Goals)
}