package client

import "github.com/wade-sam/fyp-backup-server/pkg/Entities"

type InternalClientService interface {
	//FindClient(name string) (*Entities.Client, error)
	AddClient() error
	ConfigureClientBackupList(client string, directoryscan *Entities.DirectoryScan) error
	ClientDirectoryScan(client, path string) error
	DeleteClient(name string) error
	AddClientPolicy(client, policy string) error
	DeleteClientPolicy(client, policy string) error
	//CreateClient(client *Entities.Client) error
	//UpdateClient(name string, client *Entities.Client) error

	ConnectClient() error
}

type ExternalClientService interface {
	DirectoryScanResult(directory []string)
	ConfigRequest(client string)
	NewClient(config string)
}

type ClientService interface {
	ExternalClientService
	InternalClientService
}
