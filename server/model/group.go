// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import (
	"time"
)

type Group struct {
	ID    int    `gorm:"primary_key"`
	Name  string `gorm:"type:varchar(100);not null;default:''"`
	Ctime int    `gorm:"type:int(11);not null;default:0"`
}

func (m *Group) TableName() string {
	return GetTableName("group")
}

func (m *Group) Create() bool {
	m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *Group) Update() bool {
	return UpdateByPk(m)
}

func (m *Group) List(query QueryParam) ([]Group, bool) {
	var data []Group
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *Group) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *Group) Delete() bool {
	return DeleteByPk(m)
}

func (m *Group) Get(id int) bool {
	return GetByPk(m, id)
}
