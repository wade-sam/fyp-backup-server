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
	"github.com/wade-sam/fyp-backup-server/usecase/client"
	"github.com/wade-sam/fyp-backup-server/usecase/dispatcher"
	"github.com/wade-sam/fyp-backup-server/usecase/policy"
)

func getClient(client client.UseCase, policy policy.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Reading Client"
		vars := mux.Vars(r)
		data, err := client.GetClient(vars["id"])
		w.Header().Set("Content-Type", "application/json")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		if data == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errorMessage))
			return
		}
		var policynames []string
		for _, x := range data.Policies {
			policy, err := policy.GetPolicyName(x)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}
			policynames = append(policynames, policy)
		}
		jclient := presenter.Client{
			ID:            data.ID,
			ConsumerID:    data.ConsumerID,
			Clientname:    data.Clientname,
			Policies:      policynames,
			Directorytree: data.Directorytree,
			Ignorepath:    data.Ignorepath,
			Backups:       data.Backups,
		}
		if err := json.NewEncoder(w).Encode(jclient); err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))

		}

	})
}

func updateClient(client client.UseCase, policy policy.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Could not update client"
		var input presenter.Client
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		clientinput := entity.Client{
			ID:            input.ID,
			ConsumerID:    input.ConsumerID,
			Clientname:    input.Clientname,
			Policies:      input.Policies,
			Directorytree: input.Directorytree,
			Ignorepath:    input.Ignorepath,
			Backups:       input.Backups,
		}
		log.Println("client", clientinput)
		err = client.UpdateClient(&clientinput)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not update client"))
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func newClient(client client.UseCase, policy policy.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error writing Clients"
		var input struct {
			ConsumerID string `json:"consumerid"`
			ClientName string `json:"clientname"`
		}
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if input.ClientName == "" || input.ConsumerID == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		id, err := client.CreateClient(input.ClientName, input.ConsumerID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(id); err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))

		}
	})

}

func deleteClient(client client.UseCase, policy policy.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Deleting Client"
		vars := mux.Vars(r)
		err := client.DeleteClient(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func searchNewClient(dispatcher dispatcher.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Reading Clients"
		id, err := dispatcher.SearchForNewClient()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("No New Client Found"))
			return
		}
		if err := json.NewEncoder(w).Encode(id); err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}

	})
}

func directoryScan(dispatcher dispatcher.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Requesting Directory Scan"
		vars := mux.Vars(r)
		directory, err := dispatcher.GetDirectoryScan(vars["id"])
		if err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}

		if err := json.NewEncoder(w).Encode(directory); err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))

		}

	})
}
func showPolicies(client string, policy policy.UseCase) ([]string, error) {
	clientPolicies := []string{}
	data, err := policy.ListPolicies()
	if err != nil {
		return nil, err
	}
	for _, j := range data {
		c, ok := j.GetClient(client)
		if ok != nil {
			return nil, err
		} else if c != "" {
			log.Println("showPolicies", j.PolicyID, c)
			policyname, err := policy.GetPolicyName(j.PolicyID)
			if err != nil {
				log.Println("ERROR", err)
				return nil, err
			}
			clientPolicies = append(clientPolicies, policyname)
		}
	}
	return clientPolicies, nil
}

func listClients(client client.UseCase, policy policy.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Reading Clients"
		var data []*entity.Client
		data, err := client.ListClients()
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
		var jclients []*presenter.Client
		for _, j := range data {
			policies, err := showPolicies(j.ID, policy)
			if err != nil {
				log.Println("ERROR", err, j.ID)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(errorMessage))
				return
			}
			j.Policies = policies
			var jclient presenter.Client
			mapstructure.Decode(j, &jclient)

			jclients = append(jclients, &jclient)
		}

		if err := json.NewEncoder(w).Encode(jclients); err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))

		}

	})
}

func MakeClientHandlers(r *mux.Router, clientservice client.UseCase, policyservice policy.UseCase, dispatcherservice dispatcher.UseCase) {
	r.Handle("/api/clients/list", listClients(clientservice, policyservice)).Methods("GET").Name("listClients")
	r.Handle("/api/clients/get/{id}", getClient(clientservice, policyservice)).Methods("GET").Name("getClient")
	r.Handle("/api/clients/create", newClient(clientservice, policyservice)).Methods("POST").Name("createClient")
	r.Handle("/api/clients/delete/{id}", deleteClient(clientservice, policyservice)).Methods("DELETE").Name("deleteClient")
	r.Handle("/api/clients/update", updateClient(clientservice, policyservice)).Methods("PUT").Name("updateClient")
	r.Handle("/api/clients/search", searchNewClient(dispatcherservice)).Methods("GET").Name("searchNewClient")
	r.Handle("/api/clients/scan/{id}", directoryScan(dispatcherservice)).Methods("GET").Name("directoryScan")
}
