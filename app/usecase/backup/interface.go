package backup

import (
	"github.com/wade-sam/fyp-backup-server/entity"
	"github.com/wade-sam/fyp-backup-server/rabbitBus"
)

type PolicyRepository interface {
	Get(name string) (*entity.Policy, error)
	Update(policy *entity.Policy) error
}

type ClientRepository interface {
	Get(name string) (*entity.Client, error)
}

type BackupRepository interface {
	Create(clientrun *entity.ClientRun) (string, error)
}
type RabbitRepository interface {
	StartStorageNode(clients []string, storagenode, policy string) error
	StartBackup(clients, policyID, clientname, backuptype string, ignorelist []string) error
	//CancelBackup(clients []*entity.Client)
}

type BusRepository interface {
	Subscribe(topic string) (rabbitBus.EventChannel, error)
	Unsubscribe(topic string, ch chan rabbitBus.Event) error
}
type UseCase interface {
	StartBackup(policy, Type string) error
	StartIncrementalBackup(policy string) error
}