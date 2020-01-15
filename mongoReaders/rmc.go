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

func ReadLastRMC(db *mongo.Database) ResultRMC {
	if db == nil {
		errorChan <- Error.New(Error.Fatal, "database not connected")
		return ResultRMC{
			Id:   0,
			Data: nil,
		}
	}

	coll := db.Collection("RMC")
	var entry ResultRMC
	opts := options.FindOne()
	opts.SetSort(bson.D{{"_id", -1}})
	err := coll.FindOne(context.TODO(), bson.D{}, opts).Decode(&entry)

	if err != nil {
		errorChan <- Error.Err(Error.Low, err)
		return ResultRMC{
			Id:   0,
			Data: nil,
		}
	}
	return entry
}
