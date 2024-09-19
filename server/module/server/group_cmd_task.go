// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"errors"
	"github.com/zoom-ci/zoom-ci/server/model"
	query2 "github.com/zoom-ci/zoom-ci/server/query"
)

type GroupCmdTask struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	GroupConfigAlias string `json:"group_config_alias"`
	GroupPathAlias   string `json:"group_path_alias"`
	StartCommand     string `json:"start_command"`
	StartUser        string `json:"start_user"`
	Ctime            int    `json:"ctime"`
	ServerId         int    `json:"server_id"`
	GroupId          int    `json:"group_id"`
	GroupCmdId       int    `json:"group_cmd_id"`
}

func (g *GroupCmdTask) Save() error {
	if g.ID > 0 {
		return g.Update()
	}
	return g.Create()
}

func (g *GroupCmdTask) Create() error {
	serverGroupCmdTask := model.GroupCmdTask{
		ServerId:         g.ServerId,
		GroupId:          g.GroupId,
		GroupConfigAlias: g.GroupConfigAlias,
		GroupPathAlias:   g.GroupPathAlias,
		StartCommand:     g.StartCommand,
		StartUser:        g.StartUser,
		Name:             g.Name,
	}
	if ok := serverGroupCmdTask.Create(); !ok {
		return errors.New("create group cmd data failed")
	}

	return nil
}

func (g *GroupCmdTask) Update() error {
	serverGroupCmdTask := model.GroupCmdTask{
		ID:               g.ID,
		ServerId:         g.ServerId,
		GroupId:          g.GroupId,
		GroupConfigAlias: g.GroupConfigAlias,
		GroupPathAlias:   g.GroupPathAlias,
		StartCommand:     g.StartCommand,
		StartUser:        g.StartUser,
		Name:             g.Name,
	}
	if ok := serverGroupCmdTask.Update(); !ok {
		return errors.New("update group cmd data failed")
	}

	return nil
}

func (g *GroupCmdTask) List(bind *query2.BindGroupCmdTask) ([]GroupCmdTask, error) {
	offset := bind.Offset
	limit := bind.Limit
	where := query2.ParseGroupCmdTaskQuery(bind)

	groupCmd := model.GroupCmdTask{}
	groupCmdList, ok := groupCmd.List(model.QueryParam{
		Fields: "id, name, group_config_alias, group_path_alias, start_command, start_user, group_id, server_id",
		Offset: offset,
		Limit:  limit,
		Order:  "id DESC",
		Where:  where,
	})
	if !ok {
		return nil, errors.New("get group cmd list failed")
	}

	l := make([]GroupCmdTask, len(groupCmdList))
	for i, path := range groupCmdList {
		l[i] = GroupCmdTask{
			ID:               path.ID,
			GroupConfigAlias: path.GroupConfigAlias,
			GroupPathAlias:   path.GroupPathAlias,
			StartCommand:     path.StartCommand,
			StartUser:        path.StartUser,
			GroupId:          path.GroupId,
			Name:             path.Name,
			Ctime:            path.Ctime,
			ServerId:         path.ServerId,
		}
	}

	return l, nil
}

func (g *GroupCmdTask) Total(bind *query2.BindGroupCmdTask) (int, error) {

	where := query2.ParseGroupCmdTaskQuery(bind)
	groupCmdServer := model.GroupCmdTask{}
	total, ok := groupCmdServer.Count(model.QueryParam{
		Where: where,
	})
	if !ok {
		return 0, errors.New("get group cmd count failed")
	}

	return total, nil
}

func (g *GroupCmdTask) Delete() error {

	groupCmdServer := model.GroupCmdTask{
		ID: g.ID,
	}
	ok := groupCmdServer.Delete()
	if !ok {
		return errors.New("delete group cmd failed")
	}

	return nil
}

func (g *GroupCmdTask) Detail() error {

	if g.ID == 0 || g.GroupId == 0 {
		return errors.New("group cmd path not empty")
	}

	groupCmd := &model.GroupCmdTask{}
	if ok := groupCmd.Get(g.ID); !ok {
		return errors.New("group cmd path  not found")
	}

	groupCmdServer := model.GroupCmdTask{}
	ok := groupCmdServer.Get(g.ID)
	if !ok {
		return errors.New("group cmd path  not found")
	}

	g.Name = groupCmd.Name
	g.GroupConfigAlias = groupCmd.GroupConfigAlias
	g.GroupPathAlias = groupCmd.GroupPathAlias
	g.StartCommand = groupCmd.StartCommand
	g.StartUser = groupCmd.StartUser
	g.GroupId = groupCmd.GroupId
	g.ServerId = groupCmd.ServerId

	return nil
}