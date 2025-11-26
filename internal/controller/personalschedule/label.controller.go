package personalschedule_controller

import (
	"schedule_gateway/global"
	client "schedule_gateway/internal/client/personalschedule"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/common"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type LabelController struct {
	logger log.Logger
	client client.LabelClient
}

func NewLabelController() *LabelController {
	return &LabelController{
		logger: global.Logger,
		client: client.NewLabelClient(),
	}
}

func (lc *LabelController) GetLabelPerTypes(ctx *gin.Context) {
	req := &common.EmptyRequest{}

	resp, err := lc.client.GetLabelPerTypes(ctx, req)
	if err != nil {
		lc.logger.Error("Connection error: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Error connecting to personalschedule service")
		return
	}

	if resp != nil && resp.Error != nil {
		if resp.Error != nil && resp.Error.ErrorCode != nil {
			response.InternalServerError(ctx, utils.Int32PtrToString(resp.Error.ErrorCode))
			return
		}
		response.InternalServerError(ctx, resp.Error.Message)
		return
	}

	response.Ok(ctx, "Ok", resp)
}

func (lc *LabelController) GetLabelsByTypeIDs(ctx *gin.Context) {
	var req common.IDRequest
	req.Id = ctx.Param("type_id")
	if req.Id == "" {
		response.BadRequest(ctx, "type_id is required")
		return
	}

	resp, err := lc.client.GetLabelsByTypeIDs(ctx, &req)
	if err != nil {
		lc.logger.Error("Connection error: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Error connecting to grpc service")
		return
	}

	if resp != nil && resp.Error != nil {
		if resp.Error != nil && resp.Error.ErrorCode != nil {
			println(utils.Int32PtrToString(resp.Error.ErrorCode))
			response.InternalServerError(ctx, utils.Int32PtrToString(resp.Error.ErrorCode))
			return
		}
		response.InternalServerError(ctx, resp.Error.Message)
		return
	}

	response.Ok(ctx, "Ok", resp)
}

func (lc *LabelController) GetDefaultLabel(ctx *gin.Context) {
	req := &common.EmptyRequest{}

	resp, err := lc.client.GetDefaultLabel(ctx, req)
	if err != nil {
		lc.logger.Error("Connection error: ", "", zap.Error(err))
		response.InternalServerError(ctx, "Error connecting to grpc service")
		return
	}
	if resp != nil && resp.Error != nil {
		if resp.Error != nil && resp.Error.ErrorCode != nil {
			response.InternalServerError(ctx, utils.Int32PtrToString(resp.Error.ErrorCode))
			return
		}
	}
	response.Ok(ctx, "Ok", resp)
}
