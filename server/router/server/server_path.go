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

func GroupPath(c *gin.Context) {
	var query query2.BindPathServer
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	ser := &server.GroupPath{}
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

func GroupPathSave(c *gin.Context) {
	var groupForm form.GroupServerPathForm
	if err := c.ShouldBind(&groupForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	groupPath := server.GroupPath{
		ID:            groupForm.PathId,
		Path:          groupForm.Path,
		Name:          groupForm.Name,
		ServerGroupId: groupForm.GroupId,
		ServerIds:     groupForm.ServerIds,
	}

	if err := groupPath.Save(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.Success(c)
}

func GroupPathDelete(c *gin.Context) {

	id := c.GetInt("id")
	if id == 0 {
		render.Success(c)
		return
	}

	groupPath := server.GroupPath{ID: id}
	if err := groupPath.Delete(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.Success(c)
}

func GroupPathDetail(c *gin.Context) {

	id := c.GetInt("id")
	if id == 0 {
		render.Success(c)
		return
	}

	groupPath := server.GroupPath{ID: id}
	if err := groupPath.Detail(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.JSON(c, groupPath)
}
