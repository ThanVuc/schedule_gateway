package perm_constant

// common actions
const (
	READ_ALL_ACTION           = "read_all"
	READ_ONE_ACTION           = "read_one"
	UPDATE_ACTION             = "update"
	CREATE_ACTION             = "create"
	DELETE_ACTION             = "delete"
	ENABLE_AND_DISABLE_ACTION = "enable_and_disable"
)

// action and resource constants
const (
	READ_RESOURCES_ACTION = "read_resources"
	READ_ACTIONS_ACTION   = "read_actions"
)

// authentication actions
const (
	LOGOUT_ACTION        = "logout"
	REFRESH_TOKEN_ACTION = "refresh_token"
	REVOKE_TOKEN_ACTION  = "revoke_token"
)

// admin users actions
const (
	ASSIGN_ROLE_ACTION = "assign_role"
)

// auth resource actions
const (
	ME_ACTION = "me"
)
