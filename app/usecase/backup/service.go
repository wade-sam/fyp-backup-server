package backup

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/wade-sam/fyp-backup-server/entity"
	"github.com/wade-sam/fyp-backup-server/rabbitBus"
)

type Service struct {
	client ClientRepository
	policy PolicyRepository
	backup BackupRepository
	rabbit RabbitRepository
	bus    BusRepository
}

func NewService(cr ClientRepository, pr PolicyRepository, br BackupRepository, rr RabbitRepository, bsr BusRepository) *Service {
	return &Service{
		client: cr,
		policy: pr,
		backup: br,
		rabbit: rr,
		bus:    bsr,
	}
}

type Backup struct {
	job *entity.BackupRun
	bus BusRepository
}

func newBackupRun(name, Type string) (*Backup, error) {
	backupjob, err := entity.NewBackupRun(name, Type)
	if err != nil {
		return nil, err
	}

	backup := Backup{
		job: backupjob,
	}

	return &backup, nil
}

/*
Statistics and metadata regarding the backup are stored in the database, but the files holding
all information about which files were backed up are stored as a file on the storage node
*/
func (s *Service) StartBackup(name, backuptype string) error {
	policy, err := s.policy.Get(name)
	if err != nil {
		return entity.ErrNotFound
	}

	var clients []*entity.Client
	for i := range policy.Clients {
		client, err := s.client.Get(policy.Clients[i])
		if err != nil {
			return entity.ErrNotFound
		}
		clients = append(clients, client)

	}
	cbs, err := s.bus.Subscribe("clientjob")
	if err != nil {
		return nil
	}
	cbm, err := s.bus.Subscribe("clientfile")
	if err != nil {
		return nil
	}
	snm, err := s.bus.Subscribe("storagenodefile")
	if err != nil {
		return nil
	}
	sni, err := s.bus.Subscribe("StorageNodeJob")
	if err != nil {
		return nil
	}

	fmt.Println("SUBSCRIBED", cbs, cbm, snm)

	if backuptype == "" {
		if found := checkbackupType(policy.Fullbackup); found == true {
			backuptype = "Full"
		} else if found := checkbackupType(policy.IncBackup); found == true {
			backuptype = "Incremental"
		}
	}

	job, err := newBackupRun(policy.Policyname, backuptype)
	job.bus = s.bus
	//fmt.Println("BACKUP found", job.job.PolicyName)
	if err != nil {
		return errors.New("ERROR")
	}
	clientnames := []string{}
	for i := range clients {
		clientnames = append(clientnames, clients[i].Clientname)
	}

	err = s.rabbit.StartStorageNode(clientnames, "storagenode", policy.Policyname)
	if err != nil {
		return nil
	}
	jobID, err := job.handleStorageNodeResponse(sni)
	if err != nil {
		return err
	}
	fmt.Println("JOB ID", jobID)
	job.job.ID = jobID

	//Add backuprun to policy

	for i := range clients {
		err := s.rabbit.StartBackup(clients[i].ConsumerID, job.job.ID, clients[i].Clientname, backuptype, clients[i].Ignorepath)
		if err != nil {
			return err
		}
		fmt.Println("BACKUP reached", i)
		newclient, err := entity.NewClientRun(clients[i].Clientname)
		job.job.Clients = append(job.job.Clients, newclient)

		go job.handleClientMessage(cbm, clients[i].Clientname)
		err = job.handleStorageNodeMessage(snm, clients[i].Clientname)
		if err != nil {
			log.Println("Error", err)
		}
		id, err := s.backup.Create(newclient)
		if err != nil {
			return err
		}
		newclient.ID = id
		// client, err := s.client.Get(policy.Clients[i])
		// if err != nil {
		// 	return err
		// }
		// s.rabbit.StartBackup(client.ConsumerID, "storagenode", backuptype, client.Ignorepath)
		// select {}
	}
	policy.AddBackupRun(job.job)
	err = s.policy.Update(policy)
	if err != nil {
		return err
	}

	//backuprun, err := entity.NewBackupRun()

	return nil
}

func checkbackupType(days []string) bool {
	t := time.Now()
	day := t.Weekday()
	for i := range days {
		if days[i] == string(day) {
			return true
		}
	}
	return false
}

