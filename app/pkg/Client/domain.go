package client

import (
	"errors"

	//"github.com/go-playground/validator"
	"github.com/go-playground/validator/v10"
	errs "github.com/pkg/errors"
	"github.com/wade-sam/fyp-backup-server/pkg/Entities"
)

var (
	ErrClientInvalid  = errors.New("Client Invalid")
	ErrClientNotFound = errors.New("Client Not Found")
)

type clientService struct {
	clientRepo  ClientRepository
	clientsRepo ClientManageRepository
}

var validate = validator.New()

func NewClientService(clientRepo ClientRepository, clientMRepo ClientManageRepository) ClientService {
	return &clientService{
		clientRepo,
		clientMRepo,
	}
}

func (c *clientService) FindClient(name string) (*Entities.Client, error) {
	return c.clientRepo.FindClient(name)
}

func (c *clientService) CreateClient(client *Entities.Client) error {
	if err := validate.Struct(client); err != nil {
		return errs.Wrap(ErrClientInvalid, "service.Client.Create")
	}
	return c.clientRepo.CreateClient(client)
}

func (p *clientService) UpdateClient(name string, client *Entities.Client) error {
	return p.clientRepo.UpdateClient(name, client)
}

func (p *clientService) DeleteClient(name string) error {
	err := p.clientRepo.DeleteClient(name)
	if err != nil {
		return err
	}
	return nil
}

func (p *clientService) ConnectClient() error {
	c := make(chan int)
	p.clientsRepo.ConnectClient(c)
	return nil
	//return p.ConnectClient()
}
