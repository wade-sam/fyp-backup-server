package policy

import (
	"github.com/wade-sam/fyp-backup-server/entity"
)

type policyholder struct {
	m map[entity.ID]*entity.Policy
}

func NewPolicyHolder() *policyholder {
	var p = map[entity.ID]*entity.Policy{}
	return &policyholder{
		m: p,
	}
}

func (r *policyholder) Create(policy *entity.Policy) (entity.ID, error) {
	r.m[policy.PolicyID] = policy
	return policy.PolicyID, nil
}

func (r *policyholder) Get(p entity.ID) (*entity.Policy, error) {
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

func (r *policyholder) Delete(id entity.ID) error {
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
