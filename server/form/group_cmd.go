package form

type GroupServerCmdForm struct {
	GroupConfigAlias string `form:"group_config_alias" binding:"required"`
	GroupPathAlias   string `form:"group_path_alias" binding:"required"`
	StartCommand     string `form:"start_command" binding:"required"`
	StartUser        string `form:"start_user" binding:"required"`
	GroupId          int    `form:"group_id" binding:"required"`
	Name             string `form:"name" binding:"required"`
	ServerId         int    `form:"server_id"`
	Id               int    `form:"id"`
}
