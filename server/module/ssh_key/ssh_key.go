package ssh_key

import (
	"errors"
	"github.com/zoom-ci/zoom-ci/server/model"
	"time"
)

type SSHKey struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userId"`
	Name       string    `json:"name"`
	PublicKey  string    `json:"publicKey"`
	PrivateKey string    `json:"privateKey"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (p *SSHKey) Total(where []model.WhereParam) (int, error) {
	sshKey := &model.SSHKey{}
	total, ok := sshKey.Count(model.QueryParam{
		Where: where,
	})
	if !ok {
		return 0, errors.New("get sshKey count failed")
	}
	return total, nil
}

func (p *SSHKey) List(where []model.WhereParam, offset, limit int) ([]interface{}, error) {
	sshKey := &model.SSHKey{}
	list, ok := sshKey.List(model.QueryParam{
		Fields: "id, name",
		Offset: offset,
		Limit:  limit,
		Order:  "id DESC",
		Where:  where,
	})
	if !ok {
		return nil, errors.New("get sshKey list failed")
	}

	var projList []interface{}
	for _, l := range list {
		projList = append(projList, SSHKey{
			ID:   l.ID,
			Name: l.Name,
			//NeedAudit: l.NeedAudit,
			//Status:    l.Status,
		})
	}
	return projList, nil
}

func (p *SSHKey) CreateOrUpdate() error {
	key := &model.SSHKey{
		ID:         p.ID,
		Name:       p.Name,
		UserID:     p.UserID,
		PublicKey:  p.PublicKey,
		PrivateKey: p.PrivateKey,
	}
	if key.ID > 0 {
		if ok := key.Update(); !ok {
			return errors.New("sshKey update failed")
		}
	} else {
		if ok := key.Create(); !ok {
			return errors.New("sshKey create failed")
		}
	}
	return nil
}

func (p *SSHKey) Delete() error {
	sshKey := &model.SSHKey{
		ID: p.ID,
	}
	if ok := sshKey.Delete(); !ok {
		return errors.New("sshKey delete failed")
	}
	return nil
}

func (p *SSHKey) Detail() error {
	sshKey := &model.SSHKey{}
	if ok := sshKey.Get(p.ID); !ok {
		return errors.New("get sshKey detail failed")
	}
	if sshKey.ID == 0 {
		return errors.New("sshKey detail not exists")
	}

	p.ID = sshKey.ID
	p.UserID = sshKey.UserID
	p.PublicKey = sshKey.PublicKey
	p.PrivateKey = sshKey.PrivateKey

	return nil
}
