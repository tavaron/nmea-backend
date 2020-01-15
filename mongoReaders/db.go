package mongoReaders

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"../Error"
)

type dbConfig struct {
	username string
	password string
	uri      string
	database string
}

var errorChan chan<- Error.Error

func MongoDB(ch chan<- Error.Error) *mongo.Database {
	errorChan = ch
	config := dbConfig{
		username: "",
		password: "",
		uri:      "mongodb://10.0.23.254:27017",
		database: "NMEA0183",
	}
	opts := options.Client().ApplyURI(config.uri)
	client, err := mongo.NewClient(opts)
	if err != nil {
		errorChan <- Error.Err(Error.High, err)
		return nil
	}

	err = client.Connect(context.Background())
	if err != nil {
		errorChan <- Error.Err(Error.High, err)
		return nil
	}

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		errorChan <- Error.Err(Error.High, err)
		return nil
	}

	return client.Database(config.database)
}
