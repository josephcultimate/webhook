package subscriptions

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Configure(router *mux.Router) {

	router.HandleFunc("/subscription", newSubscription).Methods("POST")
}

func newSubscription(resp http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var jsonData Subscription
	decoder.Decode(&jsonData)

	write(jsonData)
}
