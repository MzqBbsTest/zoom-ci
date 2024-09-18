package query

import "github.com/zoom-ci/zoom-ci/server/model"

type BindPathServer struct {
	Keyword   string `form:"keyword"`
	GroupId   int    `form:"group_id"`
	ServerIds []int  `form:"server_ids"`
	PathId    int    `form:"path_id"`
	Offset    int    `form:"offset"`
	Limit     int    `form:"limit,default=20" binding:"gte=1,lte=999" `
}

func ParsePathServerQuery(c *BindPathServer) []model.WhereParam {
	var builder WhereParamBuilder

	if c.GroupId > 0 {
		builder.AddCondition("server_group_id", "=", c.GroupId)
	}

	if c.PathId > 0 {
		builder.AddCondition("server_path_id", "=", c.PathId)
	}

	if len(c.ServerIds) > 0 {
		builder.AddCondition("server_id", "in", c.ServerIds)
	}

	return builder.GetParams()
}

type BindServerPath struct {
	Keyword       string `form:"keyword"`
	ServerPathIds []int  `form:"server_path_ids"`
}

func ParsePathQuery(c *BindServerPath) []model.WhereParam {
	var builder WhereParamBuilder

	if len(c.ServerPathIds) > 0 {
		builder.AddCondition("id", "in", c.ServerPathIds)
	}

	return builder.GetParams()
}
