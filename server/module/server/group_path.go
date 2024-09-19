// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"errors"
	"github.com/zoom-ci/zoom-ci/server/model"
	query2 "github.com/zoom-ci/zoom-ci/server/query"
)

type GroupPath struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Alias    string `json:"alias"`
	Name     string `json:"name"`
	Ctime    int    `json:"ctime"`
	ServerId int    `json:"server_id"`
	GroupId  int    `json:"group_id"`
}

func (g *GroupPath) Save() error {
	if g.ID > 0 {
		return g.Update()
	}
	return g.Create()
}

func (g *GroupPath) Create() error {
	serverGroupPath := model.GroupPath{
		ServerId: g.ServerId,
		GroupId:  g.GroupId,
		Path:     g.Path,
		Name:     g.Name,
	}
	if ok := serverGroupPath.Create(); !ok {
		return errors.New("create server group data failed")
	}

	return nil
}

func (g *GroupPath) Update() error {
	serverGroupPath := model.GroupPath{
		ID:       g.ID,
		ServerId: g.ServerId,
		GroupId:  g.GroupId,
		Path:     g.Path,
		Name:     g.Name,
	}
	if ok := serverGroupPath.Update(); !ok {
		return errors.New("update server group data failed")
	}

	return nil
}

func (g *GroupPath) ListAlias(bind *query2.BindGroupPath) ([]GroupPath, error) {
	where := query2.ParseGroupPathQuery(bind)

	groupPath := model.GroupPath{}
	groupPathList, ok := groupPath.List(model.QueryParam{
		Fields: "id, alias, group_id, server_id",
		Order:  "id DESC",
		Group:  "alias",
		Where:  where,
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	l := make([]GroupPath, len(groupPathList))
	for i, path := range groupPathList {
		l[i] = GroupPath{
			ID:       path.ID,
			Alias:    path.Alias,
			ServerId: path.ServerId,
		}
	}

	return l, nil
}

func (g *GroupPath) List(bind *query2.BindGroupPath) ([]GroupPath, error) {
	offset := bind.Offset
	limit := bind.Limit
	where := query2.ParseGroupPathQuery(bind)

	groupPath := model.GroupPath{}
	groupPathList, ok := groupPath.List(model.QueryParam{
		Fields: "id, name, alias, path, ctime, group_id, server_id",
		Offset: offset,
		Limit:  limit,
		Order:  "id DESC",
		Where:  where,
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	l := make([]GroupPath, len(groupPathList))
	for i, path := range groupPathList {
		l[i] = GroupPath{
			ID:       path.ID,
			Path:     path.Path,
			Name:     path.Name,
			Ctime:    path.Ctime,
			ServerId: path.ServerId,
		}
	}

	return l, nil
}

func (g *GroupPath) Total(bind *query2.BindGroupPath) (int, error) {

	where := query2.ParseGroupPathQuery(bind)
	groupPathServer := model.GroupPath{}
	total, ok := groupPathServer.Count(model.QueryParam{
		Where: where,
	})
	if !ok {
		return 0, errors.New("get group server count failed")
	}

	return total, nil
}

func (g *GroupPath) Delete() error {

	groupPathServer := model.GroupPath{
		ID: g.ID,
	}
	ok := groupPathServer.Delete()
	if !ok {
		return errors.New("delete server group failed")
	}

	return nil
}

func (g *GroupPath) Detail() error {

	if g.ID == 0 || g.GroupId == 0 {
		return errors.New("server group path not empty")
	}

	groupPath := &model.GroupPath{}
	if ok := groupPath.Get(g.ID); !ok {
		return errors.New("server group path  not found")
	}

	groupPathServer := model.GroupPath{}
	ok := groupPathServer.Get(g.ID)
	if !ok {
		return errors.New("server group path  not found")
	}

	g.Name = groupPath.Name
	g.Path = groupPath.Path
	g.GroupId = groupPath.GroupId
	g.ServerId = groupPath.ServerId

	return nil
}
