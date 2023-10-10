package backup

import (
	"errors"
	"fmt"
	"log"
	"sync"
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

func (s *Service) ListBackups() ([]*entity.ClientRun, error) {
	result, err := s.backup.ListClientRunsAll()

	if result == nil {
		return nil, entity.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	var response []*entity.ClientRun
	for _, i := range result {
		temp := entity.ClientRun{
			Policy:     i.Policy,
			Client:     i.Client,
			ID:         i.ID,
			Name:       i.Name,
			Status:     i.Status,
			TotalFiles: i.TotalFiles,
		}
		response = append(response, &temp)
	}
	return response, nil
}

// func (s *Service) ListBackupsShort()([]*entity.ClientRun, error){
// 	result, err := s.backup
// }

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
	cbm, err := s.bus.Subscribe("file")
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
		newclient, err := entity.NewClientRun(clients[i].Clientname)
		job.job.Clients = append(job.job.Clients, newclient)
		log.Println("CLIENTS: ", job.job.Clients)
		err = s.rabbit.StartBackup(clients[i].ConsumerID, job.job.ID, clients[i].Clientname, backuptype, clients[i].Ignorepath)
		if err != nil {
			newclient.Status = "Failed"
			//return err
		} else {
			fmt.Println("BACKUP reached", i)

			err = job.handleMessage(cbm, clients[i].Clientname)
			//err = job.handleStorageNodeMessage(snm, clients[i].Clientname)
			if err != nil {
				log.Println("Error", err)
				newclient.Status = "Failed"
			}
			newclient.Status = "Success"
			totalFiles := len(newclient.SuccesfullFiles) + len(newclient.FailedFiles)
			newclient.TotalFiles = totalFiles
			currentTime := time.Now()
			runtime := currentTime.Format("01-02-2006")
			newclient.Date = runtime
			id, err := s.backup.Create(newclient, clients[i].ID, policy.PolicyID, runtime)
			if err != nil {
				log.Println("ERROR:", err)
				newclient.Status = "Failed"
			}
			newclient.ID = id

			clients[i].Backups = append(clients[i].Backups, id)
			err = s.client.Update(clients[i])
			if err != nil {
				return err
			}
		}
	}
	backupstruct, err := job.finaliseBackupCompletion(policy.Retention)
	if err != nil {
		log.Println("ERROR finaliseBackupCompletion", err)
	}
	err = policy.AddBackupRun(backupstruct)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	//log.Println("JOB", policy.BackupRun)
	err = s.policy.AddBackupRun(policy.PolicyID, backupstruct)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}

	//backuprun, err := entity.NewBackupRun()
	log.Println("BACKUP SUCCESFULL")
	return nil
}

func (bj *Backup) finaliseBackupCompletion(retention int) (*entity.Backups, error) {
	currentTime := time.Now()
	expiry := currentTime.AddDate(0, 0, retention)
	date := expiry.Format("01-02-2006")
	var success []string
	var fail []string
	//log.Println(currentTime, "retention:", policy.Retention, expiry)
	for _, j := range bj.job.Clients {
		if j.Status == "Success" {
			success = append(success, j.ID)
		} else {
			fail = append(fail, j.ID)
		}
	}
	b := entity.NewBackup(bj.job.ID, "Completed", bj.job.Type, currentTime.Format("01-02-2006"), date, currentTime.Format("15:04:05"), success, fail)
	return b, nil
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
			bj.bus.Unsubscribe("StorageNodeJob", channel)
			//	close(channel)
			return policyID, nil
		default:
			time.Sleep(1 * time.Second)
		}
	}
	return "", entity.ErrTimeOut
}

