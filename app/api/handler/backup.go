package handler

import (
	"net/http"

	"github.com/gorilla/mux"
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
		//ListClientRuns
	})
}

func getBackup(client client.UseCase, policy policy.UseCase, backup backup.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func MakeBackupHandlers(r *mux.Router, clientservice client.UseCase, policyservice policy.UseCase, backupservice backup.UseCase) {
	r.Handle("/backups/start/{id}", startBackup(clientservice, policyservice, backupservice)).Methods("POST").Name("startBackup")
	r.Handle("/backups/list", listBackups(clientservice, policyservice, backupservice)).Methods("GET").Name("listBackups")
	r.Handle("/backups/get/{id}", getBackup(clientservice, policyservice, backupservice)).Methods("GET").Name("getBackup")
}
