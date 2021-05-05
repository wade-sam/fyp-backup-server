package dispatcher

import (
	"fmt"
	"log"
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
			client := ""
			mapstructure.Decode(msg.Data, &client)

			err := s.bus.Unsubscribe("newclient", chn)
			if err != nil {
				return "", entity.ErrNoMatchingTopic
			}
			log.Println("New Client Found:", client)
			return client, nil
		default:
			time.Sleep(2 * time.Second)
		}
	}
	err = s.bus.Unsubscribe("newclient", chn)
	if err != nil {
		return "", entity.ErrNoMatchingTopic
	}
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
