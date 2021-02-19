package client

import "github.com/wade-sam/fyp-backup-server/pkg/Entities"

type ClientService interface {
	FindClient(name string) (*Entities.Client, error)
	CreateClient(client *Entities.Client) error
	UpdateClient(name string, client *Entities.Client) error
	DeleteClient(name string) error
	ConnectClient() error
}
