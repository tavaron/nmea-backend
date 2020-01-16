package mongoReaders

import (
	"../Error"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RMC struct {
	// device id
	DeviceID uint32

	// North = positive
	// South = negative
	Latitude float64

	// East = positive
	// West = negative
	Longitude float64

	// speed in knots, negative if n/a
	Speed float64

	// true course in degrees, negative if n/a
	TrueCourse float64

	// East = positive
	// West = negative
	MagneticVariation float64
}

type ResultRMC struct {
	Id   int64 `bson:"_id"`
	Data []RMC `bson:"data"`
}

func ReadRMC(db *mongo.Database, amount int64, interval int64) []ResultRMC {
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

	cur, err := db.Collection("RMC").Find(context.TODO(), bson.D{}, opts)
	if err != nil {
		errorChan <- Error.Err(Error.Low, err)
		return nil
	}
	defer cur.Close(context.TODO())

	var result []ResultRMC
	var i int64 = 0
	for cur.Next(context.TODO()) {
		if i%interval == 0 {
			var rmc ResultRMC
			err := cur.Decode(&rmc)
			if err != nil {
				errorChan <- Error.Err(Error.Low, err)
			} else {
				result = append(result, rmc)
			}
		}
		i++
	}

	return result
}
