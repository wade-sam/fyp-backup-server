package client

import (
	"errors"

	//"github.com/go-playground/validator"
	"github.com/go-playground/validator/v10"
	errs "github.com/pkg/errors"
)

type Client struct {
	Clientname   string `json:"clientname" bson:"clientname" validate:"required"`
	Consumername string `json:"consumername" bson:"consumername" validate:"required"`
	Policies     string `json:"policies" bson:"policies"`
	Backuptree   string `json:"backuptree" bson:"backuptree"`
	Ignorepath   string `json:"ignorepath" bson:"ignorepath"`
	Backups      string `json:"backups" bson:"backups"`
}

var (
	ErrClientInvalid  = errors.New("Client Invalid")
	ErrClientNotFound = errors.New("Client Not Found")
)

type clientService struct {
	clientRepo ClientRepository
}

var validate = validator.New()

func NewClientService(clientRepo ClientRepository) ClientService {
	return &clientService{
		clientRepo,
	}
}

func (c *clientService) FindClient(name string) (*Client, error) {
	return c.clientRepo.FindClient(name)
}

func (c *clientService) CreateClient(client *Client) error {
	if err := validate.Struct(client); err != nil {
		return errs.Wrap(ErrClientInvalid, "service.Client.Create")
	}
	return c.clientRepo.CreateClient(client)
}

func (p *clientService) UpdateClient(name string, client *Client) error {
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
	return p.ConnectClient()
}
