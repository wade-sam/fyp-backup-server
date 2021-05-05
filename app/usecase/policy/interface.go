package policy

import (
	"github.com/wade-sam/fyp-backup-server/entity"
)

type Reader interface {
	Get(name string) (*entity.Policy, error)
	GetName(name string) (string, error)
	List() ([]*entity.Policy, error)
}

type Writer interface {
	Create(policy *entity.Policy) (string, error)
	Update(policy *entity.Policy) error
	Delete(name string) error
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
}

type UseCase interface {
	GetPolicy(name string) (*entity.Policy, error)
	GetPolicyName(id string) (string, error)
	ListPolicies() ([]*entity.Policy, error)
	CreatePolicy(policyname, runtime, backupType string, retention int, fullbackup, incrementalbackup []string, clients []string) (string, error)
	UpdatePolicy(policy *entity.Policy) error
	DeletePolicy(name string) error
}
