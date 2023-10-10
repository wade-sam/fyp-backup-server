package client

import "github.com/wade-sam/fyp-backup-server/entity"

type Service struct {
	repo Repository
}

func NewService(p Repository) *Service {
	return &Service{
		repo: p,
	}
}

//TODO: re-check how we want client created. Just a name? or is a filescan required. Should I add a policy? If so filescan should be required first

func (s *Service) CreateClient(clientname string, consumerID string) (string, error) {
	client, err := entity.NewClient(clientname, consumerID)
	err = client.ValidateClient()
	if err != nil {
		return "", entity.ErrInvalidEntity
	}
	return s.repo.Create(client)
}

func (s *Service) GetClient(name string) (*entity.Client, error) {
	return s.repo.Get(name)
}

func (s *Service) GetClientName(id string) (string, error) {
	return s.repo.GetName(id)
}

func (s *Service) ListClients() ([]*entity.Client, error) {
	result, err := s.repo.List()
	if result == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return s.repo.List()
}

func (s *Service) UpdateClient(client *entity.Client) error {
	err := client.ValidateClient()
	if err != nil {
		return entity.ErrInvalidEntity
	}
	// r, err := s.GetClient(client.ID)
	// if r == nil {
	// 	return entity.ErrNotFound
	// }
	// if err != nil {
	// 	return err
	// }
	err = s.repo.Update(client)
	if err != nil {
		return err
	}
	return nil
}

/*
	   To check for differences in the lists you have to compare both ways. r1 contains the policies to remove
	   from the client and r2 contains the policies to add to the client

	policiesRemove, _ := ComparePolicyLists(client.Policies, r.Policies)
	policiesAdd, _ := ComparePolicyLists(r.Policies, client.Policies)
	if len(policiesRemove) > 0 {
		err = s.dispatcher.RemovePolicyFromClient(client.ConsumerID, policiesRemove)
		if err != nil {
			return err
		}
	}
	if len(policiesAdd) > 0 {
		err = s.dispatcher.AddPolicyToClient(client.ConsumerID, policiesAdd)
		if err != nil {
			return err
		}
	}
	err = s.persistence.Update(client)
	if err != nil {
		return err
	}
	return nil


//Compares the two lists for differences in policies and adds them to a differences list which is returned
func ComparePolicyLists(l1, l2 []string) ([]string, error) {
	differences := []string{}
	for _, i := range l1 {
		for _, j := range l2 {
			if i == j {
				break
			}
		}
		differences = append(differences, i)
	}
	if len(differences) == 0 {
		return nil, entity.ErrNoNewItem
	}
	return differences, nil
}
*/

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
