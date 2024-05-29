// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"errors"
	"fmt"

	"github.com/zoom-ci/zoom-ci/server/model"
	"github.com/zoom-ci/zoom-ci/util/gois"
)

type Group struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Ctime     int    `json:"ctime"`
	ServerIds []int  `json:"server_ids"`
}

func GroupGetMapByIds(ids []int) (map[int]Group, error) {
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
	serverGroup := model.ServerGroup{}
	serverGroupList, ok := serverGroup.List(model.QueryParam{
		Where: []model.WhereParam{
			model.WhereParam{
				Field:   "server_id",
				Tag:     "IN",
				Prepare: ids,
			},
		},
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	groupMap := make(map[int]Group)
	for _, l := range groupList {
		var serverIds []int
		for _, server := range serverGroupList {
			if l.ID == server.GroupId {
				serverIds = append(serverIds, server.ServerId)
			}
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
	serverGroup := model.Group{
		Name: g.Name,
	}
	if ok := serverGroup.Create(); !ok {
		return errors.New("create server group data failed")
	}
	return nil
}

func (g *Group) Update() error {
	serverGroup := model.Group{
		ID:   g.ID,
		Name: g.Name,
	}
	if ok := serverGroup.Update(); !ok {
		return errors.New("update server group data failed")
	}
	return nil
}

func (g *Group) List(keyword string, offset, limit int) ([]Group, error) {
	group := model.Group{}
	list, ok := group.List(model.QueryParam{
		Fields: "id, name, ctime",
		Offset: offset,
		Limit:  limit,
		Order:  "id DESC",
		Where:  g.parseWhereConds(keyword),
	})
	if !ok {
		return nil, errors.New("get server group list failed")
	}

	var groupList []Group
	for _, l := range list {
		groupList = append(groupList, Group{
			ID:    l.ID,
			Name:  l.Name,
			Ctime: l.Ctime,
		})
	}
	return groupList, nil
}

func (g *Group) Total(keyword string) (int, error) {
	group := model.Group{}
	total, ok := group.Count(model.QueryParam{
		Where: g.parseWhereConds(keyword),
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

	return nil
}

func (g *Group) parseWhereConds(keyword string) []model.WhereParam {
	var where []model.WhereParam
	if keyword != "" {
		if gois.IsInteger(keyword) {
			where = append(where, model.WhereParam{
				Field:   "id",
				Prepare: keyword,
			})
		} else {
			where = append(where, model.WhereParam{
				Field:   "name",
				Tag:     "LIKE",
				Prepare: fmt.Sprintf("%%%s%%", keyword),
			})
		}
	}
	return where
}
