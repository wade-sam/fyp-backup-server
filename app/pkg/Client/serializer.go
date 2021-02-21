package client

import "github.com/wade-sam/fyp-backup-server/pkg/Entities"

type ClientSerializer interface {
	DecodeClient(input []byte) (*Entities.Client, error)
	EncodeClient(input *Entities.Client) ([]byte, error)
	EncodeDirectoryScan(input *Entities.FileScan) ([]byte, error)
	DecodeDirectoryScan(input []byte) (*Entities.FileScan, error)
}
