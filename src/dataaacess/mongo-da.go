package dataaccess

import (
	"os"

	"gopkg.in/mgo.v2"
)

var (
	mongourl string
	Database string
)

func GetSession() *mgo.Session {
	session, err := mgo.Dial(mongourl)
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

func init() {

	if mongourl = os.Getenv("MONGO_URL"); mongourl == "" {
		mongourl = "mongodb://localhost"
	}

	Database = "webhook"
}
