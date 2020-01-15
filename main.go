package main

import "./Error"
import "./server"

func main() {
	errorChan := make(chan Error.Error)
	go server.Server(errorChan)
	select {}
}
