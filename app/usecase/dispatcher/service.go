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

// func (s *Service) AddPolicyToClient(consumerID string, policyID []string) error {
// 	for _, j := range policyID {
// 		err := s.repo.AddPolicy(consumerID, j)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
// }

func (s *Service) SearchForNewClient() (string, error) {
	id, err := s.repo.SearchForNewClient()
	if err != nil {
		return id, entity.ErrNoNewClient
	}
	return id, nil
}
func (s *Service) GetDirectoryScan(client string) (*entity.Directory, error) {
	scanResult, err := s.repo.DirectoryScan(client)

	if err != nil {
		return nil, err
	}
	return scanResult, nil
}

// func (s *Service) RemovePolicyFromClient(consumerID string, policyID []string) error {
// 	for _, j := range policyID {
// 		err := s.repo.RemovePolicy(consumerID, j)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (s *Service) RemoveClientFromPolicy(policyID string) error {
// 	return s.repo.RemovePolicies(policyID)
