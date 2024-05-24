package table

import (
	"github.com/zoom-ci/zoom-ci/server/model"
)

type DataInterface interface {
	List(where []model.WhereParam, offset, limit int) ([]interface{}, error)
	Total(where []model.WhereParam) (int, error)
}

type Page struct {
	Where  []model.WhereParam
	Offset int
	Limit  int
}

func (p *Page) Table(table DataInterface) (map[string]interface{}, error) {
	list, err := table.List(p.Where, p.Offset, p.Limit)
	if err != nil {
		return nil, err
	}

	total, err := table.Total(p.Where)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"list":  list,
		"total": total,
	}, nil
}
