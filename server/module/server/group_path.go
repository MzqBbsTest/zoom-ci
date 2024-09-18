// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"errors"
	"github.com/zoom-ci/zoom-ci/server/model"
	query2 "github.com/zoom-ci/zoom-ci/server/query"
	"github.com/zoom-ci/zoom-ci/util/utils"
)

type GroupPath struct {
	ID            int    `json:"id"`
	Path          string `json:"path"`
	Name          string `json:"name"`
	Ctime         int    `json:"ctime"`
	ServerIds     []int  `json:"server_ids"`
	ServerGroupId int    `json:"server_group_id"`
}

func (g *GroupPath) Save() error {
	if g.ID > 0 {
		return g.Update()
	}
	return g.Create()
}

func (g *GroupPath) Create() error {
	serverGroupPath := model.GroupPath{
		Path: g.Path,
		Name: g.Name,
	}
	if ok := serverGroupPath.Create(); !ok {
		return errors.New("create server group data failed")
	}

	for _, id := range g.ServerIds {
		groupPathServer := model.GroupPathServer{
			ServerPathId:  serverGroupPath.ID,
			ServerId:      id,
			ServerGroupId: g.ServerGroupId,
		}
		groupPathServer.Create()
	}

	return nil
}

func (g *GroupPath) Update() error {
	serverGroupPath := model.GroupPath{
		ID:   g.ID,
		Path: g.Path,
		Name: g.Name,
	}
	if ok := serverGroupPath.Update(); !ok {
		return errors.New("update server group data failed")
	}

	// 已存在
	groupPathServer := model.GroupPathServer{}
	where := query2.ParsePathServerQuery(&query2.BindPathServer{GroupId: g.ServerGroupId})
	groupPathServerList, ok := groupPathServer.List(model.QueryParam{Where: where})
	if !ok {
		return errors.New("query server group failed")
	}

	serverIds, err := utils.Pluck[model.GroupPathServer, int](groupPathServerList, "ServerId")
	if err != nil {
		return err
	}

	// 差集
	serverIds2 := utils.Difference(serverIds, g.ServerIds)

	// 删除不存在
	where2 := append(where, model.WhereParam{
		Field:   "server_id",
		Tag:     "in",
		Prepare: serverIds2,
	})
	ok = groupPathServer.DeleteWhere(model.QueryParam{Where: where2})
	if !ok {
		return errors.New("delete data failed")
	}

	// 新增
	serverIds3 := utils.Difference(g.ServerIds, serverIds)
	for _, id := range serverIds3 {
		serverGroupPath := model.GroupPathServer{
			ServerId:      id,
			ServerGroupId: g.ServerGroupId,
			ServerPathId:  g.ID,
		}
		serverGroupPath.Create()
	}

	return nil
}

func (g *GroupPath) List(bind *query2.BindPathServer) ([]GroupPath, error) {
	offset := bind.Offset
	limit := bind.Limit
	where := query2.ParsePathServerQuery(bind)

	groupPathServer := model.GroupPathServer{}
	groupPathServerList, ok := groupPathServer.List(model.QueryParam{
		Fields: "id, server_group_id, server_id, server_path_id",
		Where:  where,
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	var ids []int
	for _, group := range groupPathServerList {
		ids = append(ids, group.ServerPathId)
	}

	where = query2.ParsePathQuery(&query2.BindServerPath{ServerPathIds: ids, Keyword: bind.Keyword})
	groupPath := model.GroupPath{}
	groupPathList, ok := groupPath.List(model.QueryParam{
		Fields: "id, name, path, ctime",
		Offset: offset,
		Limit:  limit,
		Order:  "id DESC",
		Where:  where,
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	var groupList []GroupPath
	for _, path := range groupPathList {
		var serverIds []int
		for _, groupPathServer := range groupPathServerList {
			if path.ID == groupPathServer.ServerPathId {
				serverIds = append(serverIds, groupPathServer.ServerId)
			}
		}
		groupList = append(groupList, GroupPath{
			ID:            path.ID,
			Path:          path.Path,
			Ctime:         path.Ctime,
			ServerIds:     serverIds,
			ServerGroupId: bind.GroupId,
		})
	}

	return groupList, nil
}

func (g *GroupPath) Total(bind *query2.BindPathServer) (int, error) {

	where := query2.ParsePathServerQuery(bind)
	groupPathServer := model.GroupPathServer{}
	groupPathServerList, ok := groupPathServer.List(model.QueryParam{
		Fields: "id, server_group_id, server_id, server_path_id",
		Where:  where,
	})
	if !ok {
		return 0, errors.New("get server group list failed")
	}

	var ids []int
	for _, group := range groupPathServerList {
		ids = append(ids, group.ServerPathId)
	}

	where = query2.ParsePathQuery(&query2.BindServerPath{ServerPathIds: ids, Keyword: bind.Keyword})
	groupPath := model.GroupPath{}
	total, ok := groupPath.Count(model.QueryParam{
		Fields: "id, name, ctime",
		Order:  "id DESC",
		Where:  where,
	})

	if !ok {
		return 0, errors.New("get server group count failed")
	}
	return total, nil
}

func (g *GroupPath) Delete() error {

	where := query2.ParsePathServerQuery(&query2.BindPathServer{PathId: g.ID})
	groupPathServer := model.GroupPathServer{}
	ok := groupPathServer.DeleteWhere(model.QueryParam{Where: where})
	if !ok {
		return errors.New("delete server group failed")
	}

	group := &model.GroupPath{
		ID: g.ID,
	}
	if ok := group.Delete(); !ok {
		return errors.New("delete server group failed")
	}

	return nil
}

func (g *GroupPath) Detail() error {

	if g.ID == 0 || g.ServerGroupId == 0 {
		return errors.New("server group path not empty")
	}

	groupPath := &model.GroupPath{}
	if ok := groupPath.Get(g.ID); !ok {
		return errors.New("server group path  not found")
	}

	where := query2.ParsePathServerQuery(&query2.BindPathServer{
		PathId:  g.ID,
		GroupId: g.ServerGroupId,
	})

	groupPathServer := model.GroupPathServer{}
	groupPathServerList, ok := groupPathServer.List(model.QueryParam{
		Fields: "server_id",
		Where:  where,
	})
	if !ok {
		return errors.New("server group path  not found")
	}

	for _, groupPathServerInfo := range groupPathServerList {
		g.ServerIds = append(g.ServerIds, groupPathServerInfo.ServerId)
	}

	return nil
}
