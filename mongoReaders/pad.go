package mongoReaders

import (
	"../Error"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PAD struct {
	// device id
	DeviceID uint32

	// degree Celsius
	Temperature float64

	// percent RH
	Humidity float64

	// hecto Pascals
	Pressure float64
}

type ResultPAD struct {
	Id   int64 `bson:"_id"`
	Data []PAD `bson:"data"`
}

func ReadLastPAD(db *mongo.Database) ResultPAD {
	if db == nil {
		errorChan <- Error.New(Error.Fatal, "database not connected")
		return ResultPAD{
			Id:   0,
			Data: nil,
		}
	}

	coll := db.Collection("PAD")
	var entry ResultPAD
	opts := options.FindOne()
	opts.SetSort(bson.D{{"_id", -1}})
	err := coll.FindOne(context.TODO(), bson.D{}, opts).Decode(&entry)

	if err != nil {
		errorChan <- Error.Err(Error.Low, err)
		return ResultPAD{
			Id:   0,
			Data: nil,
		}
	}
	return entry
}
