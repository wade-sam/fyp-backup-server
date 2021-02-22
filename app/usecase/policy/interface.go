package policy

import (
	"github.com/wade-sam/fyp-backup-server/entity"
)

type Reader interface {
	Get(name string) (*entity.Policy, error)
	List() ([]*entity.Policy, error)
}

type Writer interface {
	Create(policy *entity.Policy) error
	Update(policy *entity.Policy) error
	Delete(name string) error
}

type Communication interface {
	AddClient(client, policy string) error
	RemoveClient(client, policy string) error
}

/* TODO
TODO Choices to make:
 1)How to Implement the removal of clients?
	- Put the logic in here for removing clients from database? or in repo?
	- Should I add another interface for updating and removing policies from devices or leave it for the client?

*/
type Repository interface {
	Reader
	Writer
	Communication
}

type UseCase interface {
	GetPolicy(name string) (*entity.Policy, error)
	ListPolicies(name string) ([]*entity.Policy, error)
	CreatePolicy(policyname, backupType string, retention int, fullbackup, incrementalbackup, clients []string) error
	UpdatePolicy(policy *entity.Policy) error
	AddClientToPolicy(client, policy string) error
	RemoveClientFromPolicy(client, policy string) error
	DeletePolicy(name string) error
}