func (bj *Backup) handleMessage(channel rabbitBus.EventChannel, client string) error {
	foundClient, err := bj.job.GetClient(client)
	if err != nil {
		log.Println(err)
	}
	clmessage := make(chan *entity.ClientFile)
	snmessage := make(chan *entity.ClientFile)
	mutex := &sync.Mutex{}
	go bj.handleClientMessage(mutex, clmessage, foundClient)
	go bj.handleStorageNodeMessage(mutex, snmessage, foundClient)
	for msg := range channel {
		var data entity.ClientFileHolder
		err := mapstructure.Decode(msg.Data, &data)
		if err != nil {
			return err
		}
		if data.File.ID == "Completion" && data.File.Status == "Success" {
			err := checkClientIsConsistent(foundClient.SuccesfullFiles, foundClient.FailedFiles)
			if err != nil {
				log.Println("ERROR Failed backup inconsistent!!", err)
				err = bj.bus.Unsubscribe("file", channel)
				if err != nil {
					log.Println("ERROR:", err)
					return err
				}
				close(clmessage)
				close(snmessage)
				return err
			}
			log.Println("Storage Node has finished Client: ", client, "Succesfully")
			close(clmessage)
			close(snmessage)
			break

		}
		if data.Type == "clientfile" {
			clmessage <- data.File
		} else if data.Type == "storagenodefile" {
			snmessage <- data.File
		}
	}
	err = bj.bus.Unsubscribe("file", channel)
	if err != nil {
		log.Println("ERROR:", err)
		return err
	}
	return nil
}

//handles All the messages that come in from clients, matching them and updating status on them
func (bj *Backup) handleClientMessage(lock *sync.Mutex, channel chan (*entity.ClientFile), client *entity.ClientRun) {
	//newclient, err := entity.NewClientRun(client)
	//bj.job.Clients = append(bj.job.Clients, newclient)

	for msg := range channel {

		//var file = entity.ClientFile{}
		//mapstructure.Decode(msg.Data, &file)
		fmt.Println("recieved message from CLIENT: ", msg.ID)
		if msg.Status == "Finished" {
			fmt.Println("recieved complete message from client: ", msg.ID)
			//err := bj.bus.Unsubscribe("clientfile", channel)
			break
		} else if msg.Status == "Success" {
			lock.Lock()
			exists, err := client.GetFile(msg.ID)
			if err == nil {
				log.Println("Storage Node beat me there:", exists.ID)
				result := checkFileIsConsistent(exists.Checksum, msg.Checksum)
				if result == true {
					msg.ChangeStatus("Complete")
					log.Println("recieved file from STORAGE NODE: ", msg.ID)
				} else {
					client.FailedFiles[msg.ID] = msg
				}
			}
			client.SuccesfullFiles[msg.ID] = msg
			lock.Unlock()
		} else if msg.Status == "Failed" {
			lock.Lock()
			client.FailedFiles[msg.ID] = msg
			lock.Unlock()
		}

	}
	log.Println("Client", "has finished")
	//fmt.Println("New File", bj.job.Clients[0].Files)
}

//handles all messages that come in from the storagenode. It either marks files as confirmed or unsuccesfull
func (bj *Backup) handleStorageNodeMessage(lock *sync.Mutex, channel chan (*entity.ClientFile), client *entity.ClientRun) error {
	for msg := range channel {
		log.Println("recieved file from STORAGE NODE: ", msg.ID)
		if msg.ID == "Completion" && msg.Status == "Success" {
			lock.Lock()
			err := checkClientIsConsistent(client.SuccesfullFiles, client.FailedFiles)
			if err != nil {
				lock.Unlock()
				return err
			}
			lock.Unlock()
			log.Println("Storage Node has finished Client: ", "Succesfully")
			break
		}
		lock.Lock()
		exists, err := client.GetFile(msg.ID)

		if err != nil {
			log.Println("Could not find: ", msg)
			if msg.Status == "Success" {
				client.SuccesfullFiles[msg.ID] = msg
			}
			//return err
		} else {
			result := checkFileIsConsistent(exists.Checksum, msg.Checksum)
			if result == true {
				msg.ChangeStatus("Complete")
				log.Println("recieved file from STORAGE NODE: ", msg.ID)
			} else {
				return err
			}
		}
		lock.Unlock()

		// if exists.Checksum == msg.Checksum {
		// 	msg.ChangeStatus("Complete")
		// 	log.Println("recieved file from STORAGE NODE: ", msg.ID)
		// } else {
		// 	return err
		// }
	}
	return nil
}

func checkFileIsConsistent(checksum1, checksum2 string) bool {
	if checksum1 == checksum2 {
		return true
	} else {
		return false
	}
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
