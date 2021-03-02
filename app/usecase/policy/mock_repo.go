package policy

import (
	"github.com/wade-sam/fyp-backup-server/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type policyholder struct {
	m map[string]*entity.Policy
}

func NewPolicyHolder() *policyholder {
	var p = map[string]*entity.Policy{}
	return &policyholder{
		m: p,
	}
}

func (r *policyholder) Create(policy *entity.Policy) (string, error) {
	id := primitive.NewObjectID().Hex()
	r.m[id] = policy
	return id, nil
}

func (r *policyholder) Get(p string) (*entity.Policy, error) {
	if r.m[p] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[p], nil
}

func (r *policyholder) List() ([]*entity.Policy, error) {
	var d []*entity.Policy
	for _, j := range r.m {
		d = append(d, j)
	}
	return d, nil
}

func (r *policyholder) Delete(id string) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

func (r *policyholder) Update(policy *entity.Policy) error {
	_, err := r.Get(policy.PolicyID)
	if err != nil {
		return err
	}
	r.m[policy.PolicyID] = policy
	return nil
}
