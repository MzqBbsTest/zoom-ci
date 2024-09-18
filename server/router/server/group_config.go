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

func GroupConfig(c *gin.Context) {

	var query query2.BindGroupConfig
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	ser := &server.GroupConfig{}
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

func GroupConfigSave(c *gin.Context) {

	var groupForm form.GroupServerConfigForm
	if err := c.ShouldBind(&groupForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	groupConfig := server.GroupConfig{
		ID:       groupForm.Id,
		Path:     groupForm.Path,
		Name:     groupForm.Name,
		GroupId:  groupForm.GroupId,
		ServerId: groupForm.ServerId,
	}

	if err := groupConfig.Save(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.Success(c)
}

func GroupConfigDelete(c *gin.Context) {

	id := c.GetInt("id")
	if id == 0 {
		render.Success(c)
		return
	}

	groupConfig := server.GroupConfig{ID: id}
	if err := groupConfig.Delete(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.Success(c)
}

func GroupConfigDetail(c *gin.Context) {

	id := c.GetInt("id")
	if id == 0 {
		render.Success(c)
		return
	}

	groupConfig := server.GroupConfig{ID: id}
	if err := groupConfig.Detail(); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.JSON(c, groupConfig)
}
