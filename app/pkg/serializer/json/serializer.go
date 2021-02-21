package json

import (
	//"github.com/wade-sam/fyp-backup-server/pkg/Client"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/wade-sam/fyp-backup-server/pkg/Entities"
)

type Filescan struct{}
type BackupFiles struct{}
type Message struct{}
type Config struct{}

//Receives a directoryscan
func (f *Filescan) DecodeDirectoryScan(input []byte) (*Entities.FileScan, error) {
	filescan := &Entities.FileScan{}
	if err := json.Unmarshal(input, filescan); err != nil {
		return nil, errors.Wrap(err, "serializer.Filescan.Decode")
	}
	return filescan, nil

}

func (f *Filescan) EncodeDirectoryScan(input *Entities.FileScan) ([]byte, error) {
	rawFilescan, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Filescan.Encode")
	}
	return rawFilescan, nil
}
