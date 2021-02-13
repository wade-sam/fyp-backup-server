package policy

type PolicySerializer interface {
	Decode(input []byte) (*Policy, error)
	Encode(input *Policy) ([]byte, error)
}