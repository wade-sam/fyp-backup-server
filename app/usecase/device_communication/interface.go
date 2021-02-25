package device_communication

type Repository interface {
	SearchNewClient() (string, error)
	RemovePolicyFromClient(client, policy string) error
	AddPolicyToClient(client, policy string) error
	RemoveClientFromPolicy(client, policy string) error
}

type UseCase interface {
	SearchNewClient() (string, error)
	AddPolicyToClient(client, policy string) error
	RemovePolicyFromClient(client, policy string) error
	RemoveClientsFromPolicy(client, policy string) error
}
