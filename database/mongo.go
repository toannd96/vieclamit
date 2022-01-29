package database

import "gopkg.in/mgo.v2"

const (
	host   = "localhost"
	dbName = "vieclamit"
)

// Mongo struct
type Mongo struct {
	Db *mgo.Database
}

// CreateConn create connection to mongodb
func (m *Mongo) CreateConn() {
	sess, err := mgo.Dial(host)
	if err != nil {
		panic(err)
	}
	m.Db = sess.DB(dbName)
}
