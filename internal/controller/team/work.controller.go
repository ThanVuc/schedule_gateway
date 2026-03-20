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
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type WorkController struct {
	logger log.Logger
	client team_client.WorkClient
}

func NewWorkController() *WorkController {
	return &WorkController{
		logger: global.Logger,
		client: team_client.NewWorkClient(),
	}
}

func (wc *WorkController) CreateWork(ctx *gin.Context) {
	req := wc.buildCreateWorkRequest(ctx)
	if req == nil {
		return
	}

	resp, err := wc.client.CreateWork(ctx, req)
	if err != nil {
		wc.logger.Error("Failed to create work: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to create work")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.Created(ctx, "Create work successful", gin.H{
		"item": wc.buildWorkResponse(resp.GetWork()),
	})
}

func (wc *WorkController) GetWork(ctx *gin.Context) {
	workID := ctx.Param("work_id")
	if workID == "" {
		response.BadRequest(ctx, "work_id is required")
		return
	}

	resp, err := wc.client.GetWork(ctx, &common.IDRequest{Id: workID})
	if err != nil {
		wc.logger.Error("Failed to get work: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to get work")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.Ok(ctx, "Get work successful", gin.H{
		"item": wc.buildWorkResponse(resp.GetWork()),
	})
}

func (wc *WorkController) ListWorks(ctx *gin.Context) {
	sprintID := ctx.Query("sprint_id")

	resp, err := wc.client.ListWorks(ctx, &team_service.ListWorksRequest{SprintId: utils.Ptr(sprintID)})
	if err != nil {
		wc.logger.Error("Failed to ilst works: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to list works")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	works := make([]gin.H, 0, len(resp.GetWorks()))
	for _, work := range resp.GetWorks() {
		works = append(works, wc.buildWorkResponse(work))
	}

	response.Ok(ctx, "List works successful", gin.H{
		"items": works,
		"total": len(works),
	})
}

func (wc *WorkController) UpdateWork(ctx *gin.Context) {
	req := wc.buildUpdateWorkRequest(ctx)
	if req == nil {
		return
	}

	resp, err := wc.client.UpdateWork(ctx, req)
	if err != nil {
		wc.logger.Error("Failed to update work: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to update work")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.Ok(ctx, "Update work successful", gin.H{
		"item": wc.buildWorkResponse(resp.GetWork()),
	})
}

func (wc *WorkController) DeleteWork(ctx *gin.Context) {
	workID := ctx.Param("work_id")
	if workID == "" {
		response.BadRequest(ctx, "work_id is required")
		return
	}

	resp, err := wc.client.DeleteWork(ctx, &common.IDRequest{Id: workID})
	if err != nil {
		wc.logger.Error("Failed to delete work: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to delete work")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.NoContent(ctx, "Delete work successful", gin.H{
		"item": gin.H{"is_success": resp.GetSuccess()},
	})
}

func (wc *WorkController) CreateChecklistItem(ctx *gin.Context) {
	req := wc.buildCreateChecklistItemRequest(ctx)
	if req == nil {
		return
	}

	resp, err := wc.client.CreateChecklistItem(ctx, req)
	if err != nil {
		wc.logger.Error("Failed to create checklist item: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to create checklist item")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.Created(ctx, "Create checklist item successful", gin.H{
		"item": wc.buildChecklistItemResponse(resp.GetItem()),
	})
}

func (wc *WorkController) UpdateChecklistItem(ctx *gin.Context) {
	req := wc.buildUpdateChecklistItemRequest(ctx)
	if req == nil {
		return
	}

	resp, err := wc.client.UpdateChecklistItem(ctx, req)
	if err != nil {
		wc.logger.Error("Failed to update checklist item: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to update checklist item")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.Ok(ctx, "Update checklist item successful", gin.H{
		"item": wc.buildChecklistItemResponse(resp.GetChecklist()),
	})
}

func (wc *WorkController) DeleteChecklistItem(ctx *gin.Context) {
	checklistID := ctx.Param("checklist_id")
	if checklistID == "" {
		response.BadRequest(ctx, "checklist_id is required")
		return
	}

	resp, err := wc.client.DeleteChecklistItem(ctx, &common.IDRequest{Id: checklistID})
	if err != nil {
		wc.logger.Error("Failed to delete checklist item: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to delete checklist item")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.NoContent(ctx, "Delete checklist item successful", gin.H{
		"item": wc.buildChecklistItemResponse(resp.GetChecklist()),
	})
}

func (wc *WorkController) CreateComment(ctx *gin.Context) {
	req := wc.buildCreateCommentRequest(ctx)
	if req == nil {
		return
	}

	resp, err := wc.client.CreateComment(ctx, req)
	if err != nil {
		wc.logger.Error("Failed to create comment: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to create comment")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.Created(ctx, "Create comment successful", gin.H{
		"items": wc.buildCommentListResponse(resp.GetComment()),
	})
}

func (wc *WorkController) UpdateComment(ctx *gin.Context) {
	req := wc.buildUpdateCommentRequest(ctx)
	if req == nil {
		return
	}

	resp, err := wc.client.UpdateComment(ctx, req)
	if err != nil {
		wc.logger.Error("Failed to update comment: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to update comment")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.Ok(ctx, "Update comment successful", gin.H{
		"items": wc.buildCommentListResponse(resp.GetComment()),
	})
}

func (wc *WorkController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("comment_id")
	if commentID == "" {
		response.BadRequest(ctx, "comment_id is required")
		return
	}

	resp, err := wc.client.DeleteComment(ctx, &common.IDRequest{Id: commentID})
	if err != nil {
		wc.logger.Error("Failed to delete comment: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Failed to delete comment")
		return
	}

	if resp == nil {
		response.InternalServerError(ctx, "Empty response from service")
		return
	}

	if wc.handleServiceError(ctx, resp.GetError()) {
		return
	}

	response.NoContent(ctx, "Delete comment successful", gin.H{
		"items": wc.buildCommentListResponse(resp.GetComment()),
	})
}

func (wc *WorkController) buildCreateWorkRequest(ctx *gin.Context) *team_service.CreateWorkRequest {
	var dto dtos.CreateWorkDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	name := strings.TrimSpace(dto.Name)
	if name == "" {
		response.BadRequest(ctx, "name is required")
		return nil
	}

	return &team_service.CreateWorkRequest{
		Name:        name,
		Description: wc.normalizeOptionalString(dto.Description),
		SprintId:    dto.SprintID,
	}
}

func (wc *WorkController) buildUpdateWorkRequest(ctx *gin.Context) *team_service.UpdateWorkRequest {
	workID := ctx.Param("work_id")
	if workID == "" {
		response.BadRequest(ctx, "work_id is required")
		return nil
	}

	var dto dtos.UpdateWorkDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	if dto.Version == nil {
		response.BadRequest(ctx, "version is required")
		return nil
	}

	if *dto.Version < 0 {
		response.BadRequest(ctx, "version must be greater or equal than 0")
		return nil
	}

	if dto.Name == nil && dto.Description == nil && dto.AssigneeID == nil && dto.Status == nil && dto.StoryPoint == nil && dto.DueDate == nil && dto.Priority == nil {
		response.BadRequest(ctx, "at least one field must be provided")
		return nil
	}

	var name *string
	if dto.Name != nil {
		trimmedName := strings.TrimSpace(*dto.Name)
		nameLength := utf8.RuneCountInString(trimmedName)
		if nameLength < 1 || nameLength > 500 {
			response.BadRequest(ctx, "name must be between 1 and 500 characters")
			return nil
		}

		name = &trimmedName
	}

	var description *string
	if dto.Description != nil {
		if utf8.RuneCountInString(*dto.Description) > 5000 {
			response.BadRequest(ctx, "description must be at most 5000 characters")
			return nil
		}

		desc := *dto.Description
		description = &desc
	}

	var assigneeID *string
	if dto.AssigneeID != nil {
		trimmedAssigneeID := strings.TrimSpace(*dto.AssigneeID)
		if trimmedAssigneeID == "" {
			response.BadRequest(ctx, "assignee_id is invalid")
			return nil
		}

		if !isValidUUIDString(trimmedAssigneeID) {
			response.BadRequest(ctx, "assignee_id must be a valid UUID")
			return nil
		}

		assigneeID = &trimmedAssigneeID
	}

	var status *team_service.WorkStatus
	if dto.Status != nil {
		parsedStatus := team_service.WorkStatus(*dto.Status)
		if !isValidWorkStatus(parsedStatus) || parsedStatus == team_service.WorkStatus_WORK_STATUS_UNSPECIFIED {
			response.BadRequest(ctx, "status is invalid")
			return nil
		}

		status = &parsedStatus
	}

	var priority *team_service.WorkPriority
	if dto.Priority != nil {
		parsedPriority := team_service.WorkPriority(*dto.Priority)
		if !isValidWorkPriority(parsedPriority) || parsedPriority == team_service.WorkPriority_WORK_PRIORITY_UNSPECIFIED {
			response.BadRequest(ctx, "priority is invalid")
			return nil
		}

		priority = &parsedPriority
	}

	storyPoint := wc.normalizeOptionalString(dto.StoryPoint)

	dueDate, ok := wc.parseOptionalDate(ctx, dto.DueDate, "due_date")
	if !ok {
		return nil
	}

	return &team_service.UpdateWorkRequest{
		Id:          workID,
		Name:        name,
		Description: description,
		AssigneeId:  assigneeID,
		StoryPoint:  storyPoint,
		DueDate:     dueDate,
		Priority:    priority,
		Status:      status,
		Version:     *dto.Version,
	}
}

func (wc *WorkController) buildCreateChecklistItemRequest(ctx *gin.Context) *team_service.CreateChecklistItemRequest {
	workID := ctx.Param("work_id")
	if workID == "" {
		response.BadRequest(ctx, "work_id is required")
		return nil
	}

	var dto dtos.CreateChecklistItemDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	name := strings.TrimSpace(dto.Name)
	if name == "" {
		response.BadRequest(ctx, "name is required")
		return nil
	}

	return &team_service.CreateChecklistItemRequest{
		WorkId: workID,
		Name:   name,
	}
}

func (wc *WorkController) buildUpdateChecklistItemRequest(ctx *gin.Context) *team_service.UpdateChecklistItemRequest {
	checklistID := ctx.Param("checklist_id")
	if checklistID == "" {
		response.BadRequest(ctx, "checklist_id is required")
		return nil
	}

	var dto dtos.UpdateChecklistItemDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	if dto.Name == nil && dto.IsCompleted == nil {
		response.BadRequest(ctx, "At least one field is required to update")
		return nil
	}

	var name *string
	if dto.Name != nil {
		trimmedName := strings.TrimSpace(*dto.Name)
		if trimmedName == "" {
			response.BadRequest(ctx, "name is invalid")
			return nil
		}
		name = &trimmedName
	}

	return &team_service.UpdateChecklistItemRequest{
		Id:          checklistID,
		Name:        name,
		IsCompleted: dto.IsCompleted,
	}
}

func (wc *WorkController) buildCreateCommentRequest(ctx *gin.Context) *team_service.CreateCommentRequest {
	workID := ctx.Param("work_id")
	if workID == "" {
		response.BadRequest(ctx, "work_id is required")
		return nil
	}

	var dto dtos.CreateCommentDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	content := strings.TrimSpace(dto.Content)
	if content == "" {
		response.BadRequest(ctx, "content is required")
		return nil
	}

	return &team_service.CreateCommentRequest{
		WorkId:  workID,
		Content: content,
	}
}

func (wc *WorkController) buildUpdateCommentRequest(ctx *gin.Context) *team_service.UpdateCommentRequest {
	commentID := ctx.Param("comment_id")
	if commentID == "" {
		response.BadRequest(ctx, "comment_id is required")
		return nil
	}

	var dto dtos.UpdateCommentDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		response.BadRequest(ctx, "Invalid request body: "+err.Error())
		return nil
	}

	content := strings.TrimSpace(dto.Content)
	if content == "" {
		response.BadRequest(ctx, "content is required")
		return nil
	}

	return &team_service.UpdateCommentRequest{
		Id:      commentID,
		Content: content,
	}
}

func (wc *WorkController) normalizeOptionalString(value *string) *string {
	if value == nil {
		return nil
	}

	trimmed := strings.TrimSpace(*value)
	return &trimmed
}

func (wc *WorkController) parseOptionalDate(ctx *gin.Context, value *string, fieldName string) (*team_service.Date, bool) {
	if value == nil {
		return nil, true
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed == "" {
		response.BadRequest(ctx, fieldName+" is invalid")
		return nil, false
	}

	date, err := utils.FromStringToDate(trimmed)
	if err != nil {
		response.BadRequest(ctx, fieldName+" must be in format YYYY-MM-DD")
		return nil, false
	}

	return date, true
}

func (wc *WorkController) handleServiceError(ctx *gin.Context, serviceErr *team_service.Error) bool {
	if serviceErr == nil {
		return false
	}

	response.UnprocessableEntity(ctx, serviceErr.GetCode(), serviceErr.GetMessage(), utils.SafeString(serviceErr.Details))
	return true
}

func (wc *WorkController) buildWorkResponse(work *team_service.WorkMessage) gin.H {
	if work == nil {
		return gin.H{}
	}

	checkList := gin.H{
		"total":     int32(0),
		"completed": int32(0),
		"items":     []gin.H{},
	}

	if work.GetCheckList() != nil {
		items := make([]gin.H, 0, len(work.GetCheckList().GetItems()))
		for _, item := range work.GetCheckList().GetItems() {
			items = append(items, wc.buildChecklistItemResponse(item))
		}

		checkList = gin.H{
			"total":     work.GetCheckList().GetTotal(),
			"completed": work.GetCheckList().GetCompleted(),
			"items":     items,
		}
	}

	return gin.H{
		"id":          work.GetId(),
		"name":        work.GetName(),
		"description": work.GetDescription(),
		"status":      work.GetStatus(),
		"sprint":      wc.buildSimpleSprintResponse(work.GetSprint()),
		"assignee":    wc.buildSimpleUserResponse(work.GetAssignee()),
		"story_point": work.GetStoryPoint(),
		"due_date":    utils.FromDateToString(work.GetDueDate()),
		"check_list":  checkList,
		"comments":    wc.buildCommentListResponse(work.GetComments()),
		"created_at":  utils.TimestampToISO8601(work.GetCreatedAt()),
		"updated_at":  utils.TimestampToISO8601(work.GetUpdatedAt()),
		"version":     work.GetVersion(),
	}
}

func (wc *WorkController) buildChecklistItemResponse(item *team_service.ChecklistItemMessage) gin.H {
	if item == nil {
		return gin.H{}
	}

	return gin.H{
		"id":           item.GetId(),
		"name":         item.GetName(),
		"is_completed": item.GetIsCompleted(),
	}
}

func (wc *WorkController) buildCommentListResponse(commentList *team_service.CommentListMessage) gin.H {
	comments := make([]gin.H, 0)
	if commentList == nil {
		return gin.H{
			"total": int32(0),
			"items": comments,
		}
	}

	for _, comment := range commentList.GetComments() {
		comments = append(comments, wc.buildCommentResponse(comment))
	}

	return gin.H{
		"total": commentList.GetTotal(),
		"items": comments,
	}
}

func (wc *WorkController) buildCommentResponse(comment *team_service.CommentMessage) gin.H {
	if comment == nil {
		return gin.H{}
	}

	return gin.H{
		"id":         comment.GetId(),
		"content":    comment.GetContent(),
		"creator":    wc.buildSimpleUserResponse(comment.GetCreator()),
		"created_at": utils.TimestampToISO8601(comment.GetCreatedAt()),
	}
}

func (wc *WorkController) buildSimpleUserResponse(user *team_service.SimpleUserMessage) gin.H {
	if user == nil {
		return gin.H{}
	}

	return gin.H{
		"id":     user.GetId(),
		"email":  user.GetEmail(),
		"avatar": user.GetAvatar(),
	}
}

func (wc *WorkController) buildSimpleSprintResponse(sprint *team_service.SimpleSprintMessage) gin.H {
	if sprint == nil {
		return gin.H{}
	}

	return gin.H{
		"id":   sprint.GetId(),
		"name": sprint.GetName(),
	}
}

func isValidWorkStatus(status team_service.WorkStatus) bool {
	_, ok := team_service.WorkStatus_name[int32(status)]
	return ok
}

func isValidWorkPriority(priority team_service.WorkPriority) bool {
	_, ok := team_service.WorkPriority_name[int32(priority)]
	return ok
}

func isValidUUIDString(value string) bool {
	_, err := uuid.Parse(strings.TrimSpace(value))
	return err == nil
}
