package mongo

import "gopkg.in/mgo.v2"
import (
	log "github.com/sirupsen/logrus"
	"github.com/zxhaaa6/DataVC/module/client"
)

var mongoClient *mgo.Database

func ConnectMongo(connectionOptions client.ConnectOptionType) (*mgo.Database, error) {
	mongoUrl := "mongodb://" + connectionOptions.Host + ":" + connectionOptions.Port
	session, err := mgo.Dial(mongoUrl)
	if err != nil {
		log.Fatal("[Mongodb]connect DB server fatal err: ", err)
		return nil, err
	}
	mongoClient = session.DB(connectionOptions.Database)
	log.Info("[Mongodb]DB connection has been established successfully.")
	return mongoClient, nil
}

func GetMongoClient() *mgo.Database {
	return mongoClient
}
