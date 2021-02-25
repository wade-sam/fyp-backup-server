package client

import "github.com/wade-sam/fyp-backup-server/entity"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//TODO: re-check how we want client created. Just a name? or is a filescan required. Should I add a policy? If so filescan should be required first

func (s *Service) CreateClient(client *entity.Client) (*entity.Client, error) {
	//client, err := entity.NewClient(client)
	err := client.ValidateClient()
	if err != nil {
		return nil, entity.ErrInvalidEntity
	}
	return s.repo.Create(client)
}

func (s *Service) GetClient(name string) (*entity.Client, error) {
	return s.repo.Get(name)
}

func (s *Service) ListClients() ([]*entity.Client, error) {
	return s.repo.List()
}

func (s *Service) UpdateClient(client *entity.Client) error {
	err := client.ValidateClient()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(client)
}

func (s *Service) DeleteClient(name string) error {
	c, err := s.GetClient(name)
	if c == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	return s.repo.Delete(name)
}
