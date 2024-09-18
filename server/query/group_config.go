package query

import (
	"fmt"
	"github.com/zoom-ci/zoom-ci/server/model"
)

type BindGroupConfig struct {
	Keyword  string `form:"keyword"`
	GroupId  int    `form:"group_id"`
	ServerId int    `form:"server_id"`
	PathId   int    `form:"path_id"`
	Offset   int    `form:"offset"`
	Limit    int    `form:"limit,default=20" binding:"gte=1,lte=999" `
}

func ParseGroupConfigQuery(c *BindGroupConfig) []model.WhereParam {
	var builder WhereParamBuilder

	if c.GroupId > 0 {
		builder.AddCondition("group_id", "=", c.GroupId)
	}

	if len(c.Keyword) > 0 {
		builder.AddCondition("name", "like", fmt.Sprintf("%%%s%%", c.Keyword))
	}

	if c.ServerId > 0 {
		builder.AddCondition("server_id", "=", c.ServerId)
	}

	return builder.GetParams()
}
