package query

import (
	"github.com/zoom-ci/zoom-ci/server/model"
)

type BindGroupCmd struct {
	GroupId  int `form:"group_id"`
	ServerId int `form:"server_id"`
	Offset   int `form:"offset"`
	Limit    int `form:"limit,default=20" binding:"gte=1,lte=999" `
}

func ParseGroupCmdQuery(c *BindGroupCmd) []model.WhereParam {
	var builder WhereParamBuilder

	if c.GroupId > 0 {
		builder.AddCondition("group_id", "=", c.GroupId)
	}

	if c.ServerId > 0 {
		builder.AddCondition("server_id", "=", c.ServerId)
	}

	if c.ServerId == -1 {
		builder.AddCondition("server_id", "=", 0)
	}

	return builder.GetParams()
}
