package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wade-sam/fyp-backup-server/entity"
	"github.com/wade-sam/fyp-backup-server/usecase/client"
)

func newFixtureClient() *entity.Client {
	return &entity.Client{
		ConsumerID: "host1",
		Clientname:    "sam's mackbook pro",
		Policies:      []string{"p1", "p2", "p3"},
		Directorytree: []string{"/", "/home", "/home/sam"},
		Ignorepath:    []string{"/home/test"},
		Backups:       []string{},
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
	repo := client.NewClientHolder()
	testService := client.NewService(repo)
	p := newFixtureClient()
	_, err := testService.CreateClient(p.Clientname, p.ConsumerID)
	assert.Nil(t, err)
}

func Test_GetClient(t *testing.T) {
	repo := client.NewClientHolder()
	testService := client.NewService(repo)
	p := newFixtureClient()
	id, _ := testService.CreateClient(p.Clientname, p.ConsumerID)
	result, err := testService.GetClient(id)
	assert.Nil(t, err)
	assert.Equal(t, p.Clientname, result.Clientname)
	result, err = testService.GetClient("client2")
	assert.Nil(t, result)
	assert.Equal(t, err, entity.ErrNotFound)
}

func Test_UpdateClient(t *testing.T) {
	repo := client.NewClientHolder()
	testService := client.NewService(repo)
	p := newFixtureClient()
	id, _ := testService.CreateClient(p.Clientname, p.ConsumerID)
	get, err := testService.GetClient(id)
	get.Clientname = "jack's MacBook Pro"
	get.Policies = append(get.Policies, "p3")
	err = testService.UpdateClient(get)
	assert.Nil(t, err)
	updated, err := testService.GetClient(id)
	assert.Nil(t, err)
	assert.Equal(t, "jack's MacBook Pro", updated.Clientname)
	assert.NotNil(t, get.Policies, updated.Policies)
}

func Test_DeleteClient(t *testing.T) {
	repo := client.NewClientHolder()
	testService := client.NewService(repo)
	p1 := newFixtureClient()
	p2 := newFixtureClient()
	p2.ConsumerID = "host2"
	p1id, _ := testService.CreateClient(p1.Clientname, p1.ConsumerID)
	err := testService.DeleteClient(p2.ConsumerID)
	assert.Equal(t, entity.ErrNotFound, err)
	err = testService.DeleteClient(p1id)
	assert.Nil(t, err)
	_, err = testService.GetClient(p1id)
	assert.Equal(t, entity.ErrNotFound, err)

}

func Test_ListClients(t *testing.T) {
	repo := client.NewClientHolder()
	testService := client.NewService(repo)
	p1 := newFixtureClient()
	p2 := newFixtureClient()
	p1.ConsumerID = "host2"
	p2.ConsumerID = "host3"
	plist, err := testService.ListClients()
	assert.Equal(t, entity.ErrNotFound, err)
	_, _ = testService.CreateClient(p1.Clientname, p1.ConsumerID)
	plist, err = testService.ListClients()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(plist))
	_, _ = testService.CreateClient(p2.Clientname, p2.ConsumerID)
	plist2, err := testService.ListClients()
	assert.Nil(t, err)
	assert.Equal(t, 2, len(plist2))
}

// func Test_CreateClient(t *testing.T) {
// 	c := newFixtureClient()
// 	//Setup expectations
// 	mockRepo := new(mock.MockRepository)
// 	mockRepo.On("Create", c).Return(c, nil)
// 	testService := client.NewService(mockRepo)
// 	result, err := testService.CreateClient(c)
// 	mockRepo.AssertExpectations(t)
// 	assert.Nil(t, err)
// 	assert.Equal(t, result, c)

// }

// func Test_GetClient(t *testing.T) {
// 	mockRepo := new(mock.MockRepository)
// 	c := newFixtureClient()
// 	mockRepo.On("Get", c.Clientname).Return(c, nil)
// 	testService := client.NewService(mockRepo)
// 	result, err := testService.GetClient(c.Clientname)
// 	mockRepo.AssertExpectations(t)
// 	assert.Nil(t, err)
// 	assert.Equal(t, result, c)
// }

// func Test_ListClients(t *testing.T) {
// 	clients := listClients()
// 	mockRepo := new(mock.MockRepository)
// 	mockRepo.On("List").Return(clients, nil)
// 	testService := client.NewService(mockRepo)
// 	result, err := testService.ListClients()
// 	mockRepo.AssertExpectations(t)
// 	assert.Nil(t, err)
// 	assert.Equal(t, result, clients)
// }

// func Test_UpdateClient(t *testing.T) {
// 	c := newFixtureClient()
// 	d := entity.Client{
// 		Policies:      []string{"wednesday's backup", "Friday Full"},
// 		Directorytree: "/, /home/ /home/sam",
// 		Ignorepath:    "/home/test",
// 		Backups:       "0",
// 	}
// 	mockRepo := new(mock.MockRepository)
// 	mockRepo.On("Update", c).Return(nil)
// 	testService := client.NewService(mockRepo)
// 	err := testService.UpdateClient(c)

// 	mockRepo.AssertExpectations(t)
// 	assert.Nil(t, err)
// 	err = testService.UpdateClient(&d)
// 	mockRepo.AssertExpectations(t)
// 	assert.Equal(t, entity.ErrInvalidEntity, err)
// }

// func Test_DeleteClient(t *testing.T) {
// 	c := newFixtureClient()
// 	mockRepo := new(mock.MockRepository)
// 	mockRepo.On("Delete", c.Clientname).Return(nil)
// 	mockRepo.On("Get", c.Clientname).Return(c, nil)
// 	testService := client.NewService(mockRepo)
// 	err := testService.DeleteClient(c.Clientname)
// 	mockRepo.AssertExpectations(t)
// 	assert.Nil(t, err)

// }
