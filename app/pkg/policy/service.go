package policy

import "github.com/wade-sam/fyp-backup-server/pkg/Entities"

type PolicyService interface {
	FindPolicy(name string) (*Entities.Policy, error)
	CreatePolicy(policy *Entities.Policy) error
	UpdatePolicy(name string, policy *Entities.Policy) error
	DeletePolicy(name string) error
}
