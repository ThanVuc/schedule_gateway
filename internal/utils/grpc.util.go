package utils

import (
	"schedule_gateway/proto/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToPageQuery(c *gin.Context) *common.PageQuery {
	page := c.Query("page")
	pageSize := c.Query("page_size")

	if page == "" {
		page = "1"
	}

	if pageSize == "" {
		pageSize = "10"
	}

	// Convert page and pageSize to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1 // Default to 1 if conversion fails
	}
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10 // Default to 10 if conversion fails
	}

	return &common.PageQuery{
		Page:     int32(pageInt),
		PageSize: int32(pageSizeInt),
	}
}
