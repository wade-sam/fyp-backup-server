package client

import (
	"github.com/wade-sam/fyp-backup-server/entity"
)

type clientholder struct {
	m map[string]*entity.Client
}

func NewClientHolder() *clientholder {
	var p = map[string]*entity.Client{}
	return &clientholder{
		m: p,
	}
}

func (r *clientholder) Create(client *entity.Client) (string, error) {
	r.m[client.ConsumerID] = client
	return client.ConsumerID, nil
}

func (r *clientholder) Get(p string) (*entity.Client, error) {
	if r.m[p] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[p], nil
}

func (r *clientholder) List() ([]*entity.Client, error) {
	var d []*entity.Client
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

func (r *clientholder) Delete(id string) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

func (r *clientholder) Update(client *entity.Client) error {
	_, err := r.Get(client.ConsumerID)
	if err != nil {
		return err
	}
	r.m[client.ConsumerID] = client
	return nil
}
