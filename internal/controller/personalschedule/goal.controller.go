package personalschedule_controller

import (
	"fmt"
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

	response.Ok(ctx, "Get Goals Successful", dtos.Goals{
		Items:      goalsResp.Goals,
		TotalGoals: goalsResp.TotalGoals,
		TotalPages: goalsResp.PageInfo.TotalPages,
		PageSize:   goalsResp.PageInfo.PageSize,
		Page:       goalsResp.PageInfo.Page,
		HasPrev:    goalsResp.PageInfo.HasPrev,
		HasNext:    goalsResp.PageInfo.HasNext,
	})
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
		Search:    searchString,
		StatusId:  statusID,
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
	if id == "" {
		req.Id = nil
	} else {
		req.Id = &id
	}

	if err := c.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(c, "Invalid request body: "+err.Error())
		fmt.Println("Error binding JSON:", err)
		return nil
	}

	req.UserId = userID
	req.Id = &id
	req.Name = dto.Name
	req.DetailedDescription = dto.DetailedDescription
	req.StatusId = dto.StatusID
	req.DifficultyId = dto.DifficultyID
	req.PriorityId = dto.PriorityID
	req.ShortDescriptions = dto.ShortDescriptions
	req.StartDate = dto.StartDate
	req.EndDate = dto.EndDate
	req.Tasks = make([]*personal_schedule.GoalTaskPayload, len(dto.Tasks))

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
