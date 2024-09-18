package query

import "github.com/zoom-ci/zoom-ci/server/model"

type BindGroup struct {
	Keyword string `form:"keyword"`
	GroupId int    `form:"group_id"`
	Offset  int    `form:"offset"`
	Limit   int    `form:"limit,default=20" binding:"gte=1,lte=999" `
}

func ParseGroupQuery(c *BindGroup) []model.WhereParam {
	var builder WhereParamBuilder

	if len(c.Keyword) > 0 {
		builder.AddCondition("name", "like", c.Keyword)
	}

	if c.GroupId > 0 {
		builder.AddCondition("id", "=", c.GroupId)
	}

	return builder.GetParams()
}
