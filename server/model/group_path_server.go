// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

import (
	"time"
)

type GroupPathServer struct {
	ID            int `gorm:"primary_key"`
	ServerId      int `gorm:"type:int(11);not null;default:0"`
	ServerGroupId int `gorm:"type:int(11);not null;default:0"`
	ServerPathId  int `gorm:"type:int(11);not null;default:0"`
	Ctime         int `gorm:"type:int(11);not null;default:0"`
}

func (m *GroupPathServer) TableName() string {
	return GetTableName("server_group_path_server")
}

func (m *GroupPathServer) Create() bool {
	m.Ctime = int(time.Now().Unix())
	return Create(m)
}

func (m *GroupPathServer) Update() bool {
	return UpdateByPk(m)
}

func (m *GroupPathServer) List(query QueryParam) ([]GroupPathServer, bool) {
	var data []GroupPathServer
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *GroupPathServer) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *GroupPathServer) Delete() bool {
	return DeleteByPk(m)
}

func (m *GroupPathServer) DeleteWhere(query QueryParam) bool {
	return Delete(m, query)
}

func (m *GroupPathServer) Get(id int) bool {
	return GetByPk(m, id)
}
