package client

import (
	"github.com/wade-sam/fyp-backup-server/entity"
)

type Reader interface {
	Get(name string) (*entity.Client, error)
	List() ([]entity.Client, error)
}

type Writer interface {
	Create(client *entity.Client) error
	Update(client *entity.Client) error
	Delete(name string) error
}

/* TODO
TODO Choices to make:
 1)How to Implement the removal of clients?
	- Put the logic in here for removing clients from database? or in repo?
	- Should i put the logichere for working out which policy/'s need to be added or removed or leave that for the repo
*/

type Communication interface {
	SearchNewClient() (string, error)
	RemovePolicy(client, policy string) error
	AddPolicy(client, policy string) error
}

type Repository interface {
	Reader
	Writer
	Communication
}

type UseCase interface {
	GetClient(name string) (*entity.Client, error)
	ListClients() ([]*entity.Client, error)
	SearchNewClient() (*entity.Client, error)
	AddPolicyToClient(client, policy string) error
	RemovePolicyFromClient(client, policy string) error
	CreateClient(client *entity.Client) error
	UpdateClient(client *entity.Client) error
	DeleteClient(name string) error
}
