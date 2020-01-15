package server

import "github.com/gin-gonic/gin"

func serveRMC(r *gin.Engine) {
	r.GET("/rmc", gin.HandlerFunc(provideRMC))
}

func provideRMC(c *gin.Context) {
	c.JSON(200, gin.H{
		"lat": 		0.1,
		"latDir": 	"E",
		"lon": 		0.1,
		"lonDir": 	"N",
	})
}
