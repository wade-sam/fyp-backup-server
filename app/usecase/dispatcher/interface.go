package dispatcher

import "github.com/wade-sam/fyp-backup-server/entity"

type Repository interface {
	SearchForNewClient() (string, error)
	DirectoryScan(client string) (*entity.Directory, error) //Return directory scan struct
	// RemovePolicy(client string, policy string) error
	// AddPolicy(client string, policy string) error
	// RemovePolicies(policy string) error
}

type UseCase interface {
	SearchForNewClient() (string, error)
	GetDirectoryScan(client string) //Return Directory scan struct
	// AddPolicyToClient(consumerID, policyID []string) error
	// RemovePolicyFromClient(client string, policy []string) error
	// RemoveClientFromPolicy(client []string, policy string) error
}
