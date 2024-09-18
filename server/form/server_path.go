package form

type GroupServerPathForm struct {
	Name      string `form:"name"`
	Path      string `form:"path" binding:"required"`
	GroupId   int    `form:"group_id" binding:"required"`
	ServerIds []int  `form:"server_ids" binding:"required"`
	PathId    int    `form:"path_id"`
}
