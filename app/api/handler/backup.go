package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/wade-sam/fyp-backup-server/api/handler/presenter"
	"github.com/wade-sam/fyp-backup-server/entity"
	"github.com/wade-sam/fyp-backup-server/usecase/backup"
	"github.com/wade-sam/fyp-backup-server/usecase/client"
	"github.com/wade-sam/fyp-backup-server/usecase/policy"
)

func startBackup(client client.UseCase, policy policy.UseCase, backup backup.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Starting Backup"
		vars := mux.Vars(r)
		err := backup.StartBackup(vars["id"], "Full")
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func listBackups(client client.UseCase, policy policy.UseCase, backup backup.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Reading Backups"
		var data []*entity.ClientRun
		data, err := backup.ListBackups()
		w.Header().Set("Content-Type", "application/json")
		if err != nil && err != entity.ErrNotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var jclientruns []*presenter.ClientRun
		for _, j := range data {

			policy, err := policy.GetPolicyName(j.Policy)
			if err != nil {
				log.Println("ERROR1", err, j.ID)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(errorMessage))
				return
			}
			p := []string{policy, j.Policy}
			client, err := client.GetClientName(j.Client)
			if err != nil {
				log.Println("ERROR2", err, j.ID)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(errorMessage))
				return
			}
			c := []string{client, j.Client}
			j.Policy = policy
			j.Client = client
			var jclientrun presenter.ClientRun
			mapstructure.Decode(j, &jclientrun)
			jclientrun.Client = c
			jclientrun.Policy = p
			jclientruns = append(jclientruns, &jclientrun)
		}
		if err := json.NewEncoder(w).Encode(jclientruns); err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
		//ListClientRuns
	})
}

func getBackup(client client.UseCase, policy policy.UseCase, backup backup.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func MakeBackupHandlers(r *mux.Router, clientservice client.UseCase, policyservice policy.UseCase, backupservice backup.UseCase) {
	r.Handle("/api/backups/start/{id}", startBackup(clientservice, policyservice, backupservice)).Methods("POST").Name("startBackup")
	r.Handle("/api/backups/list", listBackups(clientservice, policyservice, backupservice)).Methods("GET").Name("listBackups")
	r.Handle("/api/backups/get/{id}", getBackup(clientservice, policyservice, backupservice)).Methods("GET").Name("getBackup")
}
