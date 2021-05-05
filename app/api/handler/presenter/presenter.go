package presenter

import "github.com/wade-sam/fyp-backup-server/entity"

type Client struct {
	ID            string   `json:"id"`
	ConsumerID    string   `json:"consumerID"`
	Clientname    string   `json:"clientname"`
	Policies      []string `json:"policies"`
	Directorytree []string `json:"tree"`
	Ignorepath    []string `json:"ignorepath"`
	Backups       []string `json:"backups"`
}

type Policy struct {
	PolicyID   string       `json:"id"`
	Policyname string       `json:"policyname"`
	RunTime    string       `josn:"runtime"`
	Clients    [][]string   `json:"clients"`
	Retention  int          `json:"retention"`
	State      string       `json:"state"`
	Type       string       `json:"type"`
	Fullbackup []string     `json:"fullbackup"`
	IncBackup  []string     `json:"incbackup"`
	BackupRun  []*BackupRun `json:"BackupRun"`
}

type BackupRun struct {
	//PolicyName string
	ID                 string   `json:"backuprunid"`
	Type               string   `json:"runtype"`
	Date               string   `json:"rundate"`
	Expiry             string   `json:"expiry"`
	RunTime            string   `json:"runtime"`
	SuccessFullClients []string `json:"successclients"`
	FailedClients      []string `json:"failclients"`
	//	Clients []*ClientRun `json:"clientruns"`
	Status string `json:"backupstatus"`
}

type ClientRun struct {
	ID              string                 `json:"clientrunid"`
	Name            string                 `json:"runname"`
	Client          []string               `json:"clientname"`
	Policy          []string               `json:"policyname"`
	Status          string                 `json:"runstatus"`
	TotalFiles      int                    `json:"totalfiles"`
	SuccesfullFiles map[string]*ClientFile `json:"successfiles"`
	FailedFiles     map[string]*ClientFile `json:"failfiles"`
}

type ClientRunSmall struct {
	ID         string `json:"clientrunid"`
	Name       string `json:"runname"`
	Client     string `json:"clientname"`
	Policy     string `json:"policyname"`
	Status     string `json:"runstatus"`
	TotalFiles int    `json:"totalfiles"`
}

type ClientFile struct {
	ID       string `bson:"fileid"`
	Status   string `bson:"status"`
	Checksum string `bson:"checksum"`
}

type Directory struct {
	Path       string                `json:"value,ommitempty"`
	Name       string                `json:"label,omitempty"`
	Properties []string              `json:"properties,omitempty"`
	Files      []*entity.File        `json:"files,omitempty"`
	Folders    map[string]*Directory `json:"folders,omitempty"`
	NewFolders []*Directory          `json:"children, ommitempty"`
}
