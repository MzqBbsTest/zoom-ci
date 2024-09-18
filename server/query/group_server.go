package query

import "github.com/zoom-ci/zoom-ci/server/model"

type BindGroupServer struct {
	Keyword   string `form:"keyword"`
	GroupId   int    `form:"group_id"`
	GroupIds  []int  `form:"group_ids"`
	ServerIds []int  `form:"server_ids"`
	ServerId  int    `form:"server_id"`
}

func ParseGroupServerQuery(c *BindGroupServer) []model.WhereParam {
	var builder WhereParamBuilder

	if c.GroupId > 0 {
		builder.AddCondition("group_id", "=", c.GroupId)
	}

	if c.ServerId > 0 {
		builder.AddCondition("server_id", "=", c.ServerId)
	}

	if len(c.GroupIds) > 0 {
		builder.AddCondition("group_id", "in", c.GroupIds)
	}

	if len(c.ServerIds) > 0 {
		builder.AddCondition("server_id", "in", c.ServerIds)
	}

	return builder.GetParams()
}
