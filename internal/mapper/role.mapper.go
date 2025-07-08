package mapper

import (
	"schedule_gateway/internal/dtos"
	"schedule_gateway/proto/auth"
)

func MapRolesToDTO(role *auth.GetRolesResponse) *dtos.RolesResponse {
	if role == nil {
		return nil
	}

	roles := make([]dtos.Role, 0, len(role.Roles))
	for _, r := range role.Roles {
		roles = append(roles, dtos.Role{
			RoleID:      r.RoleId,
			Name:        r.Name,
			Description: r.Description,
			IsRoot:      r.IsRoot,
			TotalUsers:  r.TotalUsers,
			IsActive:    r.IsActive,
		})
	}

	rolesResponse := dtos.RolesResponse{
		Items:             roles,
		TotalRoles:        role.TotalRoles,
		Root:              role.Root,
		NonRoot:           role.NonRoot,
		RootPercentage:    float32(role.RootPercentage),
		NonRootPercentage: float32(role.NonRootPercentage),
		TotalItems:        role.PageInfo.TotalItems,
		TotalPages:        role.PageInfo.TotalPages,
		PageSize:          role.PageInfo.PageSize,
		Page:              role.PageInfo.Page,
		HasPrev:           role.PageInfo.HasPrev,
		HasNext:           role.PageInfo.HasNext,
	}

	return &rolesResponse
}

func MapRoleToDTO(role *auth.GetRoleResponse) *dtos.RoleResponse {
	if role == nil {
		return nil
	}

	perms := make([]dtos.PermRole, 0, len(role.Role.Permissions))
	for _, p := range role.Role.Permissions {
		perms = append(perms, dtos.PermRole{
			PermId:      p.PermId,
			PermName:    p.PermName,
			Description: p.Description,
		})
	}

	roleDTO := dtos.RoleDetail{
		RoleID:      role.Role.RoleId,
		Name:        role.Role.Name,
		Description: role.Role.Description,
		IsRoot:      role.Role.IsRoot,
		IsActive:    role.Role.IsActive,
		CreatedAt:   *role.Role.CreatedAt,
		UpdatedAt:   *role.Role.UpdatedAt,
		Permissions: perms,
	}

	return &dtos.RoleResponse{Data: roleDTO}
}
