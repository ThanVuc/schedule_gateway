package utils

import (
	"schedule_gateway/proto/common"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ToPageQuery(c *gin.Context) *common.PageQuery {
	page := c.Query("page")
	pageSize := c.Query("page_size")
	var sortBy *string

	pageIgnore := c.Query("page_ignore") == "true"

	sortByValue := c.Query("sort_by")
	if sortByValue != "" {
		sortBy = &sortByValue
	} else {
		sortBy = nil
	}

	if page == "" {
		page = "1"
	}

	if pageSize == "" {
		pageSize = "10"
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 1
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		pageSizeInt = 10
	}

	return &common.PageQuery{
		Page:       int32(pageInt),
		PageSize:   int32(pageSizeInt),
		SortBy:     sortBy,
		PageIgnore: &pageIgnore,
	}
}
