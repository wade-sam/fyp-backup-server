package dispatcher

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/wade-sam/fyp-backup-server/entity"
)

type Service struct {
	rabbit Rabbit
	bus    Bus
}

func NewService(r Rabbit, b Bus) *Service {
	return &Service{
		rabbit: r,
		bus:    b,
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
	chn, err := s.bus.Subscribe("newclient")
	if err != nil {
		return "", err
	}
	err = s.rabbit.SearchForNewClient()
	if err != nil {
		return "", entity.ErrNoNewClient
	}
	for i := 1; i < 10; i++ {
		select {
		case msg := <-chn:
			s := ""
			mapstructure.Decode(msg.Data, &s)
			fmt.Println("mastructure", s)
			close(chn)
			return s, nil
		default:
			time.Sleep(2 * time.Second)
		}
	}
	close(chn)
	fmt.Println("NO NEW CLIENT")
	return "", entity.ErrNoNewClient
}
func (s *Service) GetDirectoryScan(client string) (*entity.Directory, error) {
	chn, err := s.bus.Subscribe("directoryscan")
	if err != nil {
		return nil, err
	}

	err = s.rabbit.DirectoryScan(client)
	if err != nil {
		return nil, err
	}
	for i := 1; i < 200; i++ {
		select {
		case msg := <-chn:
			d := entity.Directory{}
			mapstructure.Decode(msg.Data, &d)
			err = s.bus.Unsubscribe("directoryscan", chn)
			if err != nil {
				close(chn)
				return &d, err
			}
			return &d, nil

		default:
			time.Sleep(2 * time.Second)
		}
	}
	close(chn)
	return nil, entity.ErrNotFound

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
