package client

type ClientService interface {
	FindClient(name string) (*Client, error)
	CreateClient(client *Client) error
	UpdateClient(name string, client *Client) error
	DeleteClient(name string) error
	ConnectClient() error
}
