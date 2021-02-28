package policy_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
	"github.com/wade-sam/fyp-backup-server/usecase/policy"
	//"github.com/wade-sam/fyp-backup-server/usecase/policy/mock"
)

func newFixturePolicy() *entity.Policy {
	return &entity.Policy{
		PolicyID:   entity.NewID(),
		Policyname: "wednesday backup",
		Clients:    []entity.ID{entity.NewID()},
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
	repo := policy.NewPolicyHolder()
	testService := policy.NewService(repo)
	p := newFixturePolicy()
	_, err := testService.CreatePolicy(p.Policyname, p.Type, p.Retention, p.Fullbackup, p.IncBackup, p.Clients)
	assert.Nil(t, err)

}
func Test_GetPolicy(t *testing.T) {
	repo := policy.NewPolicyHolder()
	testService := policy.NewService(repo)
	p := newFixturePolicy()
	id, _ := testService.CreatePolicy(p.Policyname, p.Type, p.Retention, p.Fullbackup, p.IncBackup, p.Clients)
	result, err := testService.GetPolicy(id)
	assert.Nil(t, err)
	assert.Equal(t, p.Policyname, result.Policyname)
	result, err = testService.GetPolicy(entity.NewID())
	assert.Nil(t, result)
	assert.Equal(t, err, entity.ErrNotFound)

}

func Test_UpdatePolicy(t *testing.T) {
	repo := policy.NewPolicyHolder()
	testService := policy.NewService(repo)
	p := newFixturePolicy()
	id, _ := testService.CreatePolicy(p.Policyname, p.Type, p.Retention, p.Fullbackup, p.IncBackup, p.Clients)
	get, err := testService.GetPolicy(id)
	get.Policyname = "Thursday's Backup"
	get.Clients = append(p.Clients, entity.NewID())
	err = testService.UpdatePolicy(get)
	assert.Nil(t, err)
	updated, err := testService.GetPolicy(id)
	assert.Nil(t, err)
	assert.Equal(t, "Thursday's Backup", updated.Policyname)
	assert.NotNil(t, get.Clients, updated.Clients)
}

func Test_DeletePolicy(t *testing.T) {
	repo := policy.NewPolicyHolder()
	testService := policy.NewService(repo)
	p1 := newFixturePolicy()
	p2 := newFixturePolicy()
	p1id, _ := testService.CreatePolicy(p1.Policyname, p1.Type, p1.Retention, p1.Fullbackup, p1.IncBackup, p1.Clients)
	err := testService.DeletePolicy(p2.PolicyID)
	assert.Equal(t, entity.ErrNotFound, err)
	err = testService.DeletePolicy(p1id)
	assert.Nil(t, err)
	_, err = testService.GetPolicy(p1id)
	assert.Equal(t, entity.ErrNotFound, err)

}

func Test_ListClients(t *testing.T) {
	repo := policy.NewPolicyHolder()
	testService := policy.NewService(repo)
	p1 := newFixturePolicy()
	p2 := newFixturePolicy()
	plist, err := testService.ListPolicies()
	assert.Equal(t, entity.ErrNotFound, err)
	_, _ = testService.CreatePolicy(p1.Policyname, p1.Type, p1.Retention, p1.Fullbackup, p1.IncBackup, p1.Clients)
	plist, err = testService.ListPolicies()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(plist))
	_, _ = testService.CreatePolicy(p2.Policyname, p2.Type, p2.Retention, p2.Fullbackup, p2.IncBackup, p2.Clients)
	plist2, err := testService.ListPolicies()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(plist2))

}
