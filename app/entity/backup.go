package entity

import (
	"fmt"
	"time"
)

type BackupRun struct {
	PolicyName string
	ID         string
	Type       string
	Date       string
	Clients    []ClientRun
	Status     string
}

type ClientRun struct {
	Name   string
	Status string
	Files  map[string]ClientFile
}

type ClientFile struct {
	ID       string
	Status   string
	Checksum string
}

func NewBackupRun(policyname, Type string) (*BackupRun, error) {
	run := &BackupRun{
		PolicyName: policyname,
		Type:       Type,
	}
	return run, nil

}

func (br *BackupRun) CreateNameTimeProperties() {
	date := time.Now()
	name := fmt.Sprintf("%v-%v", br.PolicyName, date.Format("01-02-2006 15:04:05"))
	br.ID = name
	br.Date = date.Format("01-02-2006 15:04:05")
}

func (br *BackupRun) AddClient(client ClientRun) error {
	_, err := br.GetClient(client.Name)
	if err != nil {
		return ErrClientAlreadyAdded
	}
	br.Clients = append(br.Clients, client)
	return nil
}

func (br *BackupRun) GetClient(name string) (string, error) {
	for i := range br.Clients {
		if br.Clients[i].Name == name {
			return name, nil
		}
	}
	return name, ErrNotFound
}

func NewClientFile(id string, file *File) (*ClientFile, error) {
	c := &ClientFile{
		ID:     id,
		Status: "Start",
	}

	return c, nil
}

func NewClientRun(name string) (*ClientRun, error) {
	return &ClientRun{
		Name:   name,
		Status: "In Progress",
		Files:  make(map[string]ClientFile),
	}, nil
}

func (cf *ClientFile) ChangeStatus(status string) {
	cf.Status = status
}

func (cr *ClientRun) ChangeClientRunStatus(status string) {
	cr.Status = status
}
