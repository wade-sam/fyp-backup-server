package policy

import "github.com/wade-sam/fyp-backup-server/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreatePolicy(policyname, backupType string, retention int, fullbackup, incrementalbackup []string, clients []entity.ID) (entity.ID, error) {
	policy, err := entity.NewPolicy(policyname, backupType, retention, fullbackup, incrementalbackup, clients)
	if err != nil {
		return policy.PolicyID, err
	}
	return s.repo.Create(policy)
}

func (s *Service) GetPolicy(name entity.ID) (*entity.Policy, error) {
	return s.repo.Get(name)
}

func (s *Service) ListPolicies() ([]*entity.Policy, error) {
	result, err := s.repo.List()
	if result == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return s.repo.List()
}

func (s *Service) UpdatePolicy(policy *entity.Policy) error {
	err := policy.ValidatePolicy()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(policy)
}

func (s *Service) DeletePolicy(name entity.ID) error {
	p, err := s.GetPolicy(name)
	if p == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(name)
}
