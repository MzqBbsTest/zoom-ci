package model

import (
	"time"
)

type SSHKey struct {
	ID         int       `gorm:"primaryKey;column:id" json:"id"`
	UserID     int       `gorm:"column:user_id" json:"userId"`
	Name       string    `gorm:"column:name" json:"name"`
	Password   string    `gorm:"column:password" json:"password"`
	PublicKey  string    `gorm:"column:public_key" json:"publicKey"`
	PrivateKey string    `gorm:"column:private_key" json:"privateKey"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updated_at" json:"updatedAt"`
}

func (m *SSHKey) TableName() string {
	return GetTableName("ssh_key")
}

func (m *SSHKey) Create() bool {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return Create(m)
}

func (m *SSHKey) Update() bool {
	m.UpdatedAt = time.Now()
	return UpdateByPk(m)
}

func (m *SSHKey) List(query QueryParam) ([]SSHKey, bool) {
	var data []SSHKey
	ok := GetMulti(&data, query)
	return data, ok
}

func (m *SSHKey) Count(query QueryParam) (int, bool) {
	var count int
	ok := Count(m, &count, query)
	return count, ok
}

func (m *SSHKey) Delete() bool {
	return DeleteByPk(m)
}

func (m *SSHKey) Get(id int) bool {
	return GetByPk(m, id)
}
