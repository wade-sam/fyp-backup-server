package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
)

func TestNewClient(t *testing.T) {
	u, err := entity.NewClient("Sam's MacBook Pro")
	assert.Nil(t, err)
	assert.Equal(t, u.Clientname, "Sam's MacBook Pro")
}

func TestAddPolicy(t *testing.T) {
	u, _ := entity.NewClient("Sam's MacBook Pro")
	policy := "Wednesday Backup Run"
	err := u.AddPolicy(policy)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(u.Policies))
	err = u.AddPolicy(policy)
	assert.Equal(t, entity.ErrPolicyAlreadyAdded, err)
}

func TestRemovePolicy(t *testing.T) {
	u, _ := entity.NewClient("Sam's MacBook Pro")
	err := u.RemovePolicy("tuesday's backup")
	assert.Equal(t, entity.ErrNotFound, err)
	policy := "Wednesday Backup Run"
	_ = u.AddPolicy(policy)
	err = u.RemovePolicy(policy)
	assert.Nil(t, err)
}

func TestGetPolicy(t *testing.T) {
	u, _ := entity.NewClient("Sam's MacBook Pro")
	bpolicy := "Wednesday Backup Run"
	_ = u.AddPolicy(bpolicy)
	policy, err := u.GetPolicy(bpolicy)
	assert.Nil(t, err)
	assert.Equal(t, policy, bpolicy)
	_, err = u.GetPolicy("testing")
	assert.Equal(t, entity.ErrNotFound, err)

}

func TestClientValidate(t *testing.T) {
	type test struct {
		clientname string
		want       error
	}

	tests := []test{
		{
			clientname: "sam wade",
			want:       nil,
		},
		{
			clientname: "",
			want:       entity.ErrInvalidEntity,
		},

		{
			clientname: "sam wade",
			want:       entity.ErrInvalidEntity,
		},
		{
			clientname: "",
			want:       entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewClient(tc.clientname)
		assert.Equal(t, err, tc.want)
	}
}
