package personalschedule_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/personalschedule"
	dtos "schedule_gateway/internal/dtos/personal_schedule"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/personal_schedule"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type GoalController struct {
	logger log.Logger
	client client.GoalClient
}

func NewGoalController() *GoalController {
	return &GoalController{
		logger: global.Logger,
		client: client.NewGoalClient(),
	}
}

func (gc *GoalController) GetGoals(ctx *gin.Context) {
	req := gc.buildGetGoalsRequest(ctx)
	if req == nil {
		return
	}
	goalsResp, err := gc.client.GetGoals(ctx, req)

	if err != nil {
		gc.logger.Error("Connection error: ", "", zap.Error(err))
	}

	if goalsResp != nil && goalsResp.Error != nil {
		response.InternalServerError(ctx, utils.Int32PtrToString(goalsResp.Error.ErrorCode))
		return
	}
	if goalsResp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}
	items := make([]dtos.GoalItemDTO, 0, len(goalsResp.Goals))
	for _, g := range goalsResp.Goals {
		items = append(items, gc.mapProtoToDTO(g))
	}

	response.Ok(ctx, "Get Goals Successful", dtos.Goals{
		Items:      items,
		TotalGoals: goalsResp.TotalGoals,
		TotalPages: goalsResp.PageInfo.TotalPages,
		PageSize:   goalsResp.PageInfo.PageSize,
		Page:       goalsResp.PageInfo.Page,
		HasPrev:    goalsResp.PageInfo.HasPrev,
		HasNext:    goalsResp.PageInfo.HasNext,
	})
}

func (gc *GoalController) mapProtoToDTO(p *personal_schedule.Goal) dtos.GoalItemDTO {
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
	sd := ""
	if p.ShortDescriptions != nil {
		sd = *p.ShortDescriptions
	}
	dd := ""
	if p.DetailedDescription != nil {
		dd = *p.DetailedDescription
	}

	labels := make([]*dtos.LabelInfoDTO, 0)
	if p.GoalLabels != nil {
		if p.GoalLabels.Status != nil {
			labels = append(labels, mapLabel(p.GoalLabels.Status))
		}
		if p.GoalLabels.Difficulty != nil {
			labels = append(labels, mapLabel(p.GoalLabels.Difficulty))
		}
		if p.GoalLabels.Priority != nil {
			labels = append(labels, mapLabel(p.GoalLabels.Priority))
		}

	}

	return dtos.GoalItemDTO{
		ID:                  p.Id,
		Name:                p.Name,
		ShortDescriptions:   &sd,
		DetailedDescription: &dd,
		StartDate:           p.StartDate,
		EndDate:             p.EndDate,
		Category:            mapLabel(p.Category),
		Labels:              labels,
		Overdue:             mapLabel(p.Overdue),
	}
}

func (gc *GoalController) buildGetGoalsRequest(c *gin.Context) *personal_schedule.GetGoalsRequest {
	userID := c.GetString("user_id")
	if userID == "" {
		response.BadRequest(c, "user_id is required")
		return nil
	}

	pageQuery := utils.ToPageQuery(c)
	searchString := c.Query("search")
	statusID := c.Query("status_id")

	req := personal_schedule.GetGoalsRequest{
		UserId:    userID,
		PageQuery: pageQuery,
		Search:    &searchString,
		StatusId:  &statusID,
	}
	return &req
}

func (gc *GoalController) UpsertGoal(c *gin.Context) {
	req := gc.buildUpsertGoalRequest(c)
	if req == nil {
		return
	}
	upsertResp, err := gc.client.UpsertGoals(c, req)

	if err != nil {
		gc.logger.Error("Connection error: ", "", zap.Error(err))
	}
	if upsertResp != nil && upsertResp.Error != nil {
		response.InternalServerError(c, utils.Int32PtrToString(upsertResp.Error.ErrorCode))
		return
	}
	if upsertResp == nil {
		response.InternalServerError(c, "Empty response from service")
		return
	}

	if !upsertResp.IsSuccess {
		response.InternalServerError(c, "Upsert Goal Failed")
		return
	}

	response.Ok(c, "Upsert Goal Successful", gin.H{
		"is_success": upsertResp.IsSuccess,
	})

}

