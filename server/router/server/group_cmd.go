// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/form"
	"github.com/zoom-ci/zoom-ci/server/module/server"
	query2 "github.com/zoom-ci/zoom-ci/server/query"
	"github.com/zoom-ci/zoom-ci/server/render"
)

func GroupCmd(c *gin.Context) {

	var query query2.BindGroupCmd
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	ser := &server.GroupCmd{}
	list, err := ser.List(&query)
	if err != nil {
		render.AppError(c, err.Error())
		return
	}

	total, err := ser.Total(&query)
	if err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func GroupCmdSave(c *gin.Context) {

	var groupForm form.GroupServerCmdForm
	if err := c.ShouldBind(&groupForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	groupCmd := server.GroupCmd{
		ID:               groupForm.Id,
		Name:             groupForm.Name,
		GroupConfigAlias: groupForm.GroupConfigAlias,
		GroupPathAlias:   groupForm.GroupPathAlias,
		StartCommand:     groupForm.StartCommand,
		StartUser:        groupForm.StartUser,
		GroupId:          groupForm.GroupId,
		ServerId:         groupForm.ServerId,
	}

	if err := groupCmd.Save(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.Success(c)
}

func GroupCmdDelete(c *gin.Context) {

	id := c.GetInt("id")
	if id == 0 {
		render.Success(c)
		return
	}

	groupCmd := server.GroupCmd{ID: id}
	if err := groupCmd.Delete(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.Success(c)
}

func GroupCmdDetail(c *gin.Context) {

	id := c.GetInt("id")
	if id == 0 {
		render.Success(c)
		return
	}

	groupCmd := server.GroupCmd{ID: id}
	if err := groupCmd.Detail(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.JSON(c, groupCmd)
}

func GroupCmdStart(c *gin.Context) {
	id := c.GetInt("id")
	if id == 0 {
		render.Success(c)
		return
	}

	groupCmd := server.GroupCmd{ID: id}
	if err := groupCmd.Detail(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	query := query2.BindGroupCmdTask{GroupCmdId: groupCmd.ID, Status: -2}

	task := server.GroupCmdTask{
		GroupCmdId:       groupCmd.ID,
		Name:             groupCmd.Name,
		GroupConfigAlias: groupCmd.GroupConfigAlias,
		GroupPathAlias:   groupCmd.GroupPathAlias,
		StartCommand:     groupCmd.StartCommand,
		StartUser:        groupCmd.StartUser,
		ServerId:         groupCmd.ServerId,
		GroupId:          groupCmd.GroupId,
	}

	total, err := task.Total(&query)
	if err != nil {
		render.ParamError(c, err.Error())
		return
	}

	if total > 0 {
		render.ParamError(c, "已存在正在执行的任务")
		return
	}

	if err := task.Create(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.Success(c)
}
