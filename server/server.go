package server

import (
	"../Error"
	"../mongoReaders"

	"go.mongodb.org/mongo-driver/mongo"
)
import "github.com/gin-gonic/gin"

var errorChan chan<- Error.Error
var db *mongo.Database

func Server(ch chan<- Error.Error) {
	errorChan = ch
	db = mongoReaders.MongoDB(ch)
	if db == nil {
		errorChan <- Error.New(Error.Fatal, "could not connect to database")
	}

	r := gin.Default()
	serveRMC(r)
	servePAD(r)
	err := r.Run(":23500")
	if err == nil {
		errorChan <- Error.Err(Error.High, err)
	}
}
