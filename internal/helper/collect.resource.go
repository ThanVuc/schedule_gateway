package helper

import (
	"encoding/json"
	"os"
	"schedule_gateway/global"
	"schedule_gateway/internal/grpc/auth"
	"strconv"

	"go.uber.org/zap"
)

type Resource struct {
	ResourceId string `json:"resource_id"`
	Resource   string `json:"resource"`
}

type Action struct {
	ActionId string `json:"action_id"`
	Action   string `json:"action"`
}

type ResourceItem struct {
	Resource Resource `json:"resource"`
	Actions  []Action `json:"actions"`
}

var resourceList []*auth.ResourceItem = make([]*auth.ResourceItem, 0)

type ResourceRegiseter struct {
	count      int
	resourceId string
}

func NewResourceRegiseter(resourceId string) *ResourceRegiseter {
	return &ResourceRegiseter{
		count:      0,
		resourceId: resourceId,
	}
}

func (rr *ResourceRegiseter) AddResource(resource *auth.Resource, actions []*auth.Action) {
	resourceList = append(resourceList, &auth.ResourceItem{
		Resource: resource,
		Actions:  actions,
	})
}

func GetResources() []*auth.ResourceItem {
	return resourceList
}

func WriteToJsonFile(fileName string) {
	logger := global.Logger
	filePath := "./backup/" + fileName + ".json"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logger.ErrorString("Failed to open file for writing", zap.Error(err))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(resourceList); err != nil {
		logger.ErrorString("Failed to encode resource list to JSON", zap.Error(err))
	} else {
		logger.InfoString("Resource list written to JSON file successfully", zap.String("file_path", filePath))
	}
}

func (rr *ResourceRegiseter) GenerateActionId() string {
	// Generate a unique action ID based on resource ID and action name
	rr.count++
	return rr.resourceId + strconv.Itoa(rr.count)
}
