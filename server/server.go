package server

import "../Error"
import "github.com/gin-gonic/gin"

var errorChan chan<- Error.Error = nil

func Server(ch chan Error.Error) {
	errorChan = ch
	r := gin.Default()
	serveRMC(r)
	servePAD(r)
	r.Run(":23500")
}