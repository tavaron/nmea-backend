package server

import (
	"github.com/gin-gonic/gin"

	"../Error"
	"../mongoReaders"
)

func servePAD(r *gin.Engine) {
	r.GET("/pad", gin.HandlerFunc(providePAD))
}

func providePAD(c *gin.Context) {
	data := mongoReaders.ReadLastPAD(db)
	if data.Id == 0 {
		errorChan <- Error.New(Error.High, "received invalid PAD data")
		c.JSON(404, nil)
	} else {
		c.JSON(200, gin.H{
			"time":  data.Id,
			"devID": data.Data[0].DeviceID,
			"temp":  data.Data[0].Temperature,
			"pres":  data.Data[0].Pressure,
			"humi":  data.Data[0].Humidity,
		})
	}
}
