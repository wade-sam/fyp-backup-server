package ports

import (
	"github.com/wade-sam/fyp-backup-server/pkg/client"
	"github.com/wade-sam/fyp-backup-server/pkg/policy"
)

type PersistentRepository interface {
	client.ClientRepository
	policy.PolicyRepository
}

type BrokerRepository interface {
	client.ClientManageRepository
}
