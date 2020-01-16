package server

import (
	"github.com/gin-gonic/gin"

	"../Error"
	"../mongoReaders"
)

func routeRMC(router *gin.Engine) {
	router.GET("/rmc/:amount/:interval", gin.HandlerFunc(serveRMC))
	router.GET("/rmc/:amount", gin.HandlerFunc(serveRMC))
	router.GET("/rmc", gin.HandlerFunc(serveRMC))

}

func convertRMC(rmc mongoReaders.ResultRMC) gin.H {
	return gin.H{
		"time":               rmc.Id,
		"devID":              rmc.Data[0].DeviceID,
		"latitude":           rmc.Data[0].Latitude,
		"longitude":          rmc.Data[0].Longitude,
		"speed":              rmc.Data[0].Speed,
		"true_course":        rmc.Data[0].TrueCourse,
		"magnetic_variation": rmc.Data[0].MagneticVariation,
	}
}

func serveRMC(c *gin.Context) {
	data := mongoReaders.ReadRMC(db, getAmount(c), getInterval(c))
	if data == nil {
		errorChan <- Error.New(Error.Warning, c.ClientIP()+"no RMC data received from MongoDB")
		c.JSON(500, "no RMC data received from MongoDB")
		return
	}

	var send []gin.H
	for _, rmc := range data {
		send = append(send, convertRMC(rmc))
	}
	c.JSON(200, send)
}
