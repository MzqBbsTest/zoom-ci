package server

import (
	"errors"
	"github.com/zoom-ci/zoom-ci/server/model"
	query2 "github.com/zoom-ci/zoom-ci/server/query"
)

type ServerGroup struct {
	ID       int    `json:"id"`
	GroupId  string `json:"group_id"`
	ServerId int    `json:"server_id"`
	Ctime    int    `json:"ctime"`
}

func (g *ServerGroup) List(bind *query2.BindGroupServer) ([]model.ServerGroup, error) {

	where := query2.ParseGroupServerQuery(bind)
	serverGroup := model.ServerGroup{}
	serverGroupList, ok := serverGroup.List(model.QueryParam{
		Where: where,
	})

	if !ok {
		return nil, errors.New("get server group list failed")
	}

	return serverGroupList, nil
}

func (g *ServerGroup) Delete(bind *query2.BindGroupServer) bool {

	where := query2.ParseGroupServerQuery(bind)
	serverGroup := model.ServerGroup{}
	ok := serverGroup.DeleteWhere(model.QueryParam{
		Where: where,
	})
	return ok
}