func (bj *Backup) handleStorageNodeResponse(channel rabbitBus.EventChannel) (string, error) {
	for i := 1; i < 30; i++ {
		select {
		case msg := <-channel:
			policyID := ""
			mapstructure.Decode(msg.Data, &policyID)
			close(channel)
			return policyID, nil
		default:
			time.Sleep(1 * time.Second)
		}
	}
	return "", entity.ErrTimeOut
}

//Handles Initialising New clients and the files to expect from the file
// func (bj *Backup) handleClientSchedule(channel rabbitBus.EventChannel, client string) error {
// 	for i := 1; i < 10000; i++ {
// 		select {
// 		case msg := <-channel:
// 			newclient, err := entity.NewClientRun(client)
// 			if err != nil {
// 				return nil
// 			}
// 			files := make(map[string]entity.File)
// 			mapstructure.Decode(msg.Data, &files)
// 			for i, j := range files {
// 				file, err := entity.NewClientFile(files[i].Path, &j)
// 				if err != nil {
// 					return err
// 				}
// 				newclient.Files[file.ID] = file
// 			}
// 			bj.job.Clients = append(bj.job.Clients, newclient)
// 			close(channel)
// 			return nil
// 		default:
// 			time.Sleep(2 * time.Second)
// 		}
// 	}
// 	close(channel)
// 	return entity.ErrFileNotFound
// }

//handles All the messages that come in from clients, matching them and updating status on them
func (bj *Backup) handleClientMessage(channel rabbitBus.EventChannel, client string) {
	//newclient, err := entity.NewClientRun(client)
	//bj.job.Clients = append(bj.job.Clients, newclient)
	foundClient, err := bj.job.GetClient(client)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Found Client", client)
	for msg := range channel {

		var file = entity.ClientFile{}
		mapstructure.Decode(msg.Data, &file)
		fmt.Println("recieved message from CLIENT: ", file.ID)
		if file.Status == "Finished" {
			fmt.Println("recieved complete message from client: ", file.ID)
			err := bj.bus.Unsubscribe("clientfile", channel)
			if err != nil {
				log.Println("ERROR unsubscribing", err)
			}
			break
		} else if file.Status == "Success" {
			foundClient.SuccesfullFiles[file.ID] = &file
		} else if file.Status == "Failed" {
			foundClient.FailedFiles[file.ID] = &file
		}

	}
	log.Println("Client", client, "has finished")
	//fmt.Println("New File", bj.job.Clients[0].Files)
}

//handles all messages that come in from the storagenode. It either marks files as confirmed or unsuccesfull
func (bj *Backup) handleStorageNodeMessage(channel rabbitBus.EventChannel, client string) error {
	foundClient, err := bj.job.GetClient(client)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Found Client", client)
	for msg := range channel {
		var file = entity.ClientFile{}
		mapstructure.Decode(msg.Data, &file)
		if file.ID == "Completion" && file.Status == "Success" {
			err := checkClientIsConsistent(foundClient.SuccesfullFiles, foundClient.FailedFiles)
			if err != nil {
				errs := bj.bus.Unsubscribe("storagenodefile", channel)
				if errs != nil {
					log.Println("ERROR unsubscribing", err)
					return err
				}
				return err
			}
			errs := bj.bus.Unsubscribe("storagenodefile", channel)
			if errs != nil {
				log.Println("ERROR unsubscribing", err)
				return err
			}
			log.Println("Storage Node has finished Client: ", client, "Succesfully")
			break
		}
		exits, err := foundClient.GetFile(file.ID)
		if err != nil {
			log.Println("Could not find: ", msg.Data)
			return err
		}
		if exits.Checksum == file.Checksum {
			file.ChangeStatus("Complete")
			log.Println("recieved file from STORAGE NODE: ", file.ID)
		} else {
			return err
		}

	}
	return nil
}

func checkClientIsConsistent(succesful, failure map[string]*entity.ClientFile) error {
	for _, j := range succesful {
		if j.Status == "Start" {
			return entity.ErrInconsistenciesInCompletedFiles
		} else {
			continue
		}
	}
	for _, j := range failure {
		if j.Status == "Start" {
			return entity.ErrInconsistenciesInCompletedFiles
		} else {
			continue
		}
	}
	return nil
}