package client

import (
	"github.com/wade-sam/fyp-backup-server/entity"
)

type Reader interface {
	Get(name entity.ID) (*entity.Client, error)
	List() ([]*entity.Client, error)
}

type Writer interface {
	Create(client *entity.Client) (entity.ID, error)
	Update(client *entity.Client) error
	Delete(name entity.ID) error
}

/* TODO
TODO Choices to make:
 1)How to Implement the removal of clients?
	- Put the logic in here for removing clients from database? or in repo?
	- Should i put the logichere for working out which policy/'s need to be added or removed or leave that for the repo
*/

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetClient(name entity.ID) (*entity.Client, error)
	ListClients() ([]*entity.Client, error)
	CreateClient(client *entity.Client) (entity.ID, error)
	UpdateClient(client *entity.Client) error
	DeleteClient(name entity.ID) error
	SearchNewClient() (entity.ID, error)
}
