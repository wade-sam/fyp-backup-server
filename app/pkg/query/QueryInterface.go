package query

import "github.com/wade-sam/fyp-backup-server/pkg/Entities"

//Query Interface queries the database for data. These can be for the presentation of data on the UI or for the queryingservice for backup runs
type QueryInterface interface {
	ViewClient(client *Entities.Client) (*Entities.Client, error)
	ViewClients(clients []*Entities.Client) ([]*Entities.Client, error)
	ViewPolicy(policy *Entities.Policy) (*Entities.Policy, error)
	ViewPolicies(policies []*Entities.Policy) ([]*Entities.Policy, error)
}
