package client

type ClientRepository interface {
	FindClient(name string) (*Client, error)
	CreateClient(client *Client) error
	UpdateClient(name string, client *Client) error
	DeleteClient(name string) error
}

type ClientManageRepository interface {
	//ConnectClient() error
}
