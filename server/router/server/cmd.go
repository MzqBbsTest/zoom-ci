package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/module/server"
	"github.com/zoom-ci/zoom-ci/server/render"
	"github.com/zoom-ci/zoom-ci/util/gostring"
)

type CmdForm struct {
	Id      int    `form:"id" default="0"`
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func CmdList(c *gin.Context) {
	var query QueryBind
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	cmd := &server.Cmd{}
	list, err := cmd.ListByUserId(query.Keyword, c.GetInt("user_id"), query.Offset, query.Limit)
	if err != nil {
		render.AppError(c, err.Error())
		return
	}
	total, err := cmd.TotalByUserId(query.Keyword, c.GetInt("user_id"))
	if err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.JSON(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func CmdSave(c *gin.Context) {
	var cmdForm CmdForm
	if err := c.ShouldBind(&cmdForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	server := &server.Cmd{
		ID:      cmdForm.Id,
		Title:   cmdForm.Title,
		Content: cmdForm.Content,
		UserId:  c.GetInt("user_id"),
	}
	if err := server.CreateOrUpdate(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.Success(c)
}

func CmdDelete(c *gin.Context) {
	id := gostring.Str2Int(c.PostForm("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	cmd := &server.Cmd{
		ID: id,
	}
	if err := cmd.Delete(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.JSON(c, nil)
}
