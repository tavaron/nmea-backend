package server

import (
	"../Error"
	"../mongoReaders"
	"strconv"

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
	routeRMC(r)
	routePAD(r)
	err := r.Run(":23500")
	if err == nil {
		errorChan <- Error.Err(Error.High, err)
	}
}

func getInterval(c *gin.Context) int64 {
	if len(c.Param("interval")) == 0 {
		return 1
	}
	interval, err := strconv.Atoi(c.Param("interval"))
	if err != nil {
		c.JSON(400, "received malformed interval")
		errorChan <- Error.New(Error.Warning, c.ClientIP()+": received malformed amount for RMC data")
		return 1
	}
	return int64(interval)
}

func getAmount(c *gin.Context) int64 {
	if len(c.Param("amount")) == 0 {
		return 1
	}
	amount, err := strconv.Atoi(c.Param("amount"))
	if err != nil {
		c.JSON(400, "received malformed amount")
		errorChan <- Error.New(Error.Warning, c.ClientIP()+": received malformed amount for RMC data")
		return 1
	}
	return int64(amount)
}
