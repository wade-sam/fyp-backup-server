package ports

import (
	client "github.com/wade-sam/fyp-backup-server/pkg/Client"
	"github.com/wade-sam/fyp-backup-server/pkg/policy"
)

type PersistentRepository interface {
	client.ClientRepository
	policy.PolicyRepository
}

type BrokerRepository interface {
	client.ClientManageRepository
}

type ClientCommunication interface {
	client.ExternalClientService
}
