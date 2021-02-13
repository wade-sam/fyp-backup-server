package client

type ClientSerializer interface {
	Decode(input []byte) (*Client, error)
	Encode(input *Client) ([]byte, error)
}