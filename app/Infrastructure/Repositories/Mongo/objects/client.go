package objects

type Client struct {
	Clientname    string   `bson:"clientname"`
	Policies      []string `bson:"policies"`
	Directorytree string   `bson: "treepath"`
	Ignorepath    string   `bson: "ignore"`
	Backups       string   `bson: "backups"`
}
