package entity

type BackupRun struct {
	//PolicyName string
	ID                 string
	Type               string
	Date               string
	Expiry             string
	RunTime            string
	Clients            []*ClientRun
	SuccessFullClients []string
	FailedClients      []string
	Status             string
}

type ClientRun struct {
	Date            string
	Policy          string
	Client          string
	ID              string
	Name            string
	Status          string
	TotalFiles      int
	SuccesfullFiles map[string]*ClientFile
	FailedFiles     map[string]*ClientFile
}

type ClientFile struct {
	ID       string `bson:"fileid"`
	Status   string `bson:"status"`
	Checksum string `bson:"checksum"`
}

type ClientFileHolder struct {
	Type string
	File *ClientFile
}

func NewBackupRun(policyname, Type string) (*BackupRun, error) {
	run := &BackupRun{
		//	PolicyName: policyname,
		Type: Type,
	}
	return run, nil

}

// func (br *BackupRun) CreateNameTimeProperties() {
// 	date := time.Now()
// 	name := fmt.Sprintf("%v-%v", br.PolicyName, date.Format("01-02-2006 15:04:05"))
// 	br.ID = name
// 	br.Date = date.Format("01-02-2006 15:04:05")
// }

func (br *BackupRun) AddClient(client *ClientRun) error {
	_, err := br.GetClient(client.Name)
	if err != nil {
		return ErrClientAlreadyAdded
	}
	br.Clients = append(br.Clients, client)
	return nil
}

func (br *BackupRun) GetClient(name string) (*ClientRun, error) {
	for i := range br.Clients {
		if br.Clients[i].Name == name {
			return br.Clients[i], nil
		}
	}
	return nil, ErrNotFound
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
		Name:            name,
		Status:          "In Progress",
		SuccesfullFiles: make(map[string]*ClientFile),
		FailedFiles:     make(map[string]*ClientFile),
	}, nil
}

func (cr *ClientRun) GetFile(id string) (*ClientFile, error) {
	if file, exist := cr.SuccesfullFiles[id]; exist {
		return file, nil
	}
	if file, exist := cr.FailedFiles[id]; exist {
		return file, nil
	}

	return nil, ErrCouldNotFindFile
}

func (cf *ClientFile) ChangeStatus(status string) {
	cf.Status = status
}

func (cr *ClientRun) ChangeClientRunStatus(status string) {
	cr.Status = status
}
