package device_communication

import "github.com/wade-sam/fyp-backup-server/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) AddPolicyToClient(client, policy string) error {
	return s.repo.AddPolicyToClient(client, policy)
}

func (s *Service) RemovePolicyFromClient(client, policy string) error {
	return s.repo.RemovePolicyFromClient(client, policy)
}

func (s *Service) RemoveClientFromPolicy(policy *entity.Policy) error {
	for _, j := range policy.Clients {
		err := s.repo.RemoveClientFromPolicy(j, policy.Policyname)
		if err != nil {
			return entity.ErrPolicyCantBeRemoved
		}
	}
	return nil
}
func (s *Service) SearchNewClient() (string, error) {
	name, err := s.repo.SearchNewClient()
	if err != nil {
		return "", entity.ErrNoNewClient
	}
	return name, nil
}
