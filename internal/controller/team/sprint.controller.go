package team_controller

import (
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/team_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type SprintController struct {
	logger log.Logger
	client team_client.SprintClient
}

func NewSprintController() *SprintController {
	return &SprintController{
		logger: global.Logger,
		client: team_client.NewSprintClient(),
	}
}

func (sc *SprintController) CreateSprint(ctx *gin.Context) {
	req := sc.buildCreateSprintRequest(ctx)
	if req == nil {
		return
	}

	resp, err := sc.client.CreateSprint(ctx, req)
	if err != nil {
		sc.logger.Error("Failed to create sprint: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to create sprint")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Create sprint successful", gin.H{
		"sprint": sc.buildSprintResponse(resp.GetSprint()),
	})
}

func (sc *SprintController) GetSprint(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.BadRequest(ctx, "Sprint ID is required")
		return
	}

	resp, err := sc.client.GetSprint(ctx, &common.IDRequest{Id: id})
	if err != nil {
		sc.logger.Error("Failed to get sprint: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to get sprint")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Get sprint successful", gin.H{
		"sprint": sc.buildSprintResponse(resp.GetSprint()),
	})
}

func (sc *SprintController) ListSprints(ctx *gin.Context) {
	groupID := ctx.Query("group_id")
	if groupID == "" {
		groupID = ctx.Param("group_id")
	}
	if groupID == "" {
		response.BadRequest(ctx, "group_id is required")
		return
	}

	resp, err := sc.client.ListSprints(ctx, &team_service.ListSprintsRequest{GroupId: groupID})
	if err != nil {
		sc.logger.Error("Failed to list sprints: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to list sprints")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	sprints := make([]gin.H, 0, len(resp.GetSprints()))
	for _, sprint := range resp.GetSprints() {
		sprints = append(sprints, sc.buildSprintResponse(sprint))
	}

	response.Ok(ctx, "List sprints successful", gin.H{
		"sprints": sprints,
		"total":   resp.GetTotal(),
	})
}

func (sc *SprintController) UpdateSprint(ctx *gin.Context) {
	req := sc.buildUpdateSprintRequest(ctx)
	if req == nil {
		return
	}

	resp, err := sc.client.UpdateSprint(ctx, req)
	if err != nil {
		sc.logger.Error("Failed to update sprint: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to update sprint")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Update sprint successful", gin.H{
		"sprint": sc.buildSprintResponse(resp.GetSprint()),
	})
}

func (sc *SprintController) UpdateSprintStatus(ctx *gin.Context) {
	req := sc.buildUpdateSprintStatusRequest(ctx)
	if req == nil {
		return
	}

	resp, err := sc.client.UpdateSprintStatus(ctx, req)
	if err != nil {
		sc.logger.Error("Failed to update sprint status: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to update sprint status")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Update sprint status successful", gin.H{
		"id":     resp.GetId(),
		"status": resp.GetStatus(),
	})
}

func (sc *SprintController) DeleteSprint(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.BadRequest(ctx, "Sprint ID is required")
		return
	}

	resp, err := sc.client.DeleteSprint(ctx, &common.IDRequest{Id: id})
	if err != nil {
		sc.logger.Error("Failed to delete sprint: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to delete sprint")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if resp.GetError() != nil {
		response.UnprocessableEntity(ctx, resp.GetError().GetCode(), resp.GetError().GetMessage(), utils.SafeString(resp.GetError().Details))
		return
	}

	response.Ok(ctx, "Delete sprint successful", gin.H{
		"is_success": resp.GetSuccess(),
	})
}

func (sc *SprintController) buildCreateSprintRequest(ctx *gin.Context) *team_service.CreateSprintRequest {
	var req team_service.CreateSprintRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	if req.GetGroupId() == "" {
		response.BadRequest(ctx, "group_id is required")
		return nil
	}

	if req.GetName() == "" {
		response.BadRequest(ctx, "name is required")
		return nil
	}

	if !utils.IsValidDate(req.GetStartDate()) || !utils.IsValidDate(req.GetEndDate()) {
		response.BadRequest(ctx, "start_date and end_date must be valid dates")
		return nil
	}

	start := utils.DateToTime(req.GetStartDate())
	end := utils.DateToTime(req.GetEndDate())
	if start.After(end) {
		response.BadRequest(ctx, "start_date must be before or equal to end_date")
		return nil
	}

	return &req
}

func (sc *SprintController) buildUpdateSprintRequest(ctx *gin.Context) *team_service.UpdateSprintRequest {
	id := ctx.Param("id")
	if id == "" {
		response.BadRequest(ctx, "Sprint ID is required")
		return nil
	}

	var req team_service.UpdateSprintRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	req.Id = id

	if req.Name == nil && req.Goal == nil && req.StartDate == nil && req.EndDate == nil {
		response.BadRequest(ctx, "At least one field is required to update")
		return nil
	}

	if req.StartDate != nil && !utils.IsValidDate(req.StartDate) {
		response.BadRequest(ctx, "start_date is invalid")
		return nil
	}

	if req.EndDate != nil && !utils.IsValidDate(req.EndDate) {
		response.BadRequest(ctx, "end_date is invalid")
		return nil
	}

	if req.StartDate != nil && req.EndDate != nil {
		if utils.DateToTime(req.StartDate).After(utils.DateToTime(req.EndDate)) {
			response.BadRequest(ctx, "start_date must be before or equal to end_date")
			return nil
		}
	}

	return &req
}

func (sc *SprintController) buildUpdateSprintStatusRequest(ctx *gin.Context) *team_service.UpdateSprintStatusRequest {
	id := ctx.Param("id")
	if id == "" {
		response.BadRequest(ctx, "Sprint ID is required")
		return nil
	}

	var req team_service.UpdateSprintStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	req.Id = id

	if !isValidSprintStatus(req.GetStatus()) || req.GetStatus() == team_service.SprintStatus_SPRINT_STATUS_UNSPECIFIED {
		response.BadRequest(ctx, "status is invalid")
		return nil
	}

	return &req
}

func (sc *SprintController) buildSprintResponse(sprint *team_service.SprintMessage) gin.H {
	if sprint == nil {
		return gin.H{}
	}

	return gin.H{
		"id":               sprint.GetId(),
		"group_id":         sprint.GetGroupId(),
		"name":             sprint.GetName(),
		"goal":             sprint.GetGoal(),
		"status":           sprint.GetStatus(),
		"start_date":       sprint.GetStartDate(),
		"end_date":         sprint.GetEndDate(),
		"total_work":       sprint.GetTotalWork(),
		"completed_work":   sprint.GetCompletedWork(),
		"progress_percent": sprint.GetProgressPercent(),
		"created_at":       sprint.GetCreatedAt(),
		"updated_at":       sprint.GetUpdatedAt(),
	}
}

func isValidSprintStatus(status team_service.SprintStatus) bool {
	_, ok := team_service.SprintStatus_name[int32(status)]
	return ok
}
