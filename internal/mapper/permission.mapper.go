package mapper

import (
	"schedule_gateway/internal/dtos"
	"schedule_gateway/proto/auth"
)

func MapPermissionsToDTO(permission *auth.GetPermissionsResponse) *dtos.PermissionsResponse {
	if permission == nil {
		return nil
	}

	permissions := make([]dtos.Permission, 0, len(permission.Permissions))
	for _, p := range permission.Permissions {
		permissions = append(permissions, dtos.Permission{
			PermID:      p.PermId,
			PermName:    p.PermName,
			Description: p.Description,
			IsRoot:      p.IsRoot,
		})
	}

	permissionsResponse := dtos.PermissionsResponse{
		Items:             permissions,
		TotalPermissions:  permission.TotalPermissions,
		Root:              permission.Root,
		NonRoot:           permission.NonRoot,
		RootPercentage:    float32(permission.RootPercentage),
		NonRootPercentage: float32(permission.NonRootPercentage),
		TotalItems:        permission.PageInfo.TotalItems,
		TotalPages:        permission.PageInfo.TotalPages,
		PageSize:          permission.PageInfo.PageSize,
		Page:              permission.PageInfo.Page,
		HasPrev:           permission.PageInfo.HasPrev,
		HasNext:           permission.PageInfo.HasNext,
	}

	return &permissionsResponse
}

func MapPermissionToDTO(permission *auth.GetPermissionResponse) *dtos.PermissionResponse {
	if permission == nil {
		return nil
	}

	resource := dtos.Resource{
		Id:   permission.Permission.Resource.Id,
		Name: permission.Permission.Resource.Name,
	}

	actions := make([]dtos.Action, 0, len(permission.Permission.Actions))
	for _, action := range permission.Permission.Actions {
		actions = append(actions, dtos.Action{
			Id:   action.Id,
			Name: action.Name,
		})
	}

	detail := dtos.PermissionDetail{
		PermID:      permission.Permission.PermId,
		PermName:    permission.Permission.PermName,
		Description: permission.Permission.Description,
		IsRoot:      permission.Permission.IsRoot,
		Resource:    resource,
		Action:      actions,
		CreatedAt:   *permission.Permission.CreatedAt,
		UpdatedAt:   *permission.Permission.UpdatedAt,
	}

	return &dtos.PermissionResponse{Data: detail}
}
