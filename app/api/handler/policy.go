package handler

import (
	"encoding/json"
	"fmt"
	"log"

	//"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/wade-sam/fyp-backup-server/api/handler/presenter"
	"github.com/wade-sam/fyp-backup-server/entity"
	"github.com/wade-sam/fyp-backup-server/usecase/client"

	//"github.com/wade-sam/fyp-backup-server/usecase/dispatcher"
	"github.com/wade-sam/fyp-backup-server/usecase/policy"
)

func getPolicy(policy policy.UseCase, client client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Reading Policy"
		vars := mux.Vars(r)
		data, err := policy.GetPolicy(vars["id"])
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
		var clientnames [][]string
		for _, x := range data.Clients {
			name, err := client.GetClientName(x)
			if err != nil {
				err = data.RemoveClient(x)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(errorMessage))
					return
				}
				e := policy.UpdatePolicy(data)
				if e != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(errorMessage))
					return
				}
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(errorMessage))
				return
			}
			c := []string{name, x}
			clientnames = append(clientnames, c)
		}
		var jPolicy presenter.Policy
		mapstructure.Decode(data, &jPolicy)
		jPolicy.Clients = clientnames
		if err := json.NewEncoder(w).Encode(jPolicy); err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
		}
	})
}

func createPolicy(client client.UseCase, policy policy.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Could not create policy"
		var input presenter.Policy
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		var policyinput entity.Policy
		//holdClients := policyinput.Clients
		log.Println(input)
		mapstructure.Decode(input, &policyinput)
		var inputClients []string
		for _, i := range input.Clients {
			inputClients = append(inputClients, i[1])
		}
		id, err := policy.CreatePolicy(input.Policyname, input.RunTime, input.Type, input.Retention, input.Fullbackup, input.IncBackup, inputClients)
		if err != nil {
			log.Println("CREATE ERROR", err)
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

func updatePolicy(client client.UseCase, policy policy.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Could not update client"
		var input presenter.Policy
		err := json.NewDecoder(r.Body).Decode(&input)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		cl := []string{}
		for _, j := range input.Clients {
			cl = append(cl, j[1])
		}
		//cl := input.Clients[0][1]
		policyinput := entity.Policy{
			PolicyID:   input.PolicyID,
			Policyname: input.Policyname,
			Clients:    cl,
			Retention:  input.Retention,
			State:      input.State,
			Type:       input.Type,
			Fullbackup: input.Fullbackup,
			IncBackup:  input.IncBackup,
		}
		//var policyinput entity.Policy
		//mapstructure.Decode(input, &policyinput)
		log.Println("policy", policyinput)
		storedpolicy, err := policy.GetPolicy(policyinput.PolicyID)

		if storedpolicy == nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
		err = policy.UpdatePolicy(&policyinput)
		if err != nil {
			log.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Could not update client"))
		}
		w.WriteHeader(http.StatusCreated)
	})
}
func updateClientPolicies(Type string, list []string, policyID string, c client.UseCase) error {
	for _, j := range list {
		client, err := c.GetClient(j)
		if err != nil {
			return err
		}
		if client == nil {
			return err
		}
		if Type == "remove" {
			for x, z := range client.Policies {
				if z == j {
					client.Policies = append(client.Policies[:x], client.Policies[x+1])
					break
				}
			}
			err := c.UpdateClient(client)
			if err != nil {
				return err
			}
		} else {
			client.Policies = append(client.Policies, j)
			err := c.UpdateClient(client)
			if err != nil {
				return err
			}
		}
		return nil
	}
	return nil
}

func deletePolicy(policy policy.UseCase, client client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Deleting Policy"
		vars := mux.Vars(r)
		err := policy.DeletePolicy(vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func listPolicies(policy policy.UseCase, client client.UseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "Error Reading Clients"
		var data []*entity.Policy
		data, err := policy.ListPolicies()
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
		var jpolicies []*presenter.Policy
		for _, j := range data {
			var clientnames [][]string
			for _, x := range j.Clients {
				log.Println("policies", j)
				client, err := client.GetClientName(x)
				if err != nil {

					err = j.RemoveClient(x)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("errorMessage"))
						return
					}
					log.Println(j)
					e := policy.UpdatePolicy(j)
					if e != nil {
						log.Println("REMOVE CLIENT", e)
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte(errorMessage))
						return
					}

					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte(" 2 errorMessage"))
					return
				}
				c := []string{client, x}
				clientnames = append(clientnames, c)
			}
			var jPolicy presenter.Policy

			mapstructure.Decode(j, &jPolicy)
			//presenterbackuprun := presenter.BackupRun{}
			jPolicy.Clients = clientnames
			for _, b := range j.BackupRun {
				//presenterbackuprun
				log.Println("API BACKUPRUN", b)
			}
			jpolicies = append(jpolicies, &jPolicy)
		}

		if err := json.NewEncoder(w).Encode(jpolicies); err != nil {
			fmt.Println("ERROR", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))

		}
	})
}

func MakePolicyHolders(r *mux.Router, clientservice client.UseCase, policyservice policy.UseCase) {
	r.Handle("/api/policies/get/{id}", getPolicy(policyservice, clientservice)).Methods("GET").Name("getPolicy")
	r.Handle("/api/policies/list", listPolicies(policyservice, clientservice)).Methods("GET").Name("listPolicies")
	r.Handle("/api/policies/update", updatePolicy(clientservice, policyservice)).Methods("PUT").Name("updatePolicy")
	r.Handle("/api/policies/create", createPolicy(clientservice, policyservice)).Methods("POST").Name("createPolicy")
	r.Handle("/api/policies/delete/{id}", deletePolicy(policyservice, clientservice)).Methods("DELETE").Name("deletePolicy")

}
