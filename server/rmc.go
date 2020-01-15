package server

import (
	"github.com/gin-gonic/gin"

	"../Error"
	"../mongoReaders"
)

func serveRMC(r *gin.Engine) {
	r.GET("/rmc", gin.HandlerFunc(provideRMC))
}

func provideRMC(c *gin.Context) {
	data := mongoReaders.ReadLastRMC(db)
	if data.Id == 0 {
		errorChan <- Error.New(Error.High, "received invalid PAD data")
		c.JSON(404, nil)
	} else {
		c.JSON(200, gin.H{
			"time":               data.Id,
			"devID":              data.Data[0].DeviceID,
			"latitude":           data.Data[0].Latitude,
			"longitude":          data.Data[0].Longitude,
			"speed":              data.Data[0].Speed,
			"true_course":        data.Data[0].TrueCourse,
			"magnetic_variation": data.Data[0].MagneticVariation,
		})
	}
}
