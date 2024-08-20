// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import (
	"time"
)

type Cmd struct {
	ID      int    `gorm:"primary_key"`
	UserId  int    `gorm:"type:int(11);not null;default:0"`
	Title   string `gorm:"type:varchar(50);not null;default:''"`
	Content string `gorm:"type:varchar(1000);not null;default:''"`
	Ctime   int    `gorm:"type:int(11);not null;default:0"`
}

func (m *Cmd) TableName() string {
	return GetTableName("cmd")
}

func (m *Cmd) Create() bool {
	m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *Cmd) Update() bool {
	return UpdateByPk(m)
}

func (m *Cmd) UpdateByFields(data map[string]interface{}, query QueryParam) bool {
	return Update(m, data, query)
}

func (m *Cmd) List(query QueryParam) ([]Cmd, bool) {
	var data []Cmd
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *Cmd) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *Cmd) Delete() bool {
	return DeleteByPk(m)
}

func (m *Cmd) Get(id int) bool {
	return GetByPk(m, id)
}

func (m *Cmd) GetOne(query QueryParam) bool {
	return GetOne(m, query)
}
