package dispatcher

type Repository interface {
	NewClient() (string, error)
	RemovePolicy(client string, policy string) error
	AddPolicy(client string, policy string) error
	RemovePolicies(policy string) error
}

type UseCase interface {
	SearchNewClient() (string, error)
	AddPolicyToClient(consumerID, policyID []string) error
	RemovePolicyFromClient(client string, policy []string) error
	RemoveClientFromPolicy(client []string, policy string) error
}
