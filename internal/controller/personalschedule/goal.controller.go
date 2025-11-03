package personalschedule_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/personalschedule"
	dtos "schedule_gateway/internal/dtos/persional_schedule"
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
