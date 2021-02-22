package entity

type Client struct {
	Clientname    string
	Policies      []string
	Directorytree string
	Ignorepath    string
	Backups       string
}

func NewClient(clientname string) (*Client, error) {
	c := &Client{
		Clientname: clientname,
	}

	err := c.ValidateClient()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return c, nil

}

func (c *Client) ValidateClient() error {
	if c.Clientname == "" {
		return ErrInvalidEntity
	}

	return nil
}

func (c *Client) AddPolicy(policy string) error {
	_, err := c.GetPolicy(policy)
	if err == nil {
		return ErrPolicyAlreadyAdded
	}
	c.Policies = append(c.Policies, policy)
	return nil
}

func (c *Client) RemovePolicy(policy string) error {
	for i, j := range c.Policies {
		if j == policy {
			c.Policies = append(c.Policies[:i], c.Policies[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (c *Client) GetPolicy(policy string) (string, error) {
	for _, v := range c.Policies {
		if v == policy {
			return policy, nil
		}
	}
	return policy, ErrNotFound
}
