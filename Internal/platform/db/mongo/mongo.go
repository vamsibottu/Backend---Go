package mongo

import (
	"gopkg.in/mgo.v2"
)

const (//TODO need to make these configurable
	host         = "localhost"
	databaseName = "myDB"
)

func InitMongo() (*mgo.Session, error) {
	mgoSession, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}

	mgoSession.SetMode(mgo.Monotonic, true)

	return mgoSession.Clone(), nil
}

func WithCollection(collection string, s func(*mgo.Collection) error) error {
	//TODO: need refactoring
	session, _ := InitMongo()
	defer session.Close()
	c := session.DB(databaseName).C(collection)
	return s(c)
}
