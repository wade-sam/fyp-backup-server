package entity_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
)

func TestNewClient(t *testing.T) {
	id := uuid.NewV4()
	u, err := entity.NewClient("Sam's MacBook Pro", id)
	assert.Nil(t, err)
	assert.Equal(t, u.Clientname, "Sam's MacBook Pro")
}

func TestAddPolicy(t *testing.T) {
	id := uuid.NewV4()
	u, _ := entity.NewClient("Sam's MacBook Pro", id)
	//policy := "Wednesday Backup Run"
	policyid := uuid.NewV4()
	err := u.AddPolicy(policyid)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(u.Policies))
	err = u.AddPolicy(policyid)
	assert.Equal(t, entity.ErrPolicyAlreadyAdded, err)
}

func TestRemovePolicy(t *testing.T) {
	id := uuid.NewV4()
	u, _ := entity.NewClient("Sam's MacBook Pro", id)
	policyid := uuid.NewV4()
	err := u.RemovePolicy(policyid)
	assert.Equal(t, entity.ErrNotFound, err)
	//policy := "Wednesday Backup Run"
	_ = u.AddPolicy(policyid)
	err = u.RemovePolicy(policyid)
	assert.Nil(t, err)
}

func TestGetPolicy(t *testing.T) {
	id := uuid.NewV4()
	u, _ := entity.NewClient("Sam's MacBook Pro", id)
	//bpolicy := "Wednesday Backup Run"
	policyid := uuid.NewV4()
	_ = u.AddPolicy(policyid)
	policy, err := u.GetPolicy(policyid)
	assert.Nil(t, err)
	assert.Equal(t, policy, policyid)
	policyid2 := uuid.NewV4()
	_, err = u.GetPolicy(policyid2)
	assert.Equal(t, entity.ErrNotFound, err)

}

func TestClientValidate(t *testing.T) {
	type test struct {
		clientname string
		clientid   entity.ID
		want       error
	}

	tests := []test{
		{
			clientname: "sam wade",
			clientid:   uuid.NewV4(),
			want:       nil,
		},
		{
			clientname: "",
			clientid:   uuid.NewV4(),
			want:       entity.ErrInvalidEntity,
		},
		{
			clientname: "",
			clientid:   uuid.NewV4(),
			want:       entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewClient(tc.clientname, tc.clientid)
		assert.Equal(t, err, tc.want)
	}
}
