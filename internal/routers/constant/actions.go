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
	REVOKE_TOKEN_ACTION  = "revoke_token"
	SYNC_DATABASE_ACTION = "sync_database"
)

// admin users actions
const (
	ASSIGN_ROLE_ACTION    = "assign_role"
	READ_ALL_USERS_ACTION = "read_all"
	READ_ONE_USER_ACTION  = "read_one"
	LOCK_USER_ACTION      = "lock_user"
)

// auth resource actions
const (
	ME_ACTION = "me"
)

// label resource actions
const (
	READ_ALL_LABEL_PER_TYPES_ACTION = "read_all_label_per_types"
	READ_LABELS_BY_TYPE_ACTION      = "read_labels_by_type"
)

// goal resource actions
const (
	READ_GOALS_ACTION = "read_all_goals"
	CREATE_GOAL_ACTION = "create_goal"
	UPDATE_GOAL_ACTION = "update_goal"
)
