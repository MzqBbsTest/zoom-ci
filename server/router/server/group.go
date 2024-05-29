// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/module/server"
	"github.com/zoom-ci/zoom-ci/server/render"
	"github.com/zoom-ci/zoom-ci/util/gostring"
)

type GroupForm struct {
	ID      int    `form:"id"`
	Name    string `form:"name" binding:"required"`
	Servers []int  `form:"servers" `
}

func GroupAdd(c *gin.Context) {
	var groupForm GroupForm
	if err := c.ShouldBind(&groupForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	group := &server.Group{
		Name: groupForm.Name,
	}
	if err := group.Create(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.Success(c)
}

func GroupList(c *gin.Context) {
	var query QueryBind
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	group := &server.Group{}
	list, err := group.List(query.Keyword, query.Offset, query.Limit)
	if err != nil {
		render.AppError(c, err.Error())
		return
	}
	total, err := group.Total(query.Keyword)
	if err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.JSON(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func GroupDelete(c *gin.Context) {
	id := gostring.Str2Int(c.PostForm("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	group := &server.Group{
		ID: id,
	}
	if err := group.Delete(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.JSON(c, nil)
}

func GroupDetail(c *gin.Context) {
	id := gostring.Str2Int(c.Query("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	group := &server.Group{
		ID: id,
	}
	if err := group.Detail(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.JSON(c, group)
}

func GroupUpdate(c *gin.Context) {
	var groupForm GroupForm
	if err := c.ShouldBind(&groupForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	if groupForm.ID == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	group := &server.Group{
		ID:        groupForm.ID,
		Name:      groupForm.Name,
		ServerIds: groupForm.Servers,
	}
	if err := group.Update(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.Success(c)
}

func GroupPath(c *gin.Context) {
	var groupForm GroupForm
	if err := c.ShouldBind(&groupForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	if groupForm.ID == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	group := &server.Group{
		ID:   groupForm.ID,
		Name: groupForm.Name,
	}
	if err := group.Update(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.Success(c)
}

func GroupPathAdd(c *gin.Context) {
	var groupForm GroupForm
	if err := c.ShouldBind(&groupForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	if groupForm.ID == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	group := &server.Group{
		ID:   groupForm.ID,
		Name: groupForm.Name,
	}
	if err := group.Update(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.Success(c)
}
