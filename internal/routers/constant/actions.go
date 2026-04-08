package perm_constant

// common actions
const (
	READ_ALL_ACTION           = "read_all"
	READ_ONE_ACTION           = "read_one"
	UPDATE_ACTION             = "update"
	CREATE_ACTION             = "create"
	DELETE_ACTION             = "delete"
	ENABLE_AND_DISABLE_ACTION = "enable_and_disable"
	EXPORT_ACTION             = "export"
	GENERATE_ACTION           = "generate"
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
	READ_DEFAULT_LABEL_ACTION       = "read_default_label"
)

// goal for dialog actions
const (
	READ_GOALS_FOR_DIALOG_ACTION = "read_goals_for_dialog"
)

// notification resource actions
const (
	SAVE_FCM_TOKEN_ACTION = "save_token"
	MARK_AS_READ_ACTION   = "mark_as_read"
)

// work resource actions
const (
	RECOVER_WORKS_ACTION           = "recover_works"
	ACCEPT_ALL_DRAFTS_WORKS_ACTION = "accept_all_drafts_works"
)

// group resource actions
const (
	READ_LIST_MEMBERS_ACTION  = "read_list_members"
	UPDATE_MEMBER_ROLE_ACTION = "update_member_role"
	REMOVE_MEMBER_ACTION      = "remove_member"
	CREATE_INVITE_ACTION      = "create_invite"
	ACCEPT_INVITE_ACTION      = "accept_invite"
)

// team user resource actions
const (
	UPDATE_NOTIFICATION_CONFIGURATION_ACTION = "update_notification_configuration"
	GET_PRESINGED_URLS_ACTION                = "get_presigned_urls"
)
