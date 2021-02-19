package query

import "github.com/wade-sam/fyp-backup-server/pkg/Entities"

//Query Interface talks directly to the Database and is used for GUI. Consists of
//business logic
type QueryInterface interface {
	ViewClient(client *Entities.Client) (*Entities.Client, error)
	ViewClients(clients []*Entities.Client) ([]*Entities.Client, error)
	ViewPolicy(policy *Entities.Policy) (*Entities.Policy, error)
	ViewPolicies(policies []*Entities.Policy) ([]*Entities.Policy, error)
}
