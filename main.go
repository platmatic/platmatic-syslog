package main

import (
	"gopkg.in/mcuadros/go-syslog.v2"
	"fmt"
)

func main() {
	channel := make(syslog.LogPartsChannel)
	handler := syslog.NewChannelHandler(channel)

	server := syslog.NewServer()
	server.SetFormat(syslog.RFC5424)
	server.SetHandler(handler)
	server.ListenUDP("0.0.0.0:8080")
	server.ListenTCP("0.0.0.0:8080")

	server.Boot()

	go func(channel syslog.LogPartsChannel) {
		for logParts := range channel {
			fmt.Println(logParts)
		}
	}(channel)

	fmt.Println("Starting server on port: 8080")
	server.Wait()
}
