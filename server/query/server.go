package query

import (
	"fmt"
	"github.com/zoom-ci/zoom-ci/server/model"
)

type BindServer struct {
	Keyword   string `form:"keyword"`
	ServerId  int    `form:"server_id"`
	ServerIds []int  `form:"server_ids"`
	Offset    int    `form:"offset"`
	Limit     int    `form:"limit,default=20" binding:"gte=1,lte=999" `
}

func ParseServerQuery(c *BindServer) []model.WhereParam {
	var builder WhereParamBuilder

	if len(c.Keyword) > 0 {
		builder.AddCondition("name", "like", fmt.Sprintf("%%%s%%", c.Keyword))
	}

	if c.ServerId > 0 {
		builder.AddCondition("id", "=", c.ServerId)
	}

	if len(c.ServerIds) > 0 {
		builder.AddCondition("id", "IN", c.ServerIds)

	}

	return builder.GetParams()
}
