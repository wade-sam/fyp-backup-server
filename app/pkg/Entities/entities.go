package Entities

//Contains all the  entites/models used across the system.
type Client struct {
	Clientname   string `json:"clientname" bson:"clientname" validate:"required"`
	Consumername string `json:"consumername" bson:"consumername" validate:"required"`
	Policies     string `json:"policies" bson:"policies"`
	Backuptree   string `json:"backuptree" bson:"backuptree"`
	Ignorepath   string `json:"ignorepath" bson:"ignorepath"`
	Backups      string `json:"backups" bson:"backups"`
}

type Policy struct {
	Policyname  string   `json:"policyname" bson:"policyname" validate:"required"`
	Clients     []string `json:"clients" bson:"clients"`
	Retention   int      `json:"retention" bson:"retention" validate:"required"`
	Scale       string   `json:"scale" bson:"scale" validate:"required"`
	Fullbackup  []string `json:"fullbackup" bson:"fullbackup" validate:"required"`
	Incremental []string `json:"incremental" bson:"incremental"`
}

type File struct {
	FileName      string `json:"filename"`
	DirectoryPath string `json:"directory-path"`
	Permissions   string `json:"file-permissions"`
	UID           int    `json:"file-user-id"`
	GID           int    `json:"file-group-id"`
	Checksum      string `json:"checksum"`
}

type DirectoryScan struct {
	Files map[string]File `files`
	Keys  []string        `keys`
}
