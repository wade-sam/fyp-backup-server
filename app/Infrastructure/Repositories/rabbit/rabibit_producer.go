package rabbit

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
	//"github.com/wade-sam/fyp-backup-server/entity"
)

type ProducerConfig struct {
	ExchangeName string
	ExchangeType string
	MaxAttempt   int
	Interval     time.Duration
	connection   *amqp.Connection
}

func (b *Broker) Publish(routekey, messagetype string, body *DTO) error {
	channel, err := b.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()
	//dto, err := Serialize(body)
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	fmt.Println("DATA", body.Data, "ID", body.ID)
	err = channel.Publish(
		b.Producer.ExchangeName,
		routekey,
		false,
		false,
		amqp.Publishing{
			Type:        messagetype,
			ContentType: "application/json",
			Body:        []byte(data),
		})
	if err != nil {
		return err
	}
	return nil

}

type StoragenodeData struct {
	Clients    []string `json:"clients"`
	PolicyName string   `json:"policyname"`
}

type ClientData struct {
	Type     string   `json: "type"`
	Client   string   `json: "name"`
	PolicyID string   `json: "policy"`
	Data     []string `json: "ignorelist"`
}

func (b *Broker) SearchForNewClient() error {
	messagetype := "New.Client"
	key := "newclient"
	dto := DTO{}
	err := b.Publish(key, messagetype, &dto)
	if err != nil {
		return err
	}
	return nil
}

func (b *Broker) DirectoryScan(client string) error {
	messagetype := "Directory.Scan"
	key := fmt.Sprintf("%v", client)
	dto := DTO{}
	err := b.Publish(key, messagetype, &dto)
	if err != nil {
		return err
	}
	return nil
}

//Storagenode needs policyname and clients to start backup
func (b *Broker) StartBackup(clientID, policyID, clientname, backuptype string, ignorepath []string) error {

	dto := DTO{}
	data := ClientData{
		Client:   clientname,
		PolicyID: policyID,
		Data:     ignorepath,
	}
	dto.Data = data
	// temp := ClientData{
	// 	Type:     "FUCK OFF!",
	// 	Client:   "FUCK OFF!",
	// 	PolicyID: "Friday Backup",
	// 	Data:     []string{"/home"},
	// }
	dto.Data = data
	//messagetype := ""
	if backuptype == "Full" {
		err := b.Publish(clientID, "Full.Backup", &dto)
		if err != nil {
			return err
		}

	} else {
		err := b.Publish(clientID, "Inc.Backup", &dto)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Broker) StartStorageNode(clients []string, storagenode, policy string) error {
	dto := DTO{}
	temp := StoragenodeData{
		Clients:    clients,
		PolicyName: policy,
	}
	dto.Data = temp

	err := b.Publish(storagenode, "New.Backup.Job", &dto)
	if err != nil {
		return err
	}
	return nil
}
