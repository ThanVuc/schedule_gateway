package team_controller

import (
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"
	dtos "schedule_gateway/internal/dtos/team_service"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/team_service"
	"strings"

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
		"item": sc.buildSprintResponse(resp.GetSprint()),
	})
}

func (sc *SprintController) GetSprint(ctx *gin.Context) {
	sprintID := ctx.Param("sprint_id")
	if sprintID == "" {
		response.BadRequest(ctx, "Sprint ID is required")
		return
	}

	resp, err := sc.client.GetSprint(ctx, &common.IDRequest{Id: sprintID})
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
		"item": sc.buildSprintResponse(resp.GetSprint()),
	})
}

func (sc *SprintController) ListSprints(ctx *gin.Context) {
	groupID := ctx.Param("group_id")
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
		"items": sprints,
		"total": resp.GetTotal(),
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
		"item": sc.buildSprintResponse(resp.GetSprint()),
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
		"item": gin.H{
			"id":     resp.GetId(),
			"status": resp.GetStatus(),
		},
	})
}

func (sc *SprintController) DeleteSprint(ctx *gin.Context) {
	sprintID := ctx.Param("sprint_id")
	if sprintID == "" {
		response.BadRequest(ctx, "Sprint ID is required")
		return
	}

	resp, err := sc.client.DeleteSprint(ctx, &common.IDRequest{Id: sprintID})
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

	response.NoContent(ctx, "Delete sprint successful", gin.H{
		"item": gin.H{"is_success": resp.GetSuccess()},
	})
}

func (sc *SprintController) buildCreateSprintRequest(ctx *gin.Context) *team_service.CreateSprintRequest {
	groupID := ctx.Param("group_id")
	if groupID == "" {
		response.BadRequest(ctx, "group_id is required")
		return nil
	}

	var dto dtos.CreateSprintDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	name := strings.TrimSpace(dto.Name)
	if name == "" {
		response.BadRequest(ctx, "name is required")
		return nil
	}

	startDate, err := utils.FromStringToDate(dto.StartDate)
	if err != nil {
		response.BadRequest(ctx, "start_date must be in format YYYY-MM-DD")
		return nil
	}

	endDate, err := utils.FromStringToDate(dto.EndDate)
	if err != nil {
		response.BadRequest(ctx, "end_date must be in format YYYY-MM-DD")
		return nil
	}

	start := utils.DateToTime(startDate)
	end := utils.DateToTime(endDate)
	if start.After(end) {
		response.BadRequest(ctx, "start_date must be before or equal to end_date")
		return nil
	}

	return &team_service.CreateSprintRequest{
		GroupId:   groupID,
		Name:      name,
		Goal:      dto.Goal,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func (sc *SprintController) buildUpdateSprintRequest(ctx *gin.Context) *team_service.UpdateSprintRequest {
	sprintID := ctx.Param("sprint_id")
	if sprintID == "" {
		response.BadRequest(ctx, "Sprint ID is required")
		return nil
	}

	var dto dtos.UpdateSprintDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	if dto.Name == nil && dto.Goal == nil && dto.StartDate == nil && dto.EndDate == nil {
		response.BadRequest(ctx, "At least one field is required to update")
		return nil
	}

	var name *string
	if dto.Name != nil {
		trimmed := strings.TrimSpace(*dto.Name)
		if trimmed == "" {
			response.BadRequest(ctx, "name is invalid")
			return nil
		}
		name = &trimmed
	}

	var startDate *team_service.Date
	if dto.StartDate != nil {
		parsedStartDate, err := utils.FromStringToDate(*dto.StartDate)
		if err != nil {
			response.BadRequest(ctx, "start_date must be in format YYYY-MM-DD")
			return nil
		}
		startDate = parsedStartDate
	}

	var endDate *team_service.Date
	if dto.EndDate != nil {
		parsedEndDate, err := utils.FromStringToDate(*dto.EndDate)
		if err != nil {
			response.BadRequest(ctx, "end_date must be in format YYYY-MM-DD")
			return nil
		}
		endDate = parsedEndDate
	}

	if startDate != nil && endDate != nil && utils.DateToTime(startDate).After(utils.DateToTime(endDate)) {
		response.BadRequest(ctx, "start_date must be before or equal to end_date")
		return nil
	}

	return &team_service.UpdateSprintRequest{
		Id:        sprintID,
		Name:      name,
		Goal:      dto.Goal,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

func (sc *SprintController) buildUpdateSprintStatusRequest(ctx *gin.Context) *team_service.UpdateSprintStatusRequest {
	sprintID := ctx.Param("sprint_id")
	if sprintID == "" {
		response.BadRequest(ctx, "Sprint ID is required")
		return nil
	}

	var req team_service.UpdateSprintStatusRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	req.Id = sprintID

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
		"start_date":       utils.FromDateToString(sprint.GetStartDate()),
		"end_date":         utils.FromDateToString(sprint.GetEndDate()),
		"total_work":       sprint.GetTotalWork(),
		"completed_work":   sprint.GetCompletedWork(),
		"progress_percent": sprint.GetProgressPercent(),
		"created_at":       utils.TimestampToISO8601(sprint.GetCreatedAt()),
		"updated_at":       utils.TimestampToISO8601(sprint.GetUpdatedAt()),
	}
}

func isValidSprintStatus(status team_service.SprintStatus) bool {
	_, ok := team_service.SprintStatus_name[int32(status)]
	return ok
}
