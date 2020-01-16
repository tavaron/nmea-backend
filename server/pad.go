package server

import (
	"../Error"
	"../mongoReaders"
	"github.com/gin-gonic/gin"
)

func routePAD(router *gin.Engine) {
	router.GET("/pad/:amount/:interval", gin.HandlerFunc(servePAD))
	router.GET("/pad/:amount", gin.HandlerFunc(servePAD))
	router.GET("/pad", gin.HandlerFunc(servePAD))

}

func convertPAD(pad mongoReaders.ResultPAD) gin.H {
	return gin.H{
		"time":        pad.Id,
		"deviceID":    pad.Data[0].DeviceID,
		"temperature": pad.Data[0].Temperature,
		"pressure":    pad.Data[0].Pressure,
		"humidity":    pad.Data[0].Humidity,
	}
}

func servePAD(c *gin.Context) {
	data := mongoReaders.ReadPAD(db, getAmount(c), getInterval(c))
	if data == nil {
		errorChan <- Error.New(Error.Warning, c.ClientIP()+"no PAD data received from MongoDB")
		c.JSON(500, "no PAD data received from MongoDB")
		return
	}

	var send []gin.H
	for _, pad := range data {
		send = append(send, convertPAD(pad))
	}
	c.JSON(200, send)

}
