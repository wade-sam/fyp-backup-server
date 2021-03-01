package entity

type Client struct {
	ID string `bson:"_id omitempty"`
	ConsumerID    string    `bson:"consumerid"`
	Clientname    string   `bson:"clientname"`
	Policies      []string     `bson:"policies"`
	Directorytree []string `bson: "treepath"`
	Ignorepath    []string `bson: "ignore"`
	Backups       []string `bson: "backups"`
}

func NewClient(clientname, consumerid string) (*Client, error) {
	c := &Client{
		ConsumerID: consumerid,
		Clientname: clientname,
	}

	err := c.ValidateClient()
	if err != nil {
		return nil, ErrInvalidEntity
	}

	return c, nil
}

func (c *Client) ValidateClient() error {
	if c.Clientname == "" || c.ConsumerID == "" {
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
