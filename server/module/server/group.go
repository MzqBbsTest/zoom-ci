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

type Group struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Ctime     int    `json:"ctime"`
	ServerIds []int  `json:"server_ids"`
}

func (g *Group) GroupGetMapByIds(ids []int) (map[int]Group, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	group := &model.Group{}
	groupList, ok := group.List(model.QueryParam{
		Where: []model.WhereParam{
			model.WhereParam{
				Field:   "id",
				Tag:     "IN",
				Prepare: ids,
			},
		},
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	serverGroup := ServerGroup{}
	serverGroupList, err := serverGroup.List(&query2.BindGroupServer{
		GroupIds: ids,
	})
	if err != nil {
		return nil, errors.New("get server group list failed")
	}

	groupMap := make(map[int]Group)
	for _, l := range groupList {
		var serverIds []int
		serverIds, err = utils.Pluck[model.ServerGroup, int](serverGroupList, "ServerId")
		if err != nil {
			continue
		}
		groupMap[l.ID] = Group{
			ID:        l.ID,
			Name:      l.Name,
			Ctime:     l.Ctime,
			ServerIds: serverIds,
		}

	}
	return groupMap, nil
}

func (g *Group) Create() error {
	group := model.Group{
		Name: g.Name,
	}
	if ok := group.Create(); !ok {
		return errors.New("create server group data failed")
	}
	for _, id := range g.ServerIds {
		serverGroup := model.ServerGroup{
			ServerId: id,
			GroupId:  group.ID,
		}
		serverGroup.Create()
	}

	return nil
}

func (g *Group) Update() error {
	group := model.Group{
		ID:   g.ID,
		Name: g.Name,
	}
	if ok := group.Update(); !ok {
		return errors.New("update server group data failed")
	}

	serverGroup := ServerGroup{}
	serverGroupList, err := serverGroup.List(&query2.BindGroupServer{
		GroupId: g.ID,
	})

	if err != nil {
		return err
	}

	serverIds, err := utils.Pluck[model.ServerGroup, int](serverGroupList, "server_id")
	if err != nil {
		return err
	}

	serverIds2 := utils.Difference(serverIds, g.ServerIds)

	ok := serverGroup.Delete(&query2.BindGroupServer{
		GroupId:   g.ID,
		ServerIds: serverIds2,
	})
	if !ok {
		return errors.New("delete is error")
	}

	for _, id := range g.ServerIds {
		serverGroup := model.ServerGroup{
			ServerId: id,
			GroupId:  g.ID,
		}
		serverGroup.Create()
	}

	return nil
}

func (g *Group) List(bind *query2.BindGroup) ([]Group, error) {

	where := query2.ParseGroupQuery(bind)
	group := model.Group{}
	list, ok := group.List(model.QueryParam{
		Fields: "id, name, ctime",
		Offset: bind.Offset,
		Limit:  bind.Limit,
		Order:  "id DESC",
		Where:  where,
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	var ids []int
	for _, group := range list {
		ids = append(ids, group.ID)
	}

	serverGroup := ServerGroup{}
	serverGroupList, err := serverGroup.List(&query2.BindGroupServer{GroupIds: ids})
	if err != nil {
		return nil, errors.New("get server group list failed")
	}

	var groupList []Group
	for _, l := range list {
		var serverIds []int
		for _, server := range serverGroupList {
			if l.ID == server.GroupId {
				serverIds = append(serverIds, server.ServerId)
			}
		}
		groupList = append(groupList, Group{
			ID:        l.ID,
			Name:      l.Name,
			Ctime:     l.Ctime,
			ServerIds: serverIds,
		})
	}

	return groupList, nil
}

func (g *Group) Total(bind *query2.BindGroup) (int, error) {

	where := query2.ParseGroupQuery(bind)
	group := model.Group{}
	total, ok := group.Count(model.QueryParam{
		Where: where,
	})

	if !ok {
		return 0, errors.New("get server group count failed")
	}

	return total, nil
}

func (g *Group) Delete() error {
	group := &model.Group{
		ID: g.ID,
	}
	if ok := group.Delete(); !ok {
		return errors.New("delete server group failed")
	}
	return nil
}

func (g *Group) Detail() error {
	group := model.Group{}
	if ok := group.Get(g.ID); !ok {
		return errors.New("get server group detail failed")
	}
	if group.ID == 0 {
		return errors.New("server group not exists")
	}

	g.ID = group.ID
	g.Name = group.Name
	g.Ctime = group.Ctime

	serverGroup := ServerGroup{}

	serverGroupList, err := serverGroup.List(&query2.BindGroupServer{GroupId: g.ID})
	if err != nil {
		return errors.New("get server group list failed")
	}

	serverIds, err := utils.Pluck[model.ServerGroup, int](serverGroupList, "ServerId")
	if err != nil {
		return err
	}

	g.ServerIds = serverIds

	return nil
}
