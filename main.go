package main

import "./Error"
import "./server"

func main() {
	errorChan := make(chan Error.Error)
	go server.Server(errorChan)

	for err := range errorChan {
		switch err.Lvl {
		case Error.Debug:
			println("[DEBUG] " + err.Text)
		case Error.Info:
			println("[INFO]  " + err.Text)
		case Error.Warning:
			println("[WARN]  " + err.Text)
		case Error.Low:
			println("[LOW]   " + err.Text)
		case Error.High:
			println("[HIGH]  " + err.Text)
		case Error.Fatal:
			println("[FATAL] " + err.Text)
		default:
			println("[UNKWN] " + err.Text)
		}
	}
}
