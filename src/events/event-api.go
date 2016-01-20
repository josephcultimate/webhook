package events

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"subscriptions"
	"time"

	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
)

func Configure(router *mux.Router) {

	router.HandleFunc("/event", newEvent).Methods("POST")
	router.HandleFunc("/testevent", testEvent).Methods("POST")
}

func newEvent(resp http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var jsonData HookEvent
	decoder.Decode(&jsonData)

	id, _ := uuid.NewV4()

	jsonData.Id = id.String()

	writeEvent(jsonData)

	subscriptions := subscriptions.FindEventSubscriptions(jsonData.Type)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	for _, v := range subscriptions {

		jsonStr, _ := json.Marshal(jsonData)

		resp, err := client.Post(v.Url, "application/json", bytes.NewBuffer(jsonStr))

		writeBroadcast(BroadcastEvent{
			EventId: jsonData.Id,
			App:     v.App,
			Url:     v.Url,
			Sent:    time.Now().UTC(),
			Respose: resp.StatusCode,
			Error:   err,
		})

		if err != nil {
			// log for retry policy
			panic(err)
		}

		resp.Body.Close()
	}
}

func testEvent(resp http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var jsonData HookEvent
	decoder.Decode(&jsonData)

	fmt.Println("Got this: ", jsonData)
}
