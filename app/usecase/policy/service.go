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

func (s *Service) CreatePolicy(policyname, backupType string, retention int, fullbackup, incrementalbackup, clients []string) error {
	policy, err := entity.NewPolicy(policyname, backupType, retention, fullbackup, incrementalbackup, clients)
	if err != nil {
		return err
	}
	return s.repo.Create(policy)
}

func (s *Service) GetPolicy(name string) (*entity.Policy, error) {
	return s.repo.Get(name)
}

func (s *Service) ListPolicies() ([]*entity.Policy, error) {
	return s.repo.List()
}

func (s *Service) UpdatePolicy(policy *entity.Policy) error {
	err := policy.ValidatePolicy()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(policy)
}

func (s *Service) AddClientToPolicy(client, policy string) error {
	return s.repo.AddClient(client, policy)
}

func (s *Service) RemoveClientFromPolicy(client, policy string) error {
	p, err := s.GetPolicy(policy)
	if p == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.RemoveClient(client, policy)
}

func (s *Service) DeletePolicy(name string) error {
	p, err := s.GetPolicy(name)
	if p == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}

	for i := range p.Clients {
		err = s.repo.RemoveClient(p.Policyname, p.Clients[i])
		if err != nil {
			return entity.ErrClientCannotBeDeleted
		}
	}
	return s.repo.Delete(name)
}
