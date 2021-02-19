package policy

import "github.com/wade-sam/fyp-backup-server/pkg/Entities"

type PolicySerializer interface {
	Decode(input []byte) (*Entities.Policy, error)
	Encode(input *Entities.Policy) ([]byte, error)
}
