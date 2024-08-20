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

type Cmd struct {
	ID      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Ctime   int    `json:"ctime"`
}

func (s *Cmd) CreateOrUpdate() error {
	Cmd := &model.Cmd{
		ID:      s.ID,
		UserId:  s.UserId,
		Title:   s.Title,
		Content: s.Content,
	}
	if Cmd.ID == 0 {
		if ok := Cmd.Create(); !ok {
			return errors.New("create Cmd failed")
		}
	} else {
		if ok := Cmd.Update(); !ok {
			return errors.New("update Cmd failed")
		}
	}
	return nil
}

func (s *Cmd) ListByUserId(keyword string, userId, offset, limit int) ([]Cmd, error) {
	cmd := &model.Cmd{}
	conds := s.parseWhereConds(keyword)
	conds = append(conds, model.WhereParam{
		Field:   "user_id",
		Prepare: userId,
	})
	list, ok := cmd.List(model.QueryParam{
		Fields: "id, title, content, user_id, ctime",
		Offset: offset,
		Limit:  limit,
		Order:  "id DESC",
		Where:  conds,
	})
	if !ok {
		return nil, errors.New("get Cmd list failed")
	}

	var (
		CmdList []Cmd
		userIds []int
	)
	for _, l := range list {
		CmdList = append(CmdList, Cmd{
			ID:      l.ID,
			Title:   l.Title,
			Content: l.Content,
			UserId:  l.UserId,
			Ctime:   l.Ctime,
		})
		userIds = append(userIds, l.UserId)
	}
	return CmdList, nil
}

func (s *Cmd) List(keyword string, offset, limit int) ([]Cmd, error) {
	cmd := &model.Cmd{}
	list, ok := cmd.List(model.QueryParam{
		Fields: "id, title, content, user_id, ctime",
		Offset: offset,
		Limit:  limit,
		Order:  "id DESC",
		Where:  s.parseWhereConds(keyword),
	})
	if !ok {
		return nil, errors.New("get Cmd list failed")
	}

	var (
		CmdList []Cmd
		userIds []int
	)
	for _, l := range list {
		CmdList = append(CmdList, Cmd{
			ID:      l.ID,
			Title:   l.Title,
			Content: l.Content,
			UserId:  l.UserId,
			Ctime:   l.Ctime,
		})
		userIds = append(userIds, l.UserId)
	}
	return CmdList, nil
}

func (s *Cmd) Total(keyword string) (int, error) {
	Cmd := model.Cmd{}
	total, ok := Cmd.Count(model.QueryParam{
		Where: s.parseWhereConds(keyword),
	})
	if !ok {
		return 0, errors.New("get Cmd count failed")
	}
	return total, nil
}

func (s *Cmd) Delete() error {
	Cmd := &model.Cmd{
		ID: s.ID,
	}
	if ok := Cmd.Delete(); !ok {
		return errors.New("delete Cmd failed")
	}
	return nil
}

func (s *Cmd) Detail() error {
	Cmd := &model.Cmd{}
	if ok := Cmd.Get(s.ID); !ok {
		return errors.New("get Cmd detail failed")
	}
	if Cmd.ID == 0 {
		return errors.New("Cmd not exists")
	}

	s.ID = Cmd.ID
	s.Title = Cmd.Title
	s.UserId = Cmd.UserId
	s.Content = Cmd.Content
	s.Ctime = Cmd.Ctime

	return nil
}

func (s *Cmd) parseWhereConds(keyword string) []model.WhereParam {
	var where []model.WhereParam
	if keyword != "" {
		if gois.IsInteger(keyword) {
			where = append(where, model.WhereParam{
				Field:   "id",
				Prepare: keyword,
			})
		} else {
			where = append(where, model.WhereParam{
				Field:   "title",
				Tag:     "LIKE",
				Prepare: fmt.Sprintf("%%%s%%", keyword),
			})
		}
	}
	return where
}
