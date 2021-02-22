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
func (s *Service) SearchNewClient() (string, error) {
	name, err := s.repo.SearchNewClient()
	if err != nil {
		return "", entity.ErrNoNewClient
	}
	return name, nil
}

//TODO: re-check how we want client created. Just a name? or is a filescan required. Should I add a policy? If so filescan should be required first

func (s *Service) CreateClient(name string) error {
	client, err := entity.NewClient(name)
	if err != nil {
		return err
	}
	return s.repo.Create(client)
}

func (s *Service) GetClient(name string) (*entity.Client, error) {
	return s.GetClient(name)
}

func (s *Service) ListClients() ([]*entity.Client, error) {
	return s.ListClients()
}

func (s *Service) UpdateClient(client *entity.Client) error {
	err := client.ValidateClient()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	return s.repo.Update(client)
}

func (s *Service) AddPolicyToClient(client, policy string) error {
	return s.repo.AddPolicy(client, policy)
}

func (s *Service) RemovePolicyFromClient(client, policy string) error {
	c, err := s.GetClient(client)
	if err == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	_, err = c.GetPolicy(policy)
	if err != nil {
		return err
	}
	return s.repo.RemovePolicy(client, policy)
}

func Find(policies []string, value string) bool {
	for _, item := range policies {
		if item == value {
			return true
		}
	}
	return false
}

func (s *Service) DeleteClient(name string) error {
	c, err := s.GetClient(name)
	if c == nil {
		return entity.ErrNotFound
	}
	if err != nil {
		return err
	}
	for i := range c.Policies {
		err = s.repo.RemovePolicy(c.Clientname, c.Policies[i])
		if err != nil {
			return entity.ErrClientCannotBeDeleted
		}
	}
	return s.repo.Delete(name)
}
