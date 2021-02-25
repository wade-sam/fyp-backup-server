package policy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
	"github.com/wade-sam/fyp-backup-server/usecase/policy"
	"github.com/wade-sam/fyp-backup-server/usecase/policy/mock"
)

func newFixturePolicy() *entity.Policy {
	return &entity.Policy{
		Policyname: "wednesday backup",
		Clients:    []string{"sam"},
		Retention:  10,
		Type:       "full",
		State:      "active",
		Fullbackup: []string{"Monday", "Thursday", "Sunday"},
		IncBackup:  []string{},
	}
}
func listPolicies() []*entity.Policy {
	lists := []*entity.Policy{}
	for i := 0; i < 5; i++ {
		client := newFixturePolicy()
		lists = append(lists, client)
	}
	return lists
}
func Test_CreatePolicy(t *testing.T) {
	p := newFixturePolicy()
	mockRepo := new(mock.MockRepository)
	mockRepo.On("Create", p).Return(p, nil)
	testService := policy.NewService(mockRepo)
	result, err := testService.CreatePolicy(p.Policyname, p.Type, p.Retention, p.Fullbackup, p.IncBackup, p.Clients)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, result, p)

	mockRepo = new(mock.MockRepository)
	mockRepo.On("Create", p).Return(p, nil)
	testService = policy.NewService(mockRepo)
	result, err = testService.CreatePolicy("", p.Type, p.Retention, p.Fullbackup, p.IncBackup, p.Clients)
	assert.Nil(t, result)
	assert.Equal(t, err, entity.ErrInvalidEntity)
}

func Test_GetPolicy(t *testing.T) {
	p := newFixturePolicy()
	mockRepo := new(mock.MockRepository)
	mockRepo.On("Get", p.Policyname).Return(p, nil)
	testService := policy.NewService(mockRepo)
	result, err := testService.GetPolicy(p.Policyname)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, result, p)

	mockRepo = new(mock.MockRepository)
	mockRepo.On("Get", p.Policyname).Return(nil, entity.ErrNotFound)
	testService = policy.NewService(mockRepo)
	result, err = testService.GetPolicy(p.Policyname)
	assert.Nil(t, result)
	assert.Equal(t, err, entity.ErrNotFound)
}

func Test_UpdatePolicy(t *testing.T) {
	p := newFixturePolicy()
	d := entity.Policy{
		Clients:    []string{"sam"},
		Retention:  10,
		Type:       "full",
		State:      "active",
		Fullbackup: []string{"Monday", "Thursday", "Sunday"},
		IncBackup:  []string{},
	}
	mockRepo := new(mock.MockRepository)
	mockRepo.On("Update", p).Return(nil)
	testService := policy.NewService(mockRepo)
	err := testService.UpdatePolicy(p)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	err = testService.UpdatePolicy(&d)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, entity.ErrInvalidEntity, err)
}
func Test_ListPolicy(t *testing.T) {
	//p := newFixturePolicy()
	policies := listPolicies()
	mockRepo := new(mock.MockRepository)
	mockRepo.On("List").Return(policies, nil)
	testService := policy.NewService(mockRepo)
	result, err := testService.ListPolicies()
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, result, policies)
	mockRepo2 := new(mock.MockRepository)
	mockRepo2.On("List").Return(nil, entity.ErrNotFound)
	testService = policy.NewService(mockRepo2)
	result, err = testService.ListPolicies()
	assert.Equal(t, entity.ErrNotFound, err)
	assert.Nil(t, result)
}

func Test_DeletePolicy(t *testing.T) {
	p := newFixturePolicy()
	mockRepo := new(mock.MockRepository)
	mockRepo.On("Delete", p.Policyname).Return(nil)
	mockRepo.On("Get", p.Policyname).Return(p, nil)
	testService := policy.NewService(mockRepo)
	err := testService.DeletePolicy(p.Policyname)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	mockRepo2 := new(mock.MockRepository)
	mockRepo2.On("Get", p.Policyname).Return(p, entity.ErrNotFound)
	testService = policy.NewService(mockRepo2)
	err = testService.DeletePolicy(p.Policyname)
	assert.Equal(t, entity.ErrNotFound, err)

}
