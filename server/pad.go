package server

import "github.com/gin-gonic/gin"

func servePAD(r *gin.Engine) {
	r.GET("/pad", gin.HandlerFunc(providePAD))
}

func providePAD(c *gin.Context) {
	c.JSON(200, gin.H{
		"temp": 20,
		"pres": 1000,
		"humi": 70,
	})
}