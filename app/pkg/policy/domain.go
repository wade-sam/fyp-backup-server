package policy

import (
	"errors"

	"github.com/go-playground/validator/v10"
	errs "github.com/pkg/errors"
	"github.com/wade-sam/fyp-backup-server/pkg/Entities"
)

var (
	ErrPolicyInvalid  = errors.New("Policy Invalid")
	ErrPolicyNotFound = errors.New("Policy Not Found")
)

type policyService struct {
	policyRepo PolicyRepository
}

var validate = validator.New()

func NewPolicyService(policyRepo PolicyRepository) PolicyService {
	return &policyService{
		policyRepo,
	}
}

func (p *policyService) FindPolicy(name string) (*Entities.Policy, error) {
	result, err := p.policyRepo.FindPolicy(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *policyService) CreatePolicy(policy *Entities.Policy) error {
	if err := validate.Struct(policy); err != nil {
		return errs.Wrap(ErrPolicyInvalid, "service.Policy.Create")
	}
	return p.policyRepo.CreatePolicy(policy)
}
func (p *policyService) UpdatePolicy(name string, policy *Entities.Policy) error {
	return p.policyRepo.UpdatePolicy(name, policy)
}

func (p *policyService) DeletePolicy(name string) error {
	err := p.policyRepo.DeletePolicy(name)
	if err != nil {
		return err
	}
	return nil
}
