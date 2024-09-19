package form

type GroupServerConfigForm struct {
	Name     string `form:"name"`
	Path     string `form:"file_path" binding:"required"`
	GroupId  int    `form:"group_id" binding:"required"`
	ServerId int    `form:"server_id" `
	Id       int    `form:"id"`
}
