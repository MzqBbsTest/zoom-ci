// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package model

type ServerGroup struct {
	ID       int `gorm:"primary_key"`
	ServerId int `gorm:"type:int(11);not null;default:0"`
	GroupId  int `gorm:"type:int(11);not null;default:0"`
}

func (m *ServerGroup) TableName() string {
	return GetTableName("server_group")
}

func (m *ServerGroup) Create() bool {
	return Create(m)
}

func (m *ServerGroup) Update() bool {
	return UpdateByPk(m)
}

func (m *ServerGroup) List(query QueryParam) ([]ServerGroup, bool) {
	var data []ServerGroup
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *ServerGroup) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *ServerGroup) Delete() bool {
	return DeleteByPk(m)
}

func (m *ServerGroup) Get(id int) bool {
	return GetByPk(m, id)
}
