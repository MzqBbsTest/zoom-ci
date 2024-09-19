// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import (
	"time"
)

type GroupCmd struct {
	ID               int    `gorm:"primary_key"`
	Name             string `gorm:"type:varchar(255);not null;default:''"`
	GroupConfigAlias string `gorm:"type:varchar(255);not null;default:''"`
	GroupPathAlias   string `gorm:"type:varchar(255);not null;default:''"`
	StartCommand     string `gorm:"type:varchar(255);not null;default:''"`
	StartUser        string `gorm:"type:varchar(255);not null;default:''"`
	ServerId         int    `gorm:"type:int(11);not null;default:0"`
	GroupId          int    `gorm:"type:int(11);not null;default:0"`
	Ctime            int    `gorm:"type:int(11);not null;default:0"`
}

func (m *GroupCmd) TableName() string {
	return GetTableName("group_cmd")
}

func (m *GroupCmd) Create() bool {
	m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *GroupCmd) Update() bool {
	return UpdateByPk(m)
}

func (m *GroupCmd) List(query QueryParam) ([]GroupCmd, bool) {
	var data []GroupCmd
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *GroupCmd) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *GroupCmd) Delete() bool {
	return DeleteByPk(m)
}

func (m *GroupCmd) Get(id int) bool {
	return GetByPk(m, id)
}
