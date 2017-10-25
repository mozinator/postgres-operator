package userservice

/*
Copyright 2017 Crunchy Data Solutions, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	//apiserver "github.com/crunchydata/postgres-operator/apiserver"
	msgs "github.com/crunchydata/postgres-operator/apiservermsgs"
	//"github.com/gorilla/mux"
	"net/http"
)

// UserHandler ...
// pgo user XXXX
func UserHandler(w http.ResponseWriter, r *http.Request) {

	log.Infoln("userservice.UserHandler called")

	var request msgs.UserRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := User(&request)

	json.NewEncoder(w).Encode(resp)
}