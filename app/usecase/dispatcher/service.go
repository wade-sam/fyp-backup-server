package dispatcher

import "github.com/wade-sam/fyp-backup-server/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) AddPolicyToClient(consumerID entity.ID, policyID []entity.ID) error {
	for _, j := range policyID {
		err := s.repo.AddPolicy(consumerID, j)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) RemovePolicyFromClient(consumerID entity.ID, policyID []entity.ID) error {
	for _, j := range policyID {
		err := s.repo.RemovePolicy(consumerID, j)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) RemoveClientFromPolicy(policyID entity.ID) error {
	return s.repo.RemovePolicies(policyID)

}
func (s *Service) SearchNewClient() (entity.ID, error) {
	id, err := s.repo.NewClient()
	if err != nil {
		return id, entity.ErrNoNewClient
	}
	return id, nil
}
