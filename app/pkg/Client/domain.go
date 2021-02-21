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

type internalClientHandler struct {
	clientRepo  ClientRepository
	clientsRepo ClientManageRepository
}

type externalClientHandler struct {
	clientRepo  ClientRepository
	clientsRepo ClientManageRepository
}

var validate = validator.New()

func NewClientService(clientRepo ClientRepository, clientMRepo ClientManageRepository) (InternalClientService, ExternalClientService) {
	return &internalClientHandler{
			clientRepo,
			clientMRepo,
		}, &externalClientHandler{
			clientRepo,
			clientMRepo,
		}
}

func (c *externalClientHandler) DirectoryScanResult(directory []string) {}

func (c *externalClientHandler) ConfigRequest(client string) {}

func (c *externalClientHandler) NewClient(config string) {}

func (c *internalClientHandler) FindClient(name string) (*Entities.Client, error) {
	return c.clientRepo.FindClient(name)
}

func (c *internalClientHandler) CreateClient(client *Entities.Client) error {
	if err := validate.Struct(client); err != nil {
		return errs.Wrap(ErrClientInvalid, "service.Client.Create")
	}
	return c.clientRepo.CreateClient(client)
}

func (p *internalClientHandler) UpdateClient(name string, client *Entities.Client) error {
	return p.clientRepo.UpdateClient(name, client)
}

func (p *internalClientHandler) DeleteClient(name string) error {
	err := p.clientRepo.DeleteClient(name)
	if err != nil {
		return err
	}
	return nil
}

func (p *internalClientHandler) ConnectClient() error {
	c := make(chan int)
	p.clientsRepo.ConnectClient(c)
	return nil
	//return p.ConnectClient()
}
