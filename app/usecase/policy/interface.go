package policy

import (
	"github.com/wade-sam/fyp-backup-server/entity"
)

type Reader interface {
	Get(name entity.ID) (*entity.Policy, error)
	List() ([]*entity.Policy, error)
}

type Writer interface {
	Create(policy *entity.Policy) (entity.ID, error)
	Update(policy *entity.Policy) error
	Delete(name entity.ID) error
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
	GetPolicy(name entity.ID) (*entity.Policy, error)
	ListPolicies() ([]*entity.Policy, error)
	CreatePolicy(policyname, backupType string, retention int, fullbackup, incrementalbackup, clients []entity.ID) (entity.ID, error)
	UpdatePolicy(policy *entity.Policy) error
	DeletePolicy(name entity.ID) error
}
