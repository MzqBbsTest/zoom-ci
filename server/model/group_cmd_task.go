// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import (
	"time"
)

type GroupCmdTask struct {
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

func (m *GroupCmdTask) TableName() string {
	return GetTableName("group_cmd_task")
}

func (m *GroupCmdTask) Create() bool {
	m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *GroupCmdTask) Update() bool {
	return UpdateByPk(m)
}

func (m *GroupCmdTask) List(query QueryParam) ([]GroupCmdTask, bool) {
	var data []GroupCmdTask
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *GroupCmdTask) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *GroupCmdTask) Delete() bool {
	return DeleteByPk(m)
}

func (m *GroupCmdTask) Get(id int) bool {
	return GetByPk(m, id)
}
