// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import (
	"time"
)

type GroupConfig struct {
	ID       int    `gorm:"primary_key"`
	Name     string `gorm:"type:varchar(255);not null;default:''"`
	Path     string `gorm:"type:varchar(255);not null;default:''"`
	ServerId int    `gorm:"type:int(11);not null;default:0"`
	GroupId  int    `gorm:"type:int(11);not null;default:0"`
	Ctime    int    `gorm:"type:int(11);not null;default:0"`
}

func (m *GroupConfig) TableName() string {
	return GetTableName("group_config")
}

func (m *GroupConfig) Create() bool {
	m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *GroupConfig) Update() bool {
	return UpdateByPk(m)
}

func (m *GroupConfig) List(query QueryParam) ([]GroupConfig, bool) {
	var data []GroupConfig
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *GroupConfig) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *GroupConfig) Delete() bool {
	return DeleteByPk(m)
}

func (m *GroupConfig) Get(id int) bool {
	return GetByPk(m, id)
}
