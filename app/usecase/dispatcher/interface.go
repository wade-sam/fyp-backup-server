package dispatcher

import (
	"github.com/wade-sam/fyp-backup-server/entity"
)

type Repository interface {
	NewClient() (entity.ID, error)
	RemovePolicy(client entity.ID, policy entity.ID) error
	AddPolicy(client entity.ID, policy entity.ID) error
	RemovePolicies(policy entity.ID) error
}

type UseCase interface {
	SearchNewClient() (entity.ID, error)
	AddPolicyToClient(consumerID, policyID []entity.ID) error
	RemovePolicyFromClient(client entity.ID, policy []entity.ID) error
	RemoveClientFromPolicy(client []entity.ID, policy entity.ID) error
}
