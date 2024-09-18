package query

import "github.com/zoom-ci/zoom-ci/server/model"

type WhereParamBuilder struct {
	Params []model.WhereParam
}

// AddCondition 增加一个查询条件
func (b *WhereParamBuilder) AddCondition(field string, operator string, value interface{}) {
	if value != nil && value != "" {
		b.Params = append(b.Params, model.WhereParam{
			Field:   field,
			Prepare: value,
			Tag:     operator,
		})
	}
}

// GetParams 返回构造好的WhereParam
func (b *WhereParamBuilder) GetParams() []model.WhereParam {
	return b.Params
}
