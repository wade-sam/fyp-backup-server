package presenter

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
	ID      string       `json:"backuprunid"`
	Type    string       `json:"runtype"`
	Date    string       `json:"rundate"`
	Clients []*ClientRun `json:"clientruns"`
	Status  string       `json:"backupstatus"`
}

type ClientRun struct {
	ID              string                 `json:"clientrunid"`
	Name            string                 `json:"runname"`
	Status          string                 `json:"runstatus"`
	TotalFiles      int                    `json:"totalfiles"`
	SuccesfullFiles map[string]*ClientFile `json:"successfiles"`
	FailedFiles     map[string]*ClientFile `json:"failfiles"`
}

type ClientFile struct {
	ID       string `bson:"fileid"`
	Status   string `bson:"status"`
	Checksum string `bson:"checksum"`
}
