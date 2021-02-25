package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
	"github.com/wade-sam/fyp-backup-server/usecase/client"
	"github.com/wade-sam/fyp-backup-server/usecase/client/mock"
)

func newFixtureClient() *entity.Client {
	return &entity.Client{
		Clientname:    "sam's mackbook pro",
		Policies:      []string{"wednesday's backup", "Friday Full"},
		Directorytree: "/, /home/ /home/sam",
		Ignorepath:    "/home/test",
		Backups:       "0",
	}
}

func listClients() []*entity.Client {
	lists := []*entity.Client{}
	for i := 0; i < 5; i++ {
		client := newFixtureClient()
		lists = append(lists, client)
	}
	return lists
}

func Test_CreateClient(t *testing.T) {
	c := newFixtureClient()
	//Setup expectations
	mockRepo := new(mock.MockRepository)
	mockRepo.On("Create", c).Return(c, nil)
	testService := client.NewService(mockRepo)
	result, err := testService.CreateClient(c)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, result, c)

}

func Test_GetClient(t *testing.T) {
	mockRepo := new(mock.MockRepository)
	c := newFixtureClient()
	mockRepo.On("Get", c.Clientname).Return(c, nil)
	testService := client.NewService(mockRepo)
	result, err := testService.GetClient(c.Clientname)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, result, c)
}

func Test_ListClients(t *testing.T) {
	clients := listClients()
	mockRepo := new(mock.MockRepository)
	mockRepo.On("List").Return(clients, nil)
	testService := client.NewService(mockRepo)
	result, err := testService.ListClients()
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.Equal(t, result, clients)
}

func Test_UpdateClient(t *testing.T) {
	c := newFixtureClient()
	d := entity.Client{
		Policies:      []string{"wednesday's backup", "Friday Full"},
		Directorytree: "/, /home/ /home/sam",
		Ignorepath:    "/home/test",
		Backups:       "0",
	}
	mockRepo := new(mock.MockRepository)
	mockRepo.On("Update", c).Return(nil)
	testService := client.NewService(mockRepo)
	err := testService.UpdateClient(c)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	err = testService.UpdateClient(&d)
	mockRepo.AssertExpectations(t)
	assert.Equal(t, entity.ErrInvalidEntity, err)
}

func Test_DeleteClient(t *testing.T) {
	c := newFixtureClient()
	mockRepo := new(mock.MockRepository)
	mockRepo.On("Delete", c.Clientname).Return(nil)
	mockRepo.On("Get", c.Clientname).Return(c, nil)
	testService := client.NewService(mockRepo)
	err := testService.DeleteClient(c.Clientname)
	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)

}
