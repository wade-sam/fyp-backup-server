package dispatcher

import (
	"github.com/wade-sam/fyp-backup-server/rabbitBus"
)

type Rabbit interface {
	SearchForNewClient() error
	DirectoryScan(client string) error //Return directory scan struct
	// RemovePolicy(client string, policy string) error
	// AddPolicy(client string, policy string) error
	// RemovePolicies(policy string) error
}

type Bus interface {
	Subscribe(topic string) (rabbitBus.EventChannel, error)
	Unsubscribe(topic string, ch chan rabbitBus.Event) error
}

type Repository interface {
	Rabbit
	Bus
}

type UseCase interface {
	SearchForNewClient() (string, error)
	GetDirectoryScan(client string) //Return Directory scan struct
	// AddPolicyToClient(consumerID, policyID []string) error
	// RemovePolicyFromClient(client string, policy []string) error
	// RemoveClientFromPolicy(client []string, policy string) error
}
