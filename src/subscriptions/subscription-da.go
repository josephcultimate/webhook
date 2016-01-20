package subscriptions

import (
	"dataaacess"
	"log"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

type Subscription struct {
	App       string
	EventType string
	Url       string
}

func write(s Subscription) {
	session := dataaccess.GetSession()
	defer session.Close()

	c := session.DB(dataaccess.Database).C("subscriptions")
	err := c.Insert(s)
	if err != nil {
		log.Fatal(err)
	}
}

func FindEventSubscriptions(eventType string) []Subscription {
	session := dataaccess.GetSession()
	defer session.Close()

	c := session.DB(dataaccess.Database).C("subscriptions")

	namespace := strings.Split(eventType, ".")

	var results []Subscription
	c.Find(bson.M{"$or": []interface{}{bson.M{"eventtype": "*"}, bson.M{"eventtype": namespace[0] + ".*"}, bson.M{"eventtype": eventType}}}).All(&results)

	return results
}
