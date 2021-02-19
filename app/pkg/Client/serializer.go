package client

import "github.com/wade-sam/fyp-backup-server/pkg/Entities"

type ClientSerializer interface {
	Decode(input []byte) (*Entities.Client, error)
	Encode(input *Entities.Client) ([]byte, error)
}