func (gc *GoalController) buildUpsertGoalRequest(c *gin.Context) *personal_schedule.UpsertGoalRequest {
	var req personal_schedule.UpsertGoalRequest
	var dto dtos.UpsertGoalDTO

	userID := c.GetString("user_id")
	if userID == "" {
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
		gc.logger.Error("Failed to bind JSON: ", "", zap.Error(err))
		response.BadRequest(c, "Invalid request body: "+err.Error())
		return nil
	}

	req.UserId = userID
	req.Name = dto.Name
	req.DetailedDescription = dto.DetailedDescription
	req.StatusId = dto.StatusID
	req.DifficultyId = dto.DifficultyID
	req.PriorityId = dto.PriorityID
	req.ShortDescriptions = dto.ShortDescriptions
	req.CategoryId = dto.CategoryID
	req.StartDate = dto.StartDate
	req.EndDate = dto.EndDate
	req.Tasks = make([]*personal_schedule.GoalTaskPayload, len(dto.Tasks))

	if *req.StartDate > *req.EndDate {
		response.BadRequest(c, "start_date must be before end_date")
		return nil
	}

	if *req.StartDate == *req.EndDate {
		response.BadRequest(c, "start_date must not be equal to end_date")
		return nil
	}

	for i, taskDTO := range dto.Tasks {
		var taskID string
		if taskDTO.ID != nil {
			taskID = *taskDTO.ID
		}
		req.Tasks[i] = &personal_schedule.GoalTaskPayload{
			Id:          &taskID,
			Name:        taskDTO.Name,
			IsCompleted: taskDTO.IsCompleted,
		}
	}

	return &req
}

func (gc *GoalController) GetGoal(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Goal ID is required")
		return
	}

	userID := c.GetString("user_id")
	if userID == "" {
		response.BadRequest(c, "User ID is required")
		return
	}

	req := &personal_schedule.GetGoalRequest{
		GoalId: id,
		UserId: userID,
	}

	goalResp, err := gc.client.GetGoal(c, req)
	if err != nil {
		gc.logger.Error("Connection error: ", "", zap.Error(err))
	}
	if goalResp != nil && goalResp.Error != nil {
		response.InternalServerError(c, utils.Int32PtrToString(goalResp.Error.ErrorCode))
		return
	}
	if goalResp == nil {
		response.InternalServerError(c, "Empty response from service")
		return
	}
	response.Ok(c, "Get Goal Successful", goalResp.Goal)
}

func (gc *GoalController) DeleteGoal(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, "Goal ID is required")
		return
	}
	userID := c.GetString("user_id")
	if userID == "" {
		response.BadRequest(c, "User ID is required")
		return
	}

	req := &personal_schedule.DeleteGoalRequest{
		GoalId: id,
		UserId: userID,
	}

	deleteRes, err := gc.client.DeleteGoal(c, req)
	if err != nil {
		gc.logger.Error("Connection error: ", "", zap.Error(err))
	}
	if deleteRes != nil && deleteRes.Error != nil {
		response.InternalServerError(c, utils.Int32PtrToString(deleteRes.Error.ErrorCode))
		return
	}
	if deleteRes == nil {
		response.InternalServerError(c, "Empty response from service")
		return
	}

	response.Ok(c, "Delete Goal Successful", gin.H{
		"is_success": deleteRes.Success,
	})

}

func (gc *GoalController) UpdateGoalLabel(c *gin.Context) {
	userID := c.GetString("user_id")
	goalID := c.Param("id")

	var dto dtos.UpdateGoalLabelDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		gc.logger.Error("Failed to bind JSON: ", "", zap.Error(err))
		response.BadRequest(c, "Invalid body")
		return
	}

	req := &personal_schedule.UpdateGoalLabelRequest{
		UserId:    userID,
		GoalId:    goalID,
		LabelType: dto.LabelType,
		LabelId:   dto.LabelID,
	}

	resp, err := gc.client.UpdateGoalLabel(c, req)
	if err != nil {
		gc.logger.Error("Connection error: ", "", zap.Error(err))
		response.InternalServerError(c, "Error connecting to grpc service")
		return
	}
	if resp != nil && resp.Error != nil {
		response.InternalServerError(c, resp.Error.Message)
		return
	}

	response.Ok(c, "Updated", gin.H{"is_success": true})
}
