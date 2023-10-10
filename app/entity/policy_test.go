package entity_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
)

func TestNewPolicy(t *testing.T) {
	fullschedule := []string{"Monday", "Thursday", "Sunday"}
	incremental := []string{}
	clients := []string{"client 1"}
	p, err := entity.NewPolicy("Wednesday's backup", "full", 10, fullschedule, incremental, clients)
	assert.Nil(t, err)
	assert.Equal(t, p.Policyname, "Wednesday's backup")
	assert.Equal(t, p.State, "active")
	assert.Equal(t, p.Retention, 10)
	assert.Equal(t, p.Clients, clients)
	fmt.Println(p.Fullbackup)
	assert.Equal(t, p.Fullbackup, fullschedule)
	assert.Equal(t, p.IncBackup, incremental)
}

func TestAddClient(t *testing.T) {
	fullschedule := []string{"Monday", "Thursday", "Sunday"}
	incremental := []string{}
	clients := []string{"client 1"}
	p, _ := entity.NewPolicy("Wednesday's backup", "full", 10, fullschedule, incremental, clients)
	client := "client 2"
	err := p.AddClient(client)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(p.Clients))
	err = p.AddClient(client)
	assert.Equal(t, entity.ErrClientAlreadyAdded, err)
}

func TestRemoveClient(t *testing.T) {
	fullschedule := []string{"Monday", "Thursday", "Sunday"}
	incremental := []string{}
	client := "client 10"
	clients := []string{client}
	p, _ := entity.NewPolicy("Wednesday's backup", "full", 10, fullschedule, incremental, clients)
	err := p.RemoveClient("client 1")
	assert.Equal(t, entity.ErrNotFound, err)
	//client := entity.NewID()
	_ = p.AddClient(client)
	err = p.RemoveClient(client)
	assert.Nil(t, err)
}

func TestGetClient(t *testing.T) {
	fullschedule := []string{"Monday", "Thursday", "Sunday"}
	incremental := []string{}
	client := "client 10"
	clients := []string{client}
	p, _ := entity.NewPolicy("Wednesday's backup", "full", 10, fullschedule, incremental, clients)
	client, err := p.GetClient(client)
	assert.Nil(t, err)
	assert.Equal(t, client, client)
	_, err = p.GetClient("client")
	assert.Equal(t, entity.ErrNotFound, err)
}

func TestGetState(t *testing.T) {
	fullschedule := []string{"Monday", "Thursday", "Sunday"}
	incremental := []string{}
	client := "client 1"
	clients := []string{client}
	p, _ := entity.NewPolicy("Wednesday's backup", "full", 10, fullschedule, incremental, clients)
	state, err := p.GetState()
	assert.Nil(t, err)
	assert.Equal(t, state, "active")
	err = p.RemoveClient(client)
	p.AddState()
	state, err = p.GetState()
	assert.Nil(t, err)
	assert.Equal(t, state, "inactive")
}

func TestPolicyValidate(t *testing.T) {
	type test struct {
		Policyname string
		Clients    []string
		Retention  int
		State      string
		Type       string
		Fullbackup []string
		IncBackup  []string
		want       error
	}
	tests := []test{
		{
			Policyname: "wednesday backup",
			Clients:    []string{"client 1"},
			Retention:  10,
			Type:       "full",
			Fullbackup: []string{"Monday", "Thursday", "Sunday"},
			IncBackup:  []string{},
			want:       nil,
		},
		{
			Policyname: "",
			Clients:    []string{"client 1"},
			Retention:  10,
			Type:       "full",
			Fullbackup: []string{"Monday", "Thursday", "Sunday"},
			IncBackup:  []string{},
			want:       entity.ErrInvalidEntity,
		},
		{
			Policyname: "wednesday's backup",
			Clients:    []string{"client 1"},
			Retention:  0,
			Type:       "full",
			Fullbackup: []string{"Monday", "Thursday", "Sunday"},
			IncBackup:  []string{},
			want:       entity.ErrInvalidEntity,
		},
		{
			Policyname: "wednesday's backup",
			Clients:    []string{"client 1"},
			Retention:  10,
			Type:       "full",
			Fullbackup: []string{},
			IncBackup:  []string{"Monday", "Thursday", "Sunday"},
			want:       entity.ErrInvalidBackupPlan,
		},
		{
			Policyname: "wednesday's backup",
			Clients:    []string{"client 1"},
			Retention:  10,
			Type:       "both",
			Fullbackup: []string{},
			IncBackup:  []string{"Monday", "Thursday", "Sunday"},
			want:       entity.ErrInvalidBackupPlan,
		},
		{
			Policyname: "wednesday's backup",
			Clients:    []string{"client 1"},
			Retention:  10,
			Type:       "both",
			Fullbackup: []string{"Sunday"},
			IncBackup:  []string{"Monday", "Thursday", "Sunday"},
			want:       entity.ErrInvalidBackupPlan,
		},
		{
			Policyname: "wednesday's backup",
			Clients:    []string{"client 1"},
			Retention:  10,
			Type:       "",
			Fullbackup: []string{},
			IncBackup:  []string{"Monday", "Thursday", "Sunday"},
			want:       entity.ErrInvalidBackupPlan,
		},
		{
			Policyname: "wednesday's backup",
			Clients:    []string{"client 1"},
			Retention:  10,
			Type:       "both",
			Fullbackup: []string{"Monday"},
			IncBackup:  []string{"Monday", "Thursday", "Sunday"},
			want:       entity.ErrInvalidBackupPlan,
		},
		{
			Policyname: "wednesday's backup",
			Clients:    []string{"client 1"},
			Retention:  10,
			Type:       "both",
			Fullbackup: []string{},
			IncBackup:  []string{},
			want:       entity.ErrInvalidBackupPlan,
		},
	}

	for _, tc := range tests {
		_, err := entity.NewPolicy(tc.Policyname, tc.Type, tc.Retention, tc.Fullbackup, tc.IncBackup, tc.Clients)
		assert.Equal(t, err, tc.want)
	}
}
