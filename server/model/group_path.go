// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import (
	"time"
)

type GroupPath struct {
	ID       int    `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(255);not null;default:''"`
	Path     string `gorm:"type:varchar(255);not null;default:''"`
	ServerId int    `gorm:"type:int(11);not null;default:0"`
	GroupId  int    `gorm:"type:int(11);not null;default:0"`
	Ctime    int    `gorm:"type:int(11);not null;default:0"`
}

func (m *GroupPath) TableName() string {
	return GetTableName("group_path")
}

func (m *GroupPath) Create() bool {
	m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *GroupPath) Update() bool {
	return UpdateByPk(m)
}

func (m *GroupPath) List(query QueryParam) ([]GroupPath, bool) {
	var data []GroupPath
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *GroupPath) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *GroupPath) Delete() bool {
	return DeleteByPk(m)
}

func (m *GroupPath) Get(id int) bool {
	return GetByPk(m, id)
}
