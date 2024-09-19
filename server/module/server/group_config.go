// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"errors"
	"github.com/zoom-ci/zoom-ci/server/model"
	query2 "github.com/zoom-ci/zoom-ci/server/query"
)

type GroupConfig struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Ctime    int    `json:"ctime"`
	ServerId int    `json:"server_id"`
	GroupId  int    `json:"group_id"`
}

func (g *GroupConfig) Save() error {
	if g.ID > 0 {
		return g.Update()
	}
	return g.Create()
}

func (g *GroupConfig) Create() error {

	serverGroupConfig := model.GroupConfig{
		ServerId: g.ServerId,
		GroupId:  g.GroupId,
		Path:     g.Path,
		Name:     g.Name,
	}
	if ok := serverGroupConfig.Create(); !ok {
		return errors.New("create server group data failed")
	}

	return nil
}

func (g *GroupConfig) Update() error {

	serverGroupConfig := model.GroupConfig{
		ID:       g.ID,
		ServerId: g.ServerId,
		GroupId:  g.GroupId,
		Path:     g.Path,
		Name:     g.Name,
	}
	if ok := serverGroupConfig.Update(); !ok {
		return errors.New("update server group data failed")
	}

	return nil
}

func (g *GroupConfig) List(bind *query2.BindGroupConfig) ([]GroupConfig, error) {

	offset := bind.Offset
	limit := bind.Limit
	where := query2.ParseGroupConfigQuery(bind)

	groupConfig := model.GroupConfig{}
	groupConfigList, ok := groupConfig.List(model.QueryParam{
		Fields: "id, name, path, ctime, group_id, server_id",
		Offset: offset,
		Limit:  limit,
		Order:  "id DESC",
		Where:  where,
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	l := make([]GroupConfig, len(groupConfigList))
	for i, config := range groupConfigList {
		l[i] = GroupConfig{
			ID:       config.ID,
			Path:     config.Path,
			Name:     config.Name,
			Ctime:    config.Ctime,
			ServerId: config.ServerId,
		}
	}

	return l, nil
}

func (g *GroupConfig) Total(bind *query2.BindGroupConfig) (int, error) {

	where := query2.ParseGroupConfigQuery(bind)
	groupConfigServer := model.GroupConfig{}
	total, ok := groupConfigServer.Count(model.QueryParam{
		Where: where,
	})
	if !ok {
		return 0, errors.New("get group server count failed")
	}

	return total, nil
}

func (g *GroupConfig) Delete() error {

	groupConfigServer := model.GroupConfig{
		ID: g.ID,
	}
	ok := groupConfigServer.Delete()
	if !ok {
		return errors.New("delete server group failed")
	}

	return nil
}

func (g *GroupConfig) Detail() error {

	if g.ID == 0 || g.GroupId == 0 {
		return errors.New("server group config not empty")
	}

	groupConfig := &model.GroupConfig{}
	if ok := groupConfig.Get(g.ID); !ok {
		return errors.New("server group config  not found")
	}

	groupConfigServer := model.GroupConfig{}
	ok := groupConfigServer.Get(g.ID)
	if !ok {
		return errors.New("server group config  not found")
	}

	g.Name = groupConfig.Name
	g.Path = groupConfig.Path
	g.GroupId = groupConfig.GroupId
	g.ServerId = groupConfig.ServerId

	return nil
}
