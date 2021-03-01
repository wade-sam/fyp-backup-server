package entity_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
)

func TestNewClient(t *testing.T) {
	u, err := entity.NewClient("Sam's MacBook Pro", "client1")
	assert.Nil(t, err)
	assert.Equal(t, u.Clientname, "Sam's MacBook Pro")
}

func TestAddPolicy(t *testing.T) {
	u, _ := entity.NewClient("Sam's MacBook Pro", "client1")
	//policy := "Wednesday Backup Run"
	err := u.AddPolicy("p1")
	assert.Nil(t, err)
	assert.Equal(t, 1, len(u.Policies))
	err = u.AddPolicy("p1")
	assert.Equal(t, entity.ErrPolicyAlreadyAdded, err)
}

func TestRemovePolicy(t *testing.T) {
	u, err := entity.NewClient("Sam's MacBook Pro", "client1")
	err = u.RemovePolicy("p1")
	assert.Equal(t, entity.ErrNotFound, err)
	//policy := "Wednesday Backup Run"
	_ = u.AddPolicy("p1")
	err = u.RemovePolicy("p1")
	assert.Nil(t, err)
}

func TestGetPolicy(t *testing.T) {
	u, err := entity.NewClient("Sam's MacBook Pro", "client1")
	//bpolicy := "Wednesday Backup Run"
	policyid := "p1"
	_ = u.AddPolicy(policyid)
	policy, err := u.GetPolicy(policyid)
	assert.Nil(t, err)
	assert.Equal(t, policy, policyid)
	policyid2 := "p2"
	_, err = u.GetPolicy(policyid2)
	assert.Equal(t, entity.ErrNotFound, err)

}

func TestClientValidate(t *testing.T) {
	type test struct {
		clientname string
		clientid   string
		want       error
	}

	tests := []test{
		{
			clientname: "sam wade",
			clientid:   "host1",
			want:       nil,
		},
		{
			clientname: "",
			clientid:   "host1",
			want:       entity.ErrInvalidEntity,
		},
		{
			clientname: "",
			clientid:   "",
			want:       entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewClient(tc.clientname, tc.clientid)
		assert.Equal(t, err, tc.want)
	}
}
