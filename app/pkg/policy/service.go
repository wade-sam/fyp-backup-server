package policy

type PolicyService interface {
	FindPolicy(name string) (*Policy, error)
	CreatePolicy(policy *Policy) error
	UpdatePolicy(name string, policy *Policy) error
	DeletePolicy(name string) error
}
