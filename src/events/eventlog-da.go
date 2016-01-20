package events

import (
	"dataaacess"
	"log"
	"time"
)

type HookEvent struct {
	Id      string
	Type    string
	Payload string
}

type BroadcastEvent struct {
	EventId string
	App     string
	Url     string
	Sent    time.Time
	Respose int
	Error   error
}

type EventLog struct {
	Event HookEvent
}

func writeEvent(s HookEvent) {
	session := dataaccess.GetSession()
	defer session.Close()

	c := session.DB(dataaccess.Database).C("eventlog")
	err := c.Insert(s)
	if err != nil {
		log.Fatal(err)
	}
}

func writeBroadcast(s BroadcastEvent) {
	session := dataaccess.GetSession()
	defer session.Close()

	c := session.DB(dataaccess.Database).C("broadcastlog")
	err := c.Insert(s)
	if err != nil {
		log.Fatal(err)
	}
}
