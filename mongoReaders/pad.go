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

func ReadPAD(db *mongo.Database, amount int64, interval int64) []ResultPAD {
	if db == nil {
		errorChan <- Error.New(Error.Fatal, "database not connected")
		return nil
	}

	if amount <= 0 {
		amount = 1
	} else if amount > amountLimit {
		amount = amountLimit
	}

	if interval <= 0 {
		interval = 1
	}

	opts := options.Find()
	opts.SetLimit(amount * interval)
	opts.SetSort(bson.D{{"_id", -1}})

	cur, err := db.Collection("PAD").Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		errorChan <- Error.Err(Error.Low, err)
		return nil
	}
	defer cur.Close(context.TODO())

	var result []ResultPAD
	var i int64 = 0
	for cur.Next(context.TODO()) {
		if i%interval == 0 {
			var pad ResultPAD
			err := cur.Decode(&pad)
			if err != nil {
				errorChan <- Error.Err(Error.Low, err)
			} else {
				result = append(result, pad)
			}
		}
		i++
	}

	return result
}
